package types

type GetBookTickerForSymbolParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type GetBookTickerForSymbolsParam struct {
	Symbols []string `url:"symbols" validate:"required"`
}

type BookTickerParams struct {
	Symbol  string `url:"symbol,omitempty" validate:"omitempty"`
	Symbols string `url:"symbols,omitempty" validate:"omitempty"`
}

type BookTicker struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}
