package utils

var (
	CoinMarginedBaseURL = "https://dapi.binance.com"
)

type ContractType = string

var (
	All            ContractType = "ALL"
	Perpetual      ContractType = "PERPETUAL"
	CurrentQuarter ContractType = "CURRENT_QUARTER"
	NextQuarter    ContractType = "NEXT_QUARTER"
)

type OrderStatus = string

var (
	New             OrderStatus = "NEW"
	PartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	Filled          OrderStatus = "FILLED"
	Canceled        OrderStatus = "CANCELED"
	Expired         OrderStatus = "EXPIRED"
)
