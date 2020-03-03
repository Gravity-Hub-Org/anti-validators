package cacher

import (
	"anti-validators/db/models"
	"anti-validators/wavesapi"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

func StartWavesCacher(contractAddress string, nodeUrl string, ctx context.Context, db *gorm.DB) {
	var nodeClient = wavesapi.New(nodeUrl, "")

	go func() {
		for true {
			err := handleWavesState(nodeClient, contractAddress, db)
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

func handleWavesState(nodeClient *wavesapi.Node, contractAddress string, db *gorm.DB) error {
	state, err := nodeClient.GetStateByAddress(contractAddress)
	if err != nil {
		return err
	}

	for key, value := range state {
		if !strings.HasPrefix(key, "status_") {
			continue
		}

		id := strings.Split(key, "_")[1]

		var dbRq models.Request
		db.Where(&models.Request{Id: id}).First(&dbRq)
		if dbRq.Id != "" {
			status := models.Status(state["status_"+id].Value.(float64))
			if dbRq.Status != status {
				dbRq.Status = status
				if err := db.Save(&dbRq).Error; err != nil {
					return err
				}
			}
			continue
		}

		if value.Value.(float64) != float64(models.New) {
			continue
		}
		strAmount := strconv.FormatFloat(state["amount_"+id].Value.(float64), 'f', 5, 64)
		req := models.Request{
			Id:        id,
			CreatedAt: int32(state["height_"+id].Value.(float64)),
			Status:    models.New,
			Owner:     state["owner_"+id].Value.(string),
			Target:    state["target_"+id].Value.(string),
			Amount:    strAmount,
			ChainType: models.Waves,
			Type:      models.RequestType(state["type_"+id].Value.(float64)),
			Signs:     nil,
		}
		if err := db.Save(req).Error; err != nil {
			return err
		}
	}
	return nil
}
