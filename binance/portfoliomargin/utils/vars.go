package cmutils

var (
	PortfolioMarginBaseURL = "https://papi.binance.com"
)

type SecurityType = string

var (
	NONE        SecurityType = "NONE"
	TRADE       SecurityType = "TRADE"
	USER_DATA   SecurityType = "USER_DATA"
	USER_STREAM SecurityType = "USER_STREAM"
)

var DefaultContentType = map[string]string{
	"Content-Type": "application/x-www-form-urlencoded",
	"Accept":       "application/json",
}
