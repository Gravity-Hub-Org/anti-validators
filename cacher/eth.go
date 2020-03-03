package cacher

import (
	"anti-validators/contracts"
	"anti-validators/db/models"
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/jinzhu/gorm"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/ethclient"
)

func StartEthCacher(contractAddress string, host string, ctx context.Context, db *gorm.DB) error {
	client, err := ethclient.DialContext(ctx, host)
	if err != nil {
		return err
	}

	ethContractAddress := common.Address{}
	hexAddress, err := hexutil.Decode(contractAddress)
	if err != nil {
		return err
	}
	ethContractAddress.SetBytes(hexAddress)
	contract, err := contracts.NewSupersymmetry(ethContractAddress, client)
	if err != nil {
		return err
	}
	go func() {
		for true {
			err := handleEthLogs(contract, ctx, db)
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

	return nil
}

func handleEthLogs(contract *contracts.Supersymmetry, ctx context.Context, db *gorm.DB) error {
	lastRequest := new(models.Request)
	if err := db.Order("CreatedAt").Last(&lastRequest).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	logs, err := contract.FilterNewRequest(&bind.FilterOpts{
		Start:   0,
		End:     nil,
		Context: ctx,
	})
	if err != nil {
		return err
	}

	for logs.Next() {
		hash := hexutil.Encode(logs.Event.RequestHash[:])

		request, err := contract.Requests(&bind.CallOpts{
			Context: ctx,
			Pending: false,
		}, logs.Event.RequestHash)

		if err != nil {
			return err
		}

		var dbRq models.Request
		db.Where(&models.Request{Id: hash}).First(&dbRq)
		if dbRq.Id != "" {
			status := models.Status(request.Status)
			if dbRq.Status != status {
				dbRq.Status = status
				if err := db.Save(dbRq).Error; err != nil {
					return err
				}
			}
			continue
		}

		if request.Status != uint8(models.New) {
			continue
		}

		req := models.Request{
			Id:        hexutil.Encode(logs.Event.RequestHash[:]),
			CreatedAt: int32(request.Height.Int64()),
			Status:    models.New,
			Amount:    request.TokenAmount.String(),
			Owner:     request.Owner.Hex(),
			Target:    request.Target,
			ChainType: models.Ethereum,
			Type:      models.RequestType(request.RType),
			Signs:     nil,
		}
		if err := db.Save(req).Error; err != nil {
			return err
		}
	}
	return nil
}
