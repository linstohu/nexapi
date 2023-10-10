package trading

type CancelParams struct {
	OrderID string `json:"order_id"`
}

type CancelAllByInstrumentParams struct {
	InstrumentName string `json:"instrument_name"`
	Type           string `json:"type,omitempty"`
}
