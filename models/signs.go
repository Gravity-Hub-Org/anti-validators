package models

import "time"

type Signs struct {
	RequestId        string
	ValidatorAddress string
	Sign             string
	ValidatorPubKey  string
	CreatedAt        time.Time
	Status           Status
}
