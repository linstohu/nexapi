package types

type GetTickerPriceParam struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
	Pair   string `url:"pair,omitempty" validate:"omitempty"`
}

type TickerPrice struct {
	Symbol string `json:"symbol"`
	Pair   string `json:"pair"`
	Price  string `json:"price"`
	Time   int64  `json:"time"`
}
