package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Waves          `json:"Waves"`
	Ethereum       `json:"Ethereum"`
	BftCoefficient int
	Timeout        int
	Ips            []string
}

type Waves struct {
	NodeURL         string
	ApiKey          string
	ContractAddress string
}

type Ethereum struct {
	NodeURL         string
	PrivateKey      string
	ContractAddress string
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
