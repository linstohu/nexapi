package types

type ExchangeInfo struct {
	Timezone        string `json:"timezone"`
	ServerTime      int64  `json:"serverTime"`
	OptionContracts []struct {
		ID          int    `json:"id"`
		BaseAsset   string `json:"baseAsset"`
		QuoteAsset  string `json:"quoteAsset"`
		Underlying  string `json:"underlying"`
		SettleAsset string `json:"settleAsset"`
	} `json:"optionContracts"`
	OptionAssets []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"optionAssets"`
	OptionSymbols []struct {
		ContractID int64 `json:"contractId"`
		ExpiryDate int64 `json:"expiryDate"`
		Filters    []struct {
			FilterType string `json:"filterType"`
			MinPrice   string `json:"minPrice,omitempty"`
			MaxPrice   string `json:"maxPrice,omitempty"`
			TickSize   string `json:"tickSize,omitempty"`
			MinQty     string `json:"minQty,omitempty"`
			MaxQty     string `json:"maxQty,omitempty"`
			StepSize   string `json:"stepSize,omitempty"`
		} `json:"filters"`
		ID                   int64  `json:"id"`
		Symbol               string `json:"symbol"`
		Side                 string `json:"side"`
		StrikePrice          string `json:"strikePrice"`
		Underlying           string `json:"underlying"`
		Unit                 int64  `json:"unit"`
		MakerFeeRate         string `json:"makerFeeRate"`
		TakerFeeRate         string `json:"takerFeeRate"`
		MinQty               string `json:"minQty"`
		MaxQty               string `json:"maxQty"`
		InitialMargin        string `json:"initialMargin"`
		MaintenanceMargin    string `json:"maintenanceMargin"`
		MinInitialMargin     string `json:"minInitialMargin"`
		MinMaintenanceMargin string `json:"minMaintenanceMargin"`
		PriceScale           int    `json:"priceScale"`
		QuantityScale        int    `json:"quantityScale"`
		QuoteAsset           string `json:"quoteAsset"`
	} `json:"optionSymbols"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
	} `json:"rateLimits"`
}
