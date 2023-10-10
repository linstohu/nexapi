package marketdata

type Trade struct {
	TradeSeq       int     `json:"trade_seq"`
	TradeID        string  `json:"trade_id"`
	Timestamp      int64   `json:"timestamp"`
	TickDirection  int     `json:"tick_direction"`
	Price          float64 `json:"price"`
	MarkPrice      float64 `json:"mark_price"`
	InstrumentName string  `json:"instrument_name"`
	IndexPrice     float64 `json:"index_price"`
	Direction      string  `json:"direction"`
	Amount         float64 `json:"amount"`
}

type GetLastTradesResponse struct {
	Trades  []Trade `json:"trades"`
	HasMore bool    `json:"has_more"`
}

type GetLastTradesByInstrumentAndTimeParams struct {
	InstrumentName string `json:"instrument_name"`
	StartTimestamp int    `json:"start_timestamp"`
	EndTimestamp   int    `json:"end_timestamp"`
	Count          int    `json:"count,omitempty"`
	IncludeOld     bool   `json:"include_old,omitempty"`
	Sorting        string `json:"sorting,omitempty"`
}
