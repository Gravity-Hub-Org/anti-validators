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
			err := signWavesRequest(wavesClient, oracleAddress, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}

			err = getWavesRequest(cfg.Ips, oracleAddress, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}

			err = finalizeWavesRequest(wavesClient, cfg.Waves.ContractAddress, pubKeyOracles, oracleAddress, cfg.BftCoefficient, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}

			time.Sleep(10 * time.Second)
		}
	}()

	return nil
}

func signWavesRequest(wavesClient *wavesapi.Node, oracleAddress string, db *gorm.DB) error {
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
		msg := request.Id + "_" + strconv.Itoa(int(models.Success))
		signedText, err := wavesClient.SignMsg(msg, oracleAddress)
		if err != nil {
			return err
		}

		if err := db.Save(models.Signs{
			RequestId:        request.Id,
			ValidatorAddress: oracleAddress,
			Sign:             signedText.Signature,
			ValidatorPubKey:  signedText.PublicKey,
			CreatedAt:        time.Now(),
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func getWavesRequest(ips []string, oracleAddress string, db *gorm.DB) error {
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

func finalizeWavesRequest(wavesClient *wavesapi.Node, contract string, pubKeyOracles []string, oracleAddress string, bftCoefficient int, db *gorm.DB) error {
	var requests []models.Request
	if err := db.Where(&models.Request{Status: models.New}).Find(&requests).Error; err != nil {
		return err
	}

	for _, request := range requests {
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
	}
	return nil
}
