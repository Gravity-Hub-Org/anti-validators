package server

import (
	"anti-validators/db/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

var validatorDb *gorm.DB
var oracleAddress string

func StartServer(host string, oracle string, db *gorm.DB) {
	validatorDb = db
	oracleAddress = oracle
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
	var sign models.Signs
	err := validatorDb.Where(&models.Signs{ValidatorAddress: oracleAddress, RequestId: reqeustHash}).Find(&sign).Error
	if err != nil {
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
