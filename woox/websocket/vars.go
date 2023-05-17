package websocket

var (
	TestNetPublicBaseURL = "wss://wss.staging.woo.org/ws/stream/"
	PublicBaseURL        = "wss://wss.woo.org/ws/stream/"

	TestNetPrivateBaseURL = "wss://wss.staging.woo.org/v2/ws/private/stream/"
	PrivateBaseURL        = "wss://wss.woo.org/v2/ws/private/stream/"
)

const (
	MaxTryTimes = 5
)

const (
	SUBSCRIBE   = "subscribe"
	UNSUBSCRIBE = "unsubscribe"
)
