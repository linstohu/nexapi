package websocketmarket

var (
	SpotMarketStreamBaseURL = "wss://data-stream.binance.vision"
	CombinedStreamRouter    = "/stream"
)

const (
	MaxTryTimes = 5
)

const (
	SUBSCRIBE   = "SUBSCRIBE"
	UNSUBSCRIBE = "UNSUBSCRIBE"
)
