package types

type NewOrderParam struct {
	Symbol   string    `url:"symbol" validate:"required"`
	Price    float64   `url:"price,omitempty" validate:"omitempty"`
	Vol      float64   `url:"vol,omitempty" validate:"required"`
	Leverage int       `url:"leverage,omitempty" validate:"omitempty"`
	Side     OrderSide `url:"side" validate:"required,oneof=1 2 3 4"`
	Type     OrderType `url:"type" validate:"required,oneof=1 2 3 4 5 6"`
	OpenType OpenType  `url:"openType" validate:"required,oneof=1 2"`

	PositionId      int64   `url:"positionId,omitempty" validate:"omitempty"`
	StopLossPrice   float64 `url:"stopLossPrice,omitempty" validate:"omitempty"`
	TakeProfitPrice float64 `url:"takeProfitPrice,omitempty" validate:"omitempty"`
	PositionMode    int     `url:"positionMode,omitempty" validate:"omitempty"`
}

type OrderSide = int

var (
	OpenLong   OrderSide = 1
	CloseShort OrderSide = 2
	OpenShort  OrderSide = 3
	CloseLong  OrderSide = 4
)

type OrderType = int

var (
	LimitOrder             OrderType = 1
	PostOnlyMaker          OrderType = 2
	TransactOrCancel       OrderType = 3
	TransactAllOrCancelAll OrderType = 4
	MarketOrder            OrderType = 5
	ConvertToCurrentPrice  OrderType = 6
)

type OpenType = int

var ()
