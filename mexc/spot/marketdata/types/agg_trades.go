package types

type GetAggTradesParam struct {
	Symbol    string `url:"symbol" validate:"required"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type AggTrade struct {
	A  int    `json:"a"` // Aggregate tradeId
	F  int    `json:"f"` // First tradeId
	L  int    `json:"l"` // Last tradeId
	P  string `json:"p"` // Price
	Q  string `json:"q"` // Quantity
	T  int64  `json:"T"` // Timestamp
	M  bool   `json:"m"` // Was the buyer the maker?
	Ma bool   `json:"M"` // Was the trade the best price match?
}
