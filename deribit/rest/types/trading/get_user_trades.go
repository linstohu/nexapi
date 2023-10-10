package trading

type GetUserTradesResponse struct {
	Trades  []Trade `json:"trades"`
	HasMore bool    `json:"has_more"`
}

type GetUserTradesByCurrencyParams struct {
	Currency string `json:"currency"`
	Kind     string `json:"kind,omitempty"`
	StartId  string `json:"start_id,omitempty"`
	EndId    string `json:"end_id,omitempty"`
	Count    int64  `json:"count,omitempty"`
	Sorting  string `json:"sorting,omitempty"`
}

type GetUserTradesByCurrencyAndTimeParams struct {
	Currency       string `json:"currency"`
	Kind           string `json:"kind,omitempty"`
	StartTimestamp int    `json:"start_timestamp"`
	EndTimestamp   int    `json:"end_timestamp"`
	Count          int    `json:"count,omitempty"`
	Sorting        string `json:"sorting,omitempty"`
}

type GetUserTradesByInstrumentParams struct {
	InstrumentName string `json:"instrument_name"`
	StartSeq       int    `json:"start_seq,omitempty"`
	EndSeq         int    `json:"end_seq,omitempty"`
	Count          int    `json:"count,omitempty"`
	Sorting        string `json:"sorting,omitempty"`
}

type GetUserTradesByInstrumentAndTimeParams struct {
	InstrumentName string `json:"instrument_name"`
	StartTimestamp int    `json:"start_timestamp"`
	EndTimestamp   int    `json:"end_timestamp"`
	Count          int    `json:"count,omitempty"`
	Sorting        string `json:"sorting,omitempty"`
}
