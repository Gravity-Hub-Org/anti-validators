package signer

import (
	"anti-validators/config"
	"anti-validators/contracts"
	"anti-validators/models"
	"anti-validators/wavesapi"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
)

func StartSigner(cfg config.Config, ethOracleAddress string, wavesOracleAddress string, ctx context.Context, db *gorm.DB) error {
	ethClient, err := ethclient.DialContext(ctx, cfg.Ethereum.NodeURL)
	if err != nil {
		return err
	}

	ethContractAddress := common.Address{}
	hexAddress, err := hexutil.Decode(cfg.Ethereum.ContractAddress)
	if err != nil {
		return err
	}
	ethContractAddress.SetBytes(hexAddress)

	ethContract, err := contracts.NewSupersymmetry(ethContractAddress, ethClient)
	if err != nil {
		return err
	}

	var ethPubKeyOracles []string
	for i := 0; i < 5; i++ {
		admin, err := ethContract.Admins(nil, big.NewInt(int64(i)))
		if err != nil {
			return err
		}
		ethPubKeyOracles = append(ethPubKeyOracles, strings.ToLower(admin.String()))
	}

	privBytes, err := hexutil.Decode(cfg.PrivateKey)
	if err != nil {
		return err
	}
	ethPrivateKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: secp256k1.S256(),
		},
		D: new(big.Int),
	}
	ethPrivateKey.D.SetBytes(privBytes)
	ethPrivateKey.PublicKey.X, ethPrivateKey.PublicKey.Y = ethPrivateKey.PublicKey.Curve.ScalarBaseMult(privBytes)

	wavesClient := wavesapi.New(cfg.Waves.NodeURL, cfg.Waves.ApiKey)
	wavesContractState, err := wavesClient.GetStateByAddress(cfg.Waves.ContractAddress)
	if err != nil {
		return err
	}
	wavesPubKeyOracles := strings.Split(wavesContractState["admins"].Value.(string), ",")

	go func() {
		for true {
			err := signRequest(wavesClient, cfg.Waves.ContractAddress, ethContract, ethPrivateKey, ethOracleAddress, wavesOracleAddress, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}

			err = getSigns(cfg.Ips, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}
			err = finalizeEthRequest(ethContract, ethPrivateKey, ethPubKeyOracles, ethOracleAddress, cfg.BftCoefficient, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}

			err = finalizeWavesRequest(wavesClient, cfg.Waves.ContractAddress, wavesPubKeyOracles, wavesOracleAddress, cfg.BftCoefficient, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}

			time.Sleep(time.Duration(cfg.Timeout) * time.Second)
		}
	}()

	return nil
}

func signRequest(wavesClient *wavesapi.Node, wavesContractAddress string, ethContract *contracts.Supersymmetry, ethPrivKey *ecdsa.PrivateKey,
	ethOracleAddress string, wavesOracleAddress string, db *gorm.DB) error {
	var requests []models.Request
	if err := db.Where(&models.Request{Status: models.New}).Find(&requests).Error; err != nil {
		return err
	}
	for _, request := range requests {
		oracleAddress := ""
		if request.ChainType == models.Ethereum {
			oracleAddress = ethOracleAddress
		} else if request.ChainType == models.Waves {
			oracleAddress = wavesOracleAddress
		}

		var sign models.Signs
		db.Where(&models.Signs{RequestId: request.Id, ValidatorAddress: oracleAddress}).First(&sign)

		if sign.ValidatorAddress == oracleAddress {
			continue
		}

		wavesContractState, err := wavesClient.GetStateByAddress(wavesContractAddress)
		if err != nil {
			continue
		}

		var fromDecimals uint8
		var toDecimals uint8
		if request.ChainType == models.Ethereum {
			tokenFrom, err := ethContract.Tokens(nil, common.HexToAddress(request.AssetId))
			if err != nil || tokenFrom.Status != uint8(models.Success) {
				continue
			}
			ercAddress := wavesContractState["erc20_address_"+tokenFrom.AssetId].Value.(string)
			if request.AssetId != ercAddress {
				continue
			}

			status := uint8(wavesContractState["asset_status_"+tokenFrom.AssetId].Value.(float64))
			if status != uint8(models.Success) {
				continue
			}

			fromDecimals = tokenFrom.Decimals
			toDecimals = uint8(wavesContractState["asset_decimals_"+tokenFrom.AssetId].Value.(float64))
		} else if request.ChainType == models.Waves {
			ercAddress := wavesContractState["erc20_address_"+request.AssetId].Value.(string)
			tokenTo, err := ethContract.Tokens(nil, common.HexToAddress(ercAddress))
			if err != nil || tokenTo.Status != uint8(models.Success) {
				continue
			}

			if request.AssetId != tokenTo.AssetId {
				continue
			}
			status := uint8(wavesContractState["asset_status_"+request.AssetId].Value.(float64))
			if status != uint8(models.Success) {
				continue
			}

			toDecimals = tokenTo.Decimals
			fromDecimals = uint8(wavesContractState["asset_decimals_"+request.AssetId].Value.(float64))
		}

		status := models.Success
		var signature string
		var pubKey string
		if request.ChainType == models.Ethereum {
			msg, err := formatEthMsg(request.Id, status)
			if err != nil {
				continue
			}

			signatureBytes, err := signEthMsg(msg, ethPrivKey)
			pubKey = ethOracleAddress
			signature = hexutil.Encode(signatureBytes)
		} else if request.ChainType == models.Waves {
			msg := formatWavesMsg(request.Id, status)
			signedText, err := wavesClient.SignMsg(msg, oracleAddress)
			if err != nil {
				continue
			}
			pubKey = signedText.PublicKey
			signature = signedText.Signature
		}

		println(request.Id + ":" + signature)
		if err := db.Save(models.Signs{
			RequestId:        request.Id,
			ValidatorAddress: oracleAddress,
			Sign:             signature,
			ValidatorPubKey:  pubKey,
			CreatedAt:        time.Now(),
			Status:           status,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func getSigns(ips []string, db *gorm.DB) error {
	var requests []models.Request
	if err := db.Where(&models.Request{Status: models.New}).Find(&requests).Error; err != nil {
		return err
	}

	for _, request := range requests {
		for _, ip := range ips {
			var client = &http.Client{Timeout: 10 * time.Second}
			res, err := client.Get(ip + "/api/signs?requestHash=" + request.Id)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
				continue
			}

			if res.StatusCode == 200 {
				var result models.Signs
				err = json.NewDecoder(res.Body).Decode(&result)
				if err != nil {
					fmt.Printf("Error: %s \n", err.Error())
					continue
				}
				var existSign models.Signs
				db.Where(&models.Signs{RequestId: request.Id, ValidatorAddress: result.ValidatorAddress}).First(&existSign)
				if existSign.ValidatorAddress == result.ValidatorAddress {
					continue
				} else {
					if err := db.Save(result).Error; err != nil {
						return err
					}
				}
			}
			res.Body.Close()
		}
	}
	return nil
}
