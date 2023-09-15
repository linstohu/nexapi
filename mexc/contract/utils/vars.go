package ctutils

var (
	BaseURL = "https://contract.mexc.com"
)

type KlineInterval string

var (
	Minute1  KlineInterval = "Min1"
	Minute5  KlineInterval = "Min5"
	Minute15 KlineInterval = "Min15"
	Minute30 KlineInterval = "Min30"
	Minute60 KlineInterval = "Min60"
	Hour4    KlineInterval = "Hour4"
	Hour8    KlineInterval = "Hour8"
	Day1     KlineInterval = "Day1"
	Week1    KlineInterval = "Week1"
	Month1   KlineInterval = "Month1"
)
