package websocket

var (
	PublicProdBaseEndpoint  = "wss://wss.woo.org/ws/stream/"
	PrivateProdBaseEndpoint = "wss://wss.woo.org/v2/ws/private/stream/"

	PublicTestnetBaseEndpoint  = "wss://wss.staging.woo.org/ws/stream/"
	PrivateTestnetBaseEndpoint = "wss://wss.staging.woo.org/v2/ws/private/stream/"
)

const (
	MaxTryTimes = 5
)

const (
	SUBSCRIBE   = "subscribe"
	UNSUBSCRIBE = "unsubscribe"
)
