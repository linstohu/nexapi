package types

type GetOrderbookParam struct {
	MaxLevel int `url:"max_level,omitempty"`
}

type Orderbook struct {
	Response
	Asks []struct {
		Price    float64 `json:"price"`
		Quantity float64 `json:"quantity"`
	} `json:"asks"`
	Bids []struct {
		Price    float64 `json:"price"`
		Quantity float64 `json:"quantity"`
	} `json:"bids"`
	Timestamp int64 `json:"timestamp"`
}
