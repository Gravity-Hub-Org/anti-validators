package signer

import (
	"anti-validators/contracts"
	"anti-validators/models"
	"crypto/ecdsa"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/jinzhu/gorm"
)

func formatEthMsg(requestId string, status models.Status) ([]byte, error) {
	msg, err := hexutil.Decode(requestId)
	if err != nil {
		return nil, err
	}
	msg = append(msg, byte(status))

	return msg, nil
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
		if err := db.Where(&models.Signs{RequestId: request.Id, Status: status}).Find(&signs).Error; err != nil {
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

func signEthMsg(message []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	validationMsg := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message))
	validationHash := crypto.Keccak256(append([]byte(validationMsg), message[:]...))
	return crypto.Sign(validationHash, privKey)
}
