package marketdata

type BookSummary struct {
	Volume            float64 `json:"volume"`
	UnderlyingPrice   float64 `json:"underlying_price"`
	UnderlyingIndex   string  `json:"underlying_index"`
	QuoteCurrency     string  `json:"quote_currency"`
	OpenInterest      float64 `json:"open_interest"`
	MidPrice          float64 `json:"mid_price"`
	MarkPrice         float64 `json:"mark_price"`
	Low               float64 `json:"low"`
	Last              float64 `json:"last"`
	InterestRate      float64 `json:"interest_rate"`
	InstrumentName    string  `json:"instrument_name"`
	High              float64 `json:"high"`
	CreationTimestamp int64   `json:"creation_timestamp"`
	BidPrice          float64 `json:"bid_price"`
	BaseCurrency      string  `json:"base_currency"`
	AskPrice          float64 `json:"ask_price"`
}

type GetBookSummaryByCurrencyParams struct {
	Currency string `json:"currency"`
	Kind     string `json:"kind,omitempty"`
}

type GetBookSummaryByInstrumentParams struct {
	InstrumentName string `json:"instrument_name"`
}
