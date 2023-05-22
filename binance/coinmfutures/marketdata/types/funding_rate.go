package types

type GetFundingRateParam struct {
	Symbol    string `url:"symbol" validate:"required"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type FundingRate struct {
	Symbol      string `json:"symbol"`
	FundingRate string `json:"fundingRate"`
	FundingTime int64  `json:"fundingTime"`
}
