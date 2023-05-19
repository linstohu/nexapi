package types

type GetTickerPriceForSymbolParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type GetTickerPriceForSymbolsParam struct {
	Symbols []string `url:"symbols" validate:"required"`
}

type TickerPriceParams struct {
	Symbol  string `url:"symbol,omitempty" validate:"omitempty"`
	Symbols string `url:"symbols,omitempty" validate:"omitempty"`
}

type TickerPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
