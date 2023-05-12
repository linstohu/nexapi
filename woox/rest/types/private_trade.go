package types

type Trade struct {
	ID                int     `json:"id"`
	Symbol            string  `json:"symbol"`
	OrderID           int     `json:"order_id"`
	OrderTag          string  `json:"order_tag"`
	ExecutedPrice     float64 `json:"executed_price"`
	ExecutedQuantity  float64 `json:"executed_quantity"`
	IsMaker           int     `json:"is_maker"`
	Side              string  `json:"side"`
	Fee               float64 `json:"fee"`
	FeeAsset          string  `json:"fee_asset"`
	ExecutedTimestamp string  `json:"executed_timestamp"`
}

type GetTrade struct {
	Response
	Trade
}

type GetTrades struct {
	Response
	Rows []Trade `json:"rows"`
}

type GetTradeHistoryParam struct {
	Symbol    string `url:"symbol,omitempty" json:"symbol,omitempty"`
	OrderTag  string `url:"order_tag,omitempty" json:"order_tag,omitempty" validate:"omitempty"`
	StartTime int64  `url:"start_t,omitempty" json:"start_t,omitempty" validate:"omitempty,gt=999999999999"`
	EndTime   int64  `url:"end_t,omitempty" json:"end_t,omitempty" validate:"omitempty,gt=999999999999"`
	Page      int64  `url:"page,omitempty" json:"page,omitempty" validate:"omitempty"`
	Size      int64  `url:"size,omitempty" json:"size,omitempty" validate:"omitempty"`
}

type GetTradeHistory struct {
	Response
	Meta struct {
		Total          int `json:"total"`
		RecordsPerPage int `json:"records_per_page"`
		CurrentPage    int `json:"current_page"`
	} `json:"meta"`
	Rows []Trade `json:"rows"`
}
