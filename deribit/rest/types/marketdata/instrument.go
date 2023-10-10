package marketdata

type Instrument struct {
	TickSize            float64 `json:"tick_size"`
	MakerCommission     float64 `json:"maker_commission"`
	TakerCommission     float64 `json:"taker_commission"`
	Strike              float64 `json:"strike"`
	SettlementPeriod    string  `json:"settlement_period"`
	SettlementCurrency  string  `json:"settlement_currency"`
	QuoteCurrency       string  `json:"quote_currency"`
	PriceIndex          string  `json:"price_index"`
	OptionType          string  `json:"option_type"`
	MinTradeAmount      float64 `json:"min_trade_amount"`
	Kind                string  `json:"kind"`
	IsActive            bool    `json:"is_active"`
	InstrumentName      string  `json:"instrument_name"`
	ExpirationTimestamp int64   `json:"expiration_timestamp"`
	CreationTimestamp   int64   `json:"creation_timestamp"`
	ContractSize        float64 `json:"contract_size"`
	BaseCurrency        string  `json:"base_currency"`
}

type GetInstrumentParams struct {
	InstrumentName string `json:"instrument_name"`
}

type GetInstrumentsParams struct {
	Currency string `json:"currency"`
	Kind     string `json:"kind,omitempty"`
	Expired  bool   `json:"expired,omitempty"`
}
