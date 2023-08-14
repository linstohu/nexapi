package types

type Position struct {
	Symbol           string  `json:"symbol"`
	Holding          float64 `json:"holding"`
	PendingLongQty   float64 `json:"pending_long_qty"`
	PendingShortQty  float64 `json:"pending_short_qty"`
	SettlePrice      float64 `json:"settle_price"`
	AverageOpenPrice float64 `json:"average_open_price"`
	Pnl24H           float64 `json:"pnl_24_h"`
	Fee24H           float64 `json:"fee_24_h"`
	MarkPrice        float64 `json:"mark_price"`
	EstLiqPrice      float64 `json:"est_liq_price"`
	Timestamp        int64  `json:"timestamp"`
}

type GetOnePositionInfo struct {
	Response
	Position
}

type GetAllPositionInfo struct {
	Response
	Data struct {
		Positions []Position `json:"positions"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}
