package types

type FuturesInfo struct {
	Symbol          string  `json:"symbol"`
	IndexPrice      float64 `json:"index_price"`
	MarkPrice       float64 `json:"mark_price"`
	EstFundingRate  float64 `json:"est_funding_rate"`
	LastFundingRate float64 `json:"last_funding_rate"`
	NextFundingTime int64   `json:"next_funding_time"`
	OpenInterest    float64 `json:"open_interest"`
	Two4HOpen       float64 `json:"24h_open"`
	Two4HClose      float64 `json:"24h_close"`
	Two4HHigh       float64 `json:"24h_high"`
	Two4HLow        float64 `json:"24h_low"`
	Two4HVolume     float64 `json:"24h_volume"`
	Two4HAmount     float64 `json:"24h_amount"`
}

type AllFuturesInfo struct {
	Success   bool          `json:"success"`
	Rows      []FuturesInfo `json:"rows"`
	Timestamp int64         `json:"timestamp"`
}

type OneFuturesInfo struct {
	Response
	Info      FuturesInfo `json:"info"`
	Timestamp int64       `json:"timestamp"`
}
