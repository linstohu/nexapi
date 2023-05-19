package websocketmarket

var (
	MarketStreamBaseURL = "wss://data-stream.binance.com"
)

const CombinedStreamRouter = "/stream"

const (
	MaxTryTimes = 5
)

const (
	SUBSCRIBE   = "SUBSCRIBE"
	UNSUBSCRIBE = "UNSUBSCRIBE"
)
