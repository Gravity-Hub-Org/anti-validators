package signer

import (
	"anti-validators/config"
	"anti-validators/db/models"
	"anti-validators/wavesapi"
	"anti-validators/wavesapi/transactions"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

func StartWavesSigner(cfg config.Config, oracleAddress string, ctx context.Context, db *gorm.DB) error {
	wavesClient := wavesapi.New(cfg.Waves.NodeURL, cfg.Waves.ApiKey)

	contractState, err := wavesClient.GetStateByAddress(cfg.Waves.ContractAddress)
	if err != nil {
		return err
	}

	pubKeyOracles := strings.Split(contractState["admins"].Value.(string), ",")
	if err != nil {
		return err
	}

	go func() {
		for true {
			err := signWavesRequest(wavesClient, cfg.Waves.ContractAddress, oracleAddress, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}

			err = getWavesRequest(cfg.Ips, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}

			err = finalizeWavesRequest(wavesClient, cfg.Waves.ContractAddress, pubKeyOracles, oracleAddress, cfg.BftCoefficient, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}

			time.Sleep(time.Duration(cfg.Timeout) * time.Second)
		}
	}()

	return nil
}

func signWavesRequest(wavesClient *wavesapi.Node, contractAddress string, oracleAddress string, db *gorm.DB) error {
	var requests []models.Request
	if err := db.Where(&models.Request{Status: models.New, ChainType: models.Waves}).Find(&requests).Error; err != nil {
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
		/*
			if request.Type == models.BurOrLock {
				valid = true
			} else if request.Type == models.MintOrUnlock {
				/*state, err := wavesClient.GetStateByAddress(contractAddress)
				if err != nil {
					return err
				}
				//	valid = isWavesValid(request, state["erc20_address_"+request.AssetId].Value.(string), request.AssetId, db)
			}*/

		if !valid {
			continue
		}

		status := models.Success
		msg := request.Id + "_" + strconv.Itoa(int(models.Success))
		signedText, err := wavesClient.SignMsg(msg, oracleAddress)
		if err != nil {
			return err
		}
		println(request.Id + ":" + signedText.Message)
		if err := db.Save(models.Signs{
			RequestId:        request.Id,
			ValidatorAddress: oracleAddress,
			Sign:             signedText.Signature,
			ValidatorPubKey:  signedText.PublicKey,
			CreatedAt:        time.Now(),
			Status:           uint8(status),
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func getWavesRequest(ips []string, db *gorm.DB) error {
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

func finalizeWavesRequest(wavesClient *wavesapi.Node, contract string, pubKeyOracles []string, oracleAddress string, bftCoefficient int, db *gorm.DB) error {
	var requests []models.Request
	if err := db.Where(&models.Request{ChainType: models.Waves, Status: models.New}).Find(&requests).Error; err != nil {
		return err
	}

	state, err := wavesClient.GetStateByAddress(contract)
	if err != nil {
		return err
	}

	for _, request := range requests {
		status := models.Status(state["status_"+request.Id].Value.(float64))
		if status == models.Success {
			request.Status = models.Success
			db.Save(request)
			continue
		}

		var signs []models.Signs
		if err := db.Where(&models.Signs{RequestId: request.Id}).Find(&signs).Error; err != nil {
			return err
		}

		if len(signs) < bftCoefficient {
			continue
		}

		sortedSigns := ""
		for _, pubKey := range pubKeyOracles {
			if len(sortedSigns) > 0 {
				sortedSigns += ","
			}

			foundSign := "q"
			for _, sign := range signs {
				if sign.ValidatorPubKey == pubKey {
					foundSign = sign.Sign
					break
				}
			}

			sortedSigns += foundSign
		}

		tx := transactions.New(transactions.InvokeScript, oracleAddress)
		tx.NewInvokeScript(contract, transactions.FuncCall{
			Function: "changeStatus",
			Args: []transactions.FuncArg{
				{
					Value: request.Id,
					Type:  "string",
				},
				{
					Value: sortedSigns,
					Type:  "string",
				},
				{
					Value: int(models.Success),
					Type:  "integer",
				},
			},
		}, nil, 500000)

		err := wavesClient.SignTx(&tx)
		if err != nil {
			return err
		}

		err = wavesClient.Broadcast(tx)
		if err != nil {
			return err
		}

		err = <-wavesClient.WaitTx(tx.ID)
		if err != nil {
			return err
		}

		fmt.Printf("Tx finilize: %s \n", tx.ID)
		request.Status = models.Success
		db.Save(request)
	}
	return nil
}

//TODO
func isWavesValid(request models.Request, tokenAddress string, assetId string, fromDecimals int, toDecimals int, db *gorm.DB) bool {
	/*var inputRequests []models.Request
	db.Where(&models.Request{Target: request.Owner, Amount: converter.StrConvert(request.Amount, fromDecimals, toDecimals), Type: models.BurOrLock,
		Owner: strings.ToLower(request.Target), Status: models.Success, AssetId: tokenAddress}).Where("chain_type = ?", models.Ethereum).Find(&inputRequests)

	var outSuccessRequests []models.Request
	db.Where(&models.Request{Target: strings.ToLower(request.Target), ChainType: models.Waves, Amount: request.Amount, Type: models.MintOrUnlock,
		Owner: request.Owner, Status: models.Success, AssetId: assetId}).Find(&outSuccessRequests)

	if len(inputRequests) <= len(outSuccessRequests) {
		return false
	}

	var outNewRequests []models.Request
	db.Where(&models.Request{Target: strings.ToLower(request.Target), ChainType: models.Waves, Amount: request.Amount, Type: models.MintOrUnlock,
		Owner: request.Owner, Status: models.New, AssetId: assetId}).Find(&outNewRequests)

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
