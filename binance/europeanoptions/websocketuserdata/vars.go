package websocketuserdata

var (
	OptionsUserDataStreamBaseURL = "wss://nbstream.binance.com/eoptions"
	UserDataStreamRouter         = "/ws/"
)

const (
	MaxTryTimes = 5
)

const (
	SUBSCRIBE   = "SUBSCRIBE"
	UNSUBSCRIBE = "UNSUBSCRIBE"
)
