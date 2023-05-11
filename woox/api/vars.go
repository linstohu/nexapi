package api

var DefaultContentType = map[string]string{
	"Content-Type": "application/json",
}

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
