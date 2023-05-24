package websocketmarket

var (
	OptionsMarketStreamBaseURL = "wss://nbstream.binance.com/eoptions"
	CombinedStreamRouter       = "/stream"
)

const (
	MaxTryTimes = 5
)

const (
	SUBSCRIBE   = "SUBSCRIBE"
	UNSUBSCRIBE = "UNSUBSCRIBE"
)
