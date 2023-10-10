package trading

type GetOrderStateParams struct {
	OrderID string `json:"order_id"`
}

type GetOpenOrdersByCurrencyParams struct {
	Currency string `json:"currency"`
	Kind     string `json:"kind,omitempty"`
	Type     string `json:"type,omitempty"`
}

type GetOpenOrdersByInstrumentParams struct {
	InstrumentName string `json:"instrument_name"`
	Type           string `json:"type,omitempty"`
}
