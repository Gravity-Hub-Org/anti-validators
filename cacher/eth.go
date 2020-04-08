package cacher

import (
	"anti-validators/contracts"
	"anti-validators/db/models"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/jinzhu/gorm"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/ethclient"
)

func StartEthCacher(contractAddress string, host string, rqTimeout int64, blockTimeout int64, ctx context.Context, db *gorm.DB) error {
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
			err := handleEthLogs(contract, client, rqTimeout, blockTimeout, ctx, db)
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

func handleEthLogs(contract *contracts.Supersymmetry, client *ethclient.Client, rqTimeout int64, blockTimeout int64, ctx context.Context, db *gorm.DB) error {
	lastRequest := new(models.Request)
	if err := db.Order("CreatedAt").Last(&lastRequest).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	height, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
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

		if request.Height.Int64() > height.Number.Int64()-blockTimeout {
			continue
		}

		if height.Number.Int64() > request.Height.Int64()+rqTimeout {
			break
		}

		var dbRq models.Request
		db.Where(&models.Request{Id: hash}).First(&dbRq)
		status := models.Status(request.Status)
		if dbRq.Id != "" {
			if dbRq.Status != status {
				dbRq.Status = status
				if err := db.Save(dbRq).Error; err != nil {
					return err
				}
			}
			continue
		}

		req := models.Request{
			Id:        hexutil.Encode(logs.Event.RequestHash[:]),
			CreatedAt: int32(request.Height.Int64()),
			Status:    status,
			Amount:    request.TokenAmount.String(),
			Owner:     strings.ToLower(request.Owner.Hex()),
			Target:    request.Target,
			ChainType: models.Ethereum,
			Type:      models.RequestType(request.RType),
			AssetId:   request.TokenAddress.String(),
			Signs:     nil,
		}
		if err := db.Save(req).Error; err != nil {
			return err
		}
	}
	db.Where("chain_type = ?", models.Ethereum).Where("create_at <= ?", height.Number.Int64()-rqTimeout).Update("status", models.Rejected)
	return nil
}
