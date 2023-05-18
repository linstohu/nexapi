package utils

var (
	BaseURL = "https://api.binance.com"
)

type SecurityType string

var (
	NONE        SecurityType = "NONE"
	TRADE       SecurityType = "TRADE"
	MARGIN      SecurityType = "MARGIN"
	USER_DATA   SecurityType = "USER_DATA"
	USER_STREAM SecurityType = "USER_STREAM"
	MARKET_DATA SecurityType = "MARKET_DATA"
)

var DefaultContentType = map[string]string{
	"Content-Type": "application/x-www-form-urlencoded",
	"Accept":       "application/json",
}
