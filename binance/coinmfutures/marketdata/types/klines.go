package types

type Kline struct {
	OpenTime                int64  `json:"openTime"`
	OpenPrice               string `json:"openPrice"`
	HighPrice               string `json:"highPrice"`
	LowPrice                string `json:"lowPrice"`
	ClosePrice              string `json:"closePrice"`
	Volume                  string `json:"volume"`
	CloseTime               int64  `json:"closeTime"`
	BaseAssetVolume         string `json:"baseAssetVolume"`
	NumberOfTrades          int64  `json:"numberOfTrades"`
	TakerBuyVolume          string `json:"takerBuyVolume"`
	TakerBuyBaseAssetVolume string `json:"takerBuyBaseAssetVolume"`
}
