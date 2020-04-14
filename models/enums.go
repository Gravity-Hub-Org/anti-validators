package models

type ChainType int
const (
	Waves ChainType = iota
	Ethereum
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