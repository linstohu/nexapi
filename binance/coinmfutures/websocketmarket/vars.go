package websocketmarket

var (
	CoinMarginedMarketStreamBaseURL = "wss://dstream.binance.com"
	CombinedStreamRouter            = "/stream"
)

const (
	MaxTryTimes = 5
)

const (
	SUBSCRIBE   = "SUBSCRIBE"
	UNSUBSCRIBE = "UNSUBSCRIBE"
)
