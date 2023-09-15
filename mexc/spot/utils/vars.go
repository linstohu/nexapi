package spotutils

var (
	BaseURL = "https://api.mexc.com"
)

type KlineInterval string

var (
	Minute1  KlineInterval = "1m"
	Minute5  KlineInterval = "5m"
	Minute15 KlineInterval = "15m"
	Minute30 KlineInterval = "30m"
	Minute60 KlineInterval = "60m"
	Hour4    KlineInterval = "4h"
	Day1     KlineInterval = "1d"
	Month1   KlineInterval = "1M"
)
