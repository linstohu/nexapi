package types

type GetTradeParams struct {
	Symbol string `url:"symbol" validate:"required"`
	Limit  int    `url:"limit,omitempty" validate:"omitempty,max=500"`
}

type Trade struct {
	ID       int64  `json:"id"`
	Symbol   string `json:"symbol"`
	Price    string `json:"price"`
	Qty      string `json:"qty"`
	QuoteQty string `json:"quoteQty"`
	Side     int    `json:"side"` // Completed trade direction（-1 Sell，1 Buy）
	Time     int64  `json:"time"`
}
