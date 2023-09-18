package okxutils

var (
	RestURL      = "https://www.okx.com"
	PublicWsURL  = "wss://ws.okx.com:8443/ws/v5/public"
	PrivateWsURL = "wss://ws.okx.com:8443/ws/v5/private"

	AWSRestURL      = "https://aws.okx.com"
	AWSPublicWsURL  = "wss://wsaws.okx.com:8443/ws/v5/public"
	AWSPrivateWsURL = "wss://wsaws.okx.com:8443/ws/v5/private"
)

type InstrumentType = string

const (
	Spot    = "SPOT"
	Margin  = "MARGIN"
	Swap    = "SWAP"
	Futures = "FUTURES"
	Option  = "OPTION"
)
