package signer

import (
	"anti-validators/config"
	"anti-validators/contracts"
	"anti-validators/db/models"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/jinzhu/gorm"
)

func StartEthSigner(cfg config.Config, oracleAddress string, ctx context.Context, db *gorm.DB) error {
	client, err := ethclient.DialContext(ctx, cfg.Ethereum.NodeURL)
	if err != nil {
		return err
	}

	ethContractAddress := common.Address{}
	hexAddress, err := hexutil.Decode(cfg.Ethereum.ContractAddress)
	if err != nil {
		return err
	}
	ethContractAddress.SetBytes(hexAddress)

	contract, err := contracts.NewSupersymmetry(ethContractAddress, client)
	if err != nil {
		return err
	}

	var pubKeyOracles []string
	for i := 0; i < 5; i++ {
		admin, err := contract.Admins(nil, big.NewInt(int64(i)))
		if err != nil {
			return err
		}
		pubKeyOracles = append(pubKeyOracles, admin.String())
	}
	if err != nil {
		return err
	}

	privBytes, err := hexutil.Decode(cfg.PrivateKey)
	if err != nil {
		return err
	}
	privateKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: secp256k1.S256(),
		},
		D: new(big.Int),
	}
	privateKey.D.SetBytes(privBytes)
	privateKey.PublicKey.X, privateKey.PublicKey.Y = privateKey.PublicKey.Curve.ScalarBaseMult(privBytes)

	go func() {
		for true {
			err := signOrTestConfirmEthRequest(contract, privateKey, oracleAddress, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}

			err = getEthRequest(cfg.Ips, oracleAddress, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}
			//TODO: add finilize sign

			time.Sleep(10 * time.Second)
		}
	}()

	return nil
}

func signOrTestConfirmEthRequest(contract *contracts.Supersymmetry, privKey *ecdsa.PrivateKey, oracleAddress string, db *gorm.DB) error {
	var requests []models.Request
	if err := db.Where(&models.Request{Status: models.New, ChainType: models.Ethereum}).Find(&requests).Error; err != nil {
		return err
	}
	transactOpts := bind.NewKeyedTransactor(privKey)
	for _, request := range requests {
		var sign models.Signs
		db.Where(&models.Signs{RequestId: request.Id, ValidatorAddress: oracleAddress}).First(&sign)
		if sign.ValidatorAddress == oracleAddress {
			continue
		}
		hash, err := hexutil.Decode(request.Id)
		if err != nil {
			continue
		}
		var hash32 [32]byte
		copy(hash32[:], hash)
		_, err = contract.ChangeStatusTest(transactOpts, hash32, uint8(models.Success))
		if err != nil {
			return err
		}
		//TODO: Sign
		/*if err := db.Save(models.Signs{
			RequestId:        request.Id,
			ValidatorAddress: oracleAddress,
			Sign:             signedText.Signature,
			ValidatorPubKey:  signedText.PublicKey,
			CreatedAt:        time.Now(),
		}).Error; err != nil {
			return err
		}*/
	}
	return nil
}

func getEthRequest(ips []string, oracleAddress string, db *gorm.DB) error {
	var requests []models.Request
	if err := db.Where(&models.Request{Status: models.New, ChainType: models.Waves}).Find(&requests).Error; err != nil {
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
				if existSign.ValidatorAddress == oracleAddress {
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
