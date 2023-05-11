package types

type GetSymbolInfoParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type SymbolInfo struct {
	Response
	Info struct {
		Symbol      string  `json:"symbol"`
		QuoteMin    float64 `json:"quote_min"`
		QuoteMax    float64 `json:"quote_max"`
		QuoteTick   float64 `json:"quote_tick"`
		BaseMin     float64 `json:"base_min"`
		BaseMax     float64 `json:"base_max"`
		BaseTick    float64 `json:"base_tick"`
		MinNotional float64 `json:"min_notional"`
		PriceRange  float64 `json:"price_range"`
		PriceScope  float64 `json:"price_scope"`
		CreatedTime string  `json:"created_time"`
		UpdatedTime string  `json:"updated_time"`
		IsStable    int     `json:"is_stable"`
	} `json:"info"`
}

type AvailableSymbols struct {
	Response
	Rows []struct {
		Symbol      string  `json:"symbol"`
		QuoteMin    float64 `json:"quote_min"`
		QuoteMax    float64 `json:"quote_max"`
		QuoteTick   float64 `json:"quote_tick"`
		BaseMin     float64 `json:"base_min"`
		BaseMax     float64 `json:"base_max"`
		BaseTick    float64 `json:"base_tick"`
		MinNotional float64 `json:"min_notional"`
		PriceRange  float64 `json:"price_range"`
		PriceScope  float64 `json:"price_scope"`
		CreatedTime string  `json:"created_time"`
		UpdatedTime string  `json:"updated_time"`
		IsStable    int     `json:"is_stable"`
	} `json:"rows"`
}
