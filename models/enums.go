package models

type ChainType int
const (
	Ethereum ChainType = iota
	Waves
)

type Status int
const (
	None Status = iota
	New
	Rejected
	Success
	Returned
)

type RqType int
const (
	Lock RqType = iota
	Unlock
	Mint
	Burn
)