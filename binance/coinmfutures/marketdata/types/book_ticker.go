package types

type GetBookTickerParam struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
	Pair   string `url:"pair,omitempty" validate:"omitempty"`
}

type BookTicker struct {
	Symbol   string `json:"symbol"`
	Pair     string `json:"pair"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
	Time     int64  `json:"time"`
}
