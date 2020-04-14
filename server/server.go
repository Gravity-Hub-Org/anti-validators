package server

import (
	"anti-validators/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

var validatorDb *gorm.DB
var wavesOracleAddress, ethOracleAddress string

func StartServer(host string, newWavesOracleAddress string, newEthOracleAddress string, db *gorm.DB) {
	validatorDb = db
	wavesOracleAddress = newWavesOracleAddress
	ethOracleAddress = newEthOracleAddress
	for {
		http.HandleFunc("/api/signs/", handleGetSign)
		http.ListenAndServe(host, nil)
	}
}

func handleGetSign(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["requestHash"]

	if !ok || len(keys[0]) < 1 {
		http.Error(w, "Url Param 'requestHash' is missing", 404)
		return
	}

	reqeustHash := keys[0]

	var rq models.Request

	if err := validatorDb.Where(&models.Request{Id: reqeustHash}).First(&rq).Error; err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	oracle := ""
	if rq.ChainType == models.Ethereum {
		oracle = ethOracleAddress
	} else {
		oracle = wavesOracleAddress
	}

	var sign models.Signs
	if err := validatorDb.Where(&models.Signs{ValidatorAddress: oracle, RequestId: reqeustHash}).Find(&sign).Error; err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if sign.RequestId != reqeustHash {
		http.Error(w, "not found", 400)
		return
	}
	result, err := json.Marshal(sign)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprintf(w, string(result))
}
