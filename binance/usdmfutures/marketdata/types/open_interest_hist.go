package types

type GetOpenInterestHistParam struct {
	Symbol    string `url:"symbol" validate:"required"`
	Period    string `url:"period" validate:"required,oneof=5m 15m 30m 1h 2h 4h 6h 12h 1d"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=500"`
}

type OpenInterestHist struct {
	Symbol               string `json:"symbol"`
	SumOpenInterest      string `json:"sumOpenInterest"`
	SumOpenInterestValue string `json:"sumOpenInterestValue"`
	Timestamp            int64  `json:"timestamp"`
}
