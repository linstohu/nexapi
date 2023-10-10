package marketdata

type GetTradingviewChartDataParams struct {
	InstrumentName string `json:"instrument_name"`
	StartTimestamp int64  `json:"start_timestamp"`
	EndTimestamp   int64  `json:"end_timestamp"`
	Resolution     string `json:"resolution"`
}

type GetTradingviewChartDataResponse struct {
	Close  []float64 `json:"close"`
	Cost   []float64 `json:"cost"`
	High   []float64 `json:"high"`
	Low    []float64 `json:"low"`
	Open   []float64 `json:"open"`
	Volume []float64 `json:"volume"`
	Ticks  []int64   `json:"ticks"`
	Status string    `json:"status"`
}
