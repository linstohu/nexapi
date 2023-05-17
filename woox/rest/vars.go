package rest

const (
	TestNetBaseURL = "https://api.staging.woo.org"
	BaseURL        = "https://api.woo.org"
)

var (
	V1DefaultContentType = map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	V3DefaultContentType = map[string]string{
		"Content-Type": "application/json",
	}
)

const (
	LimitOrderType    = "LIMIT"
	MarketOrderType   = "MARKET"
	IocOrderType      = "IOC"
	FokOrderType      = "FOK"
	PostOnlyOrderType = "POST_ONLY"
	AskOrderType      = "ASK"
	BidOrderType      = "BID"
)

const (
	SELL = "SELL"
	BUY  = "BUY"
)
