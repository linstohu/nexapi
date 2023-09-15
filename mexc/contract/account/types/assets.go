package types

type GetAccountAsset struct {
	Response
	Data []*ContractAsset `json:"data"`
}

type ContractAsset struct {
	Currency         string  `json:"currency"`
	PositionMargin   float64 `json:"positionMargin"`
	AvailableBalance float64 `json:"availableBalance"`
	CashBalance      float64 `json:"cachBalance"`
	FrozenBalance    float64 `json:"frozenBalance"`
	Equity           float64 `json:"equity"`
	Unrealized       float64 `json:"unrealized"`
	Bonus            float64 `json:"bonus"`
}
