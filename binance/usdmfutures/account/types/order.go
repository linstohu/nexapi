package types

import (
	umutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type NewOrderParam struct {
	Symbol           string                   `url:"symbol" validate:"required"`
	Side             umutils.OrderSide        `url:"side" validate:"required,oneof=BUY SELL"`
	PositionSide     umutils.PositionSide     `url:"positionSide,omitempty" validate:"omitempty"`
	Type             umutils.OrderType        `url:"type" validate:"required"`
	TimeInForce      umutils.TimeInForce      `url:"timeInForce,omitempty" validate:"omitempty"`
	Quantity         float64                  `url:"quantity,omitempty" validate:"omitempty"`
	ReduceOnly       string                   `url:"reduceOnly,omitempty" validate:"omitempty,oneof=true false"`
	Price            float64                  `url:"price,omitempty" validate:"omitempty"`
	NewClientOrderId string                   `url:"newClientOrderId,omitempty" validate:"omitempty"`
	StopPrice        float64                  `url:"stopPrice,omitempty" validate:"omitempty"`
	ClosePosition    string                   `url:"closePosition,omitempty" validate:"omitempty,oneof=true false"`
	PriceProtect     string                   `url:"priceProtect,omitempty" validate:"omitempty,oneof=true false"`
	NewOrderRespType umutils.NewOrderRespType `url:"newOrderRespType,omitempty" validate:"omitempty"`
}

type NewOrderParams struct {
	NewOrderParam
	bnutils.DefaultParam
}
