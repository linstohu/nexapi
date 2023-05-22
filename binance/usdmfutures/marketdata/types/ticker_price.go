package types

type GetTickerPriceParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type TickerPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Time   int64  `json:"time"`
}
