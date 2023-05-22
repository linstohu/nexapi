package types

import "github.com/linstohu/nexapi/binance/usdmfutures/utils"

type GetKlineParam struct {
	Symbol    string              `url:"symbol" validate:"required"`
	Interval  utils.KlineInterval `url:"interval" validate:"required,oneof=1m 3m 5m 15m 30m 1h 2h 4h 6h 8h 12h 1d 3d 1w 1M"`
	StartTime int64               `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64               `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int                 `url:"limit,omitempty" validate:"omitempty,max=1500"`
}
