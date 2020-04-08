package cacher

import (
	"anti-validators/db/models"
	"anti-validators/wavesapi"
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

func StartWavesCacher(contractAddress string, nodeUrl string, rqTimeout int, blockTimeout int, ctx context.Context, db *gorm.DB) {
	var nodeClient = wavesapi.New(nodeUrl, "")

	go func() {
		for true {
			err := handleWavesState(nodeClient, rqTimeout, blockTimeout, contractAddress, db)
			if err != nil {
				fmt.Printf("Error: %s \n", err.Error())
			}

			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(5 * time.Second)
			}
		}
	}()
}

func handleWavesState(nodeClient *wavesapi.Node, rqTimeout int, blockTimeout int, contractAddress string, db *gorm.DB) error {
	state, err := nodeClient.GetStateByAddress(contractAddress)
	if err != nil {
		return err
	}
	height, err := nodeClient.GetHeight()
	if err != nil {
		return err
	}

	lastElement := state["last_element"]
	orderId := lastElement.Value.(string)
	for {
		rqHeight := int(state["height_"+orderId].Value.(float64))
		if rqHeight > height-blockTimeout {
			continue
		}

		if height > rqHeight+rqTimeout {
			break
		}

		var dbRq models.Request
		db.Where(&models.Request{Id: orderId}).First(&dbRq)

		status := models.Status(state["status_"+orderId].Value.(float64))
		if dbRq.Id != "" {
			if dbRq.Status != status {
				dbRq.Status = status
				if err := db.Save(&dbRq).Error; err != nil {
					return err
				}
			}
			continue
		}

		strAmount := strconv.FormatFloat(state["amount_"+orderId].Value.(float64), 'f', 5, 64)
		bigAmount := big.NewInt(0)
		bigAmount.SetString(strAmount, 10)
		req := models.Request{
			Id:        orderId,
			CreatedAt: int32(rqHeight),
			Status:    status,
			Owner:     state["owner_"+orderId].Value.(string),
			Target:    strings.ToLower(state["target_"+orderId].Value.(string)),
			Amount:    bigAmount.String(),
			ChainType: models.Waves,
			Type:      models.RequestType(state["type_"+orderId].Value.(float64)),
			Signs:     nil,
			AssetId:   state["erc20_address_"+orderId].Value.(string),
		}
		if err := db.Save(req).Error; err != nil {
			return err
		}

		prevOrder := state["prev_rq_"+orderId]
		orderId = prevOrder.Value.(string)
	}
	db.Where("chain_type = ?", models.Waves).Where("create_at <= ?", height-rqTimeout).Update("status", models.Rejected)
	return nil
}
