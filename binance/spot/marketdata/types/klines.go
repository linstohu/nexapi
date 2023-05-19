package types

type GetKlineParam struct {
	Symbol    string        `url:"symbol" validate:"required"`
	Interval  KlineInterval `url:"interval" validate:"required,oneof=1s 1m 3m 5m 15m 30m 1h 2h 4h 6h 8h 12h 1d 3d 1w 1M"`
	StartTime int64         `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64         `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int           `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type KlineInterval string

var (
	Second1  KlineInterval = "1s"
	Minute1  KlineInterval = "1m"
	Minute3  KlineInterval = "3m"
	Minute5  KlineInterval = "5m"
	Minute15 KlineInterval = "15m"
	Minute30 KlineInterval = "30m"
	Hour1    KlineInterval = "1h"
	Hour2    KlineInterval = "2h"
	Hour4    KlineInterval = "4h"
	Hour6    KlineInterval = "6h"
	Hour8    KlineInterval = "8h"
	Hour12   KlineInterval = "12h"
	Day1     KlineInterval = "1d"
	Day3     KlineInterval = "3d"
	Week1    KlineInterval = "1w"
	Month1   KlineInterval = "1M"
)

type Kline struct {
	OpenTime                 int64  `json:"openTime"`
	OpenPrice                string `json:"openPrice"`
	HighPrice                string `json:"highPrice"`
	LowPrice                 string `json:"lowPrice"`
	ClosePrice               string `json:"closePrice"`
	Volume                   string `json:"volume"`
	CloseTime                int64  `json:"closeTime"`
	QuoteAssetVolume         string `json:"quoteAssetVolume"`
	NumberOfTrades           int64  `json:"numberOfTrades"`
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
}
