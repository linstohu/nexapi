package types

type GetMarketTradesParam struct {
	Symbol string `url:"symbol" validate:"required"`
	Limit  int    `url:"limit,omitempty"`
}

type MarketTrade struct {
	Response
	Rows []struct {
		Symbol            string  `json:"symbol"`
		Side              string  `json:"side"`
		ExecutedPrice     float64 `json:"executed_price"`
		ExecutedQuantity  float64 `json:"executed_quantity"`
		ExecutedTimestamp string  `json:"executed_timestamp"`
		Source            int     `json:"source"`
	} `json:"rows"`
}
