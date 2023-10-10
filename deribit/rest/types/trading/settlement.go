package trading

type Settlement struct {
	IndexPrice        float64 `json:"index_price"`
	InstrumentName    string  `json:"instrument_name"`
	MarkPrice         float64 `json:"mark_price"`
	Position          float64 `json:"position"`
	ProfitLoss        float64 `json:"profit_loss"`
	SessionBankrupcy  float64 `json:"session_bankrupcy"`
	SessionProfitLoss float64 `json:"session_profit_loss"`
	SessionTax        float64 `json:"session_tax"`
	SessionTaxRate    float64 `json:"session_tax_rate"`
	Socialized        float64 `json:"socialized"`
	Timestamp         int64   `json:"timestamp"`
	Type              string  `json:"type"`
}

type GetSettlementHistoryResponse struct {
	Settlements  []Settlement `json:"settlements"`
	Continuation string       `json:"continuation"`
}

type GetSettlementHistoryByInstrumentParams struct {
	InstrumentName string `json:"instrument_name"`
	Type           string `json:"type,omitempty"`
	Count          int    `json:"count,omitempty"`
	Continuation   string `json:"continuation,omitempty"`
}

type GetSettlementHistoryByCurrencyParams struct {
	Currency     string `json:"currency"`
	Type         string `json:"type,omitempty"`
	Count        int    `json:"count,omitempty"`
	Continuation string `json:"continuation,omitempty"`
}
