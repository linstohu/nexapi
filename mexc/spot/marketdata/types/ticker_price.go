package types

type GetTickerPriceForSymbolParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type TickerPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
