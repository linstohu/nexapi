package websocketuserdata

var (
	PortfolioMarginUserDataStreamBaseURL = "wss://fstream.binance.com/pm"
	UserDataStreamRouter                 = "/ws/"
)

const (
	MaxTryTimes = 5
)

const (
	SUBSCRIBE   = "SUBSCRIBE"
	UNSUBSCRIBE = "UNSUBSCRIBE"
)
