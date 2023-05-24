package types

type GetPriceTickerParam struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
	Pair   string `url:"pair,omitempty" validate:"omitempty"`
}

type PriceTicker struct {
	Symbol string `json:"symbol"`
	Pair   string `json:"pair"`
	Price  string `json:"price"`
	Time   int64  `json:"time"`
}
