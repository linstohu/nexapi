package spotutils

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

type KlineInterval string

var (
	Second1  KlineInterval = "1s"
	Minute1  KlineInterval = "1m"
	Minute3  KlineInterval = "3m"
	Minute5  KlineInterval = "5m"
	Minute15 KlineInterval = "15m"
	Minute30 KlineInterval = "30m"
	Hour1    KlineInterval = "1h"
	Hour2    KlineInterval = "2h"
	Hour4    KlineInterval = "4h"
	Hour6    KlineInterval = "6h"
	Hour8    KlineInterval = "8h"
	Hour12   KlineInterval = "12h"
	Day1     KlineInterval = "1d"
	Day3     KlineInterval = "3d"
	Week1    KlineInterval = "1w"
	Month1   KlineInterval = "1M"
)
