package types

type GetOrderbookParams struct {
	Symbol string `url:"symbol" validate:"required"`
	Limit  int    `url:"limit,omitempty" validate:"omitempty,max=5000"`
}

type Orderbook struct {
	LastUpdateID int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}
