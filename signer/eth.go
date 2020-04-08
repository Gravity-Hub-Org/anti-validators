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
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"

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
		pubKeyOracles = append(pubKeyOracles, strings.ToLower(admin.String()))
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
			err := signEthRequest(contract, privateKey, oracleAddress, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}

			err = getEthRequest(cfg.Ips, oracleAddress, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}
			err = finalizeEthRequest(contract, privateKey, pubKeyOracles, oracleAddress, cfg.BftCoefficient, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}

			time.Sleep(time.Duration(cfg.Timeout) * time.Second)
		}
	}()

	return nil
}

func signEthRequest(contract *contracts.Supersymmetry, privKey *ecdsa.PrivateKey, oracleAddress string, db *gorm.DB) error {

	var requests []models.Request
	if err := db.Where(&models.Request{Status: models.New}).Where("chain_type = ?", models.Ethereum).Find(&requests).Error; err != nil {
		return err
	}
	for _, request := range requests {
		var sign models.Signs
		db.Where(&models.Signs{RequestId: request.Id, ValidatorAddress: oracleAddress}).First(&sign)

		if sign.ValidatorAddress == oracleAddress {
			continue
		}

		var valid bool
		//TODO
		switch request.Type {
		case models.Lock:
		case models.Unlock:
		case models.Mint:
		case models.Burn:
		}
		/*if request.Type == models.BurOrLock {
			valid = true
		} else if request.Type == models.MintOrUnlock {
			token, err := contract.Tokens(nil, common.HexToAddress(request.AssetId))
			if err != nil {
				return err
			}
			//	valid = isEthValid(request, request.AssetId, token.AssetId, db)
		}*/
		if !valid {
			continue
		}

		msg, err := hexutil.Decode(request.Id)
		if err != nil {
			return err
		}

		status := models.Success
		msg = append(msg, byte(status))
		signature, err := signMsg(msg, privKey)
		println(request.Id + ":" + hexutil.Encode(signature))
		if err := db.Save(models.Signs{
			RequestId:        request.Id,
			ValidatorAddress: oracleAddress,
			Sign:             hexutil.Encode(signature),
			ValidatorPubKey:  oracleAddress,
			CreatedAt:        time.Now(),
			Status:           uint8(status),
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func getEthRequest(ips []string, oracleAddress string, db *gorm.DB) error {
	var requests []models.Request
	if err := db.Where(&models.Request{Status: models.New}).Where("chain_type = ?", models.Ethereum).Find(&requests).Error; err != nil {
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

func finalizeEthRequest(contract *contracts.Supersymmetry, privKey *ecdsa.PrivateKey, pubKeyOracles []string, oracleAddress string, bftCoefficient int, db *gorm.DB) error {
	var requests []models.Request
	if err := db.Where(&models.Request{Status: models.New}).Where("chain_type = ?", models.Ethereum).Find(&requests).Error; err != nil {
		return err
	}
	transactOpt := bind.NewKeyedTransactor(privKey)
	status := models.Success
	for _, request := range requests {
		bytesRequest, err := hexutil.Decode(request.Id)
		if err != nil {
			fmt.Printf("Error: %s \n", err.Error())
			continue
		}

		var bytesArrayRequest [32]byte
		copy(bytesArrayRequest[:], bytesRequest[:])
		rq, err := contract.Requests(nil, bytesArrayRequest)
		if rq.Status == uint8(models.Success) {
			request.Status = models.Success
			db.Save(request)

			continue
		}

		var signs []models.Signs
		if err := db.Where(&models.Signs{RequestId: request.Id, Status: uint8(status)}).Find(&signs).Error; err != nil {
			return err
		}

		if len(signs) < bftCoefficient {
			continue
		}

		var r [5][32]byte
		var s [5][32]byte
		var v [5]uint8

		for i, pubKey := range pubKeyOracles {
			for _, sign := range signs {
				if sign.ValidatorPubKey == pubKey {
					bytesSign, err := hexutil.Decode(sign.Sign)
					if err != nil {
						fmt.Printf("Error: %s \n", err.Error())
						continue
					}
					copy(r[i][:], bytesSign[:32])
					copy(s[i][:], bytesSign[32:64])
					v[i] = bytesSign[64] + 27
				}
			}
		}

		bytesSliceRequestId, err := hexutil.Decode(request.Id)
		if err != nil {
			fmt.Printf("Error: %s \n", err.Error())
			continue
		}
		var bytesArrayRequestId [32]byte
		copy(bytesArrayRequestId[:], bytesSliceRequestId[:])

		tx, err := contract.ChangeStatus(transactOpt, bytesArrayRequestId, v, r, s, uint8(status))
		if err != nil {
			fmt.Printf("Error: %s \n", err.Error())
			continue
		}

		fmt.Printf("Tx finilize: %s \n", tx.Hash().String())
		request.Status = models.Success
		db.Save(request)
	}
	return nil
}

func signMsg(message []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	validationMsg := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message))
	validationHash := crypto.Keccak256(append([]byte(validationMsg), message[:]...))
	return crypto.Sign(validationHash, privKey)
}

//TODO
func isEthValid(request models.Request, tokenAddress string, assetId string, fromDecimals int, toDecimals int, db *gorm.DB) bool {
	/*var inputRequests []models.Request
	db.Where(&models.Request{Target: strings.ToLower(request.Owner), ChainType: models.Waves, Amount: converter.StrConvert(request.Amount, fromDecimals, toDecimals),
		Type: models.BurOrLock, Owner: request.Target, Status: models.Success, AssetId: assetId}).Find(&inputRequests)

	var outSuccessRequests []models.Request
	db.Where(&models.Request{Target: request.Target, Amount: request.Amount, Type: models.MintOrUnlock,
		Owner: strings.ToLower(request.Owner), Status: models.Success, AssetId: tokenAddress}).Where("chain_type = ?", models.Ethereum).Find(&outSuccessRequests)

	if len(inputRequests) <= len(outSuccessRequests) {
		return false
	}

	var outNewRequests []models.Request
	db.Where(&models.Request{Target: request.Target, Amount: request.Amount, Type: models.MintOrUnlock,
		Owner: strings.ToLower(request.Owner), Status: models.New, AssetId: tokenAddress}).Where("chain_type = ?", models.Ethereum).Find(&outNewRequests)
	if len(outNewRequests) == 0 {
		return false
	}
	outNewRequest := outNewRequests[0]
	if len(outNewRequests) > 1 {
		minHeight := outNewRequests[0].CreatedAt
		for _, rq := range outNewRequests {
			if minHeight > rq.CreatedAt {
				minHeight = rq.CreatedAt
				outNewRequest = rq
			}
		}
	}

	return request.Id == outNewRequest.Id*/
	return false
}
