package utils

var (
	OptionsBaseURL = "https://eapi.binance.com"
)

type ContractType = string

var (
	CALL ContractType = "CALL"
	PUT  ContractType = "PUT"
)

type OrderSide = string

var (
	BuySide  OrderSide = "BUY"
	SellSide OrderSide = "SELL"
)

type PositionSide = string

var (
	Long  PositionSide = "LONG"
	Short PositionSide = "SHORT"
)

type TimeInForce = string

var (
	// GTC - Good Till Cancel
	GTC TimeInForce = "GTC"
	// IOC - Immediate or Cancel
	IOC TimeInForce = "IOC"
	// FOK - Fill or Kill
	FOK TimeInForce = "FOK"
)

type NewOrderRespType = string

var (
	ACK    NewOrderRespType = "ACK"
	RESULT NewOrderRespType = "RESULT"
)

type OrderStatus = string

var (
	Accepted        OrderStatus = "ACCEPTED"
	Rejected        OrderStatus = "REJECTED"
	PartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	Filled          OrderStatus = "FILLED"
	Canceled        OrderStatus = "CANCELED"
)
