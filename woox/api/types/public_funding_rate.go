package types

type FundingRates struct {
	Success   bool                `json:"success"`
	Rows      []SymbolFundingRate `json:"rows"`
	Timestamp int64               `json:"timestamp"`
}

type FundingRate struct {
	Response
	SymbolFundingRate
	Timestamp int64 `json:"timestamp"`
}

type SymbolFundingRate struct {
	Symbol                   string  `json:"symbol"`
	EstFundingRate           float64 `json:"est_funding_rate"`
	EstFundingRateTimestamp  int64   `json:"est_funding_rate_timestamp"`
	LastFundingRate          float64 `json:"last_funding_rate"`
	LastFundingRateTimestamp int64   `json:"last_funding_rate_timestamp"`
	NextFundingTime          int64   `json:"next_funding_time"`
	LastFundingRateInterval  int     `json:"last_funding_rate_interval"`
	EstFundingRateInterval   int     `json:"est_funding_rate_interval"`
}
