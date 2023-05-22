package types

type GetOrderbookParams struct {
	Symbol string `url:"symbol" validate:"required"`
	Limit  int    `url:"limit,omitempty" validate:"omitempty,oneof=5 10 20 50 100 500 1000"`
}

type Orderbook struct {
	Symbol            string     `json:"symbol"`
	Pair              string     `json:"pair"`
	MessageOutputTime int64      `json:"E"`
	TransactionTime   int64      `json:"T"`
	LastUpdateID      int64      `json:"lastUpdateId"`
	Bids              [][]string `json:"bids"`
	Asks              [][]string `json:"asks"`
}
