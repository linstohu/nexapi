package types

type GetTradeFeeParam struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
}

type TradeFeeParams struct {
	GetTradeFeeParam
	DefaultParam
}

type TradeFee struct {
	Symbol          string `json:"symbol"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}
