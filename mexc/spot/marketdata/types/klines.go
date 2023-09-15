package types

import (
	spotutils "github.com/linstohu/nexapi/mexc/spot/utils"
)

type GetKlineParam struct {
	Symbol    string                  `url:"symbol" validate:"required"`
	Interval  spotutils.KlineInterval `url:"interval" validate:"required,oneof=1m 5m 15m 30m 60m 4h 1d 1M"`
	StartTime int64                   `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64                   `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int                     `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type Kline struct {
	OpenTime         int64  `json:"openTime"`
	OpenPrice        string `json:"openPrice"`
	HighPrice        string `json:"highPrice"`
	LowPrice         string `json:"lowPrice"`
	ClosePrice       string `json:"closePrice"`
	Volume           string `json:"volume"`
	CloseTime        int64  `json:"closeTime"`
	QuoteAssetVolume string `json:"quoteAssetVolume"`
}
