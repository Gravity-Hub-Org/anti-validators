package signer

import (
	"anti-validators/models"
	"anti-validators/wavesapi"
	"anti-validators/wavesapi/transactions"
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
)

func formatWavesMsg(requestId string, status models.Status) string {
	return requestId + "_" + strconv.Itoa(int(status))
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
