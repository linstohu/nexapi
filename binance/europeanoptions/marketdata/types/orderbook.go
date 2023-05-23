package types

type GetOrderbookParams struct {
	Symbol string `url:"symbol" validate:"required"`
	Limit  int    `url:"limit,omitempty" validate:"omitempty,oneof=10 20 50 100 500 1000"`
}

type Orderbook struct {
	TransactionTime int64      `json:"T"`
	UpdateID        int64      `json:"u"`
	Bids            [][]string `json:"bids"`
	Asks            [][]string `json:"asks"`
}
