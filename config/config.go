package config

import (
	"anti-validators/db/models"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Host            string
	Db              string
	Waves           `json:"Waves"`
	Ethereum        `json:"Ethereum"`
	BftCoefficient  int
	Timeout         int
	RqTimeout       int
	Ips             []string
	ApproveNewToken []ApproveRqToken
}

type ApproveRqToken struct {
	ChainType models.ChainType
	TokenId   string
}

type Waves struct {
	NodeURL         string
	ApiKey          string
	ContractAddress string
	BlockInterval   int
	OracleAddress   string
}

type Ethereum struct {
	NodeURL         string
	PrivateKey      string
	ContractAddress string
	BlockInterval   int
	OracleAddress   string
}

func Load(filename string) (Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	if err := json.Unmarshal(file, &config); err != nil {
		return Config{}, err
	}
	return config, err
}
