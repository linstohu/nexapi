package types

type ExchangeInfo struct {
	Timezone   string `json:"timezone"`
	ServerTime int64  `json:"serverTime"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
	} `json:"rateLimits"`
	Symbols []struct {
		Symbol                string   `json:"symbol"`
		Pair                  string   `json:"pair"`
		ContractType          string   `json:"contractType"`
		DeliveryDate          int64    `json:"deliveryDate"`
		OnboardDate           int64    `json:"onboardDate"`
		ContractStatus        string   `json:"contractStatus"`
		ContractSize          int      `json:"contractSize"`
		MarginAsset           string   `json:"marginAsset"`
		MaintMarginPercent    string   `json:"maintMarginPercent"`
		RequiredMarginPercent string   `json:"requiredMarginPercent"`
		BaseAsset             string   `json:"baseAsset"`
		QuoteAsset            string   `json:"quoteAsset"`
		PricePrecision        int      `json:"pricePrecision"`
		QuantityPrecision     int      `json:"quantityPrecision"`
		BaseAssetPrecision    int      `json:"baseAssetPrecision"`
		QuotePrecision        int      `json:"quotePrecision"`
		EqualQtyPrecision     int      `json:"equalQtyPrecision"`
		MaxMoveOrderLimit     int      `json:"maxMoveOrderLimit"`
		TriggerProtect        string   `json:"triggerProtect"`
		UnderlyingType        string   `json:"underlyingType"`
		UnderlyingSubType     []string `json:"underlyingSubType"`
		Filters               []struct {
			MinPrice          string `json:"minPrice,omitempty"`
			MaxPrice          string `json:"maxPrice,omitempty"`
			FilterType        string `json:"filterType"`
			TickSize          string `json:"tickSize,omitempty"`
			StepSize          string `json:"stepSize,omitempty"`
			MaxQty            string `json:"maxQty,omitempty"`
			MinQty            string `json:"minQty,omitempty"`
			Limit             int    `json:"limit,omitempty"`
			MultiplierDown    string `json:"multiplierDown,omitempty"`
			MultiplierUp      string `json:"multiplierUp,omitempty"`
			MultiplierDecimal string `json:"multiplierDecimal,omitempty"`
		} `json:"filters"`
		OrderTypes      []string `json:"orderTypes"`
		TimeInForce     []string `json:"timeInForce"`
		LiquidationFee  string   `json:"liquidationFee"`
		MarketTakeBound string   `json:"marketTakeBound"`
	} `json:"symbols"`
}
