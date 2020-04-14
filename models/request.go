package models

type Request struct {
	Id              string
	CreatedAt       int32
	Status          Status
	Amount          string
	Owner           string
	Target          string
	ChainType       ChainType
	Type            RqType
	AssetId         string
	TargetRequestId string
	Signs           []Signs
}
