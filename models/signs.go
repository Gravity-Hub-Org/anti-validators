package models

import "time"

type Signs struct {
	RequestId        string
	ValidatorAddress string
	ConfirmSign      string
	Sign             string
	ValidatorPubKey  string
	CreatedAt        time.Time
	Status           Status
}
