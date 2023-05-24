package types

import (
	"github.com/linstohu/nexapi/binance/europeanoptions/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type Order struct {
	OrderID       int64  `json:"orderId"`
	Symbol        string `json:"symbol"`
	Price         string `json:"price"`
	Quantity      string `json:"quantity"`
	ExecutedQty   string `json:"executedQty"`
	Fee           string `json:"fee"`
	Side          string `json:"side"`
	Type          string `json:"type"`
	TimeInForce   string `json:"timeInForce"`
	ReduceOnly    bool   `json:"reduceOnly"`
	PostOnly      bool   `json:"postOnly"`
	CreateTime    int64  `json:"createTime"`
	UpdateTime    int64  `json:"updateTime"`
	Status        string `json:"status"`
	AvgPrice      string `json:"avgPrice"`
	Source        string `json:"source,omitempty"`
	ClientOrderID string `json:"clientOrderId"`
	PriceScale    int    `json:"priceScale"`
	QuantityScale int    `json:"quantityScale"`
	OptionSide    string `json:"optionSide"`
	QuoteAsset    string `json:"quoteAsset"`
	Mmp           bool   `json:"mmp"`
}

type NewOrderParam struct {
	Symbol           string                 `url:"symbol" validate:"required"`
	Side             utils.OrderSide        `url:"side" validate:"required,oneof=BUY SELL"`
	Type             utils.OrderType        `url:"type" validate:"required"`
	Quantity         float64                `url:"quantity" validate:"required"`
	Price            float64                `url:"price,omitempty" validate:"omitempty"`
	TimeInForce      utils.TimeInForce      `url:"timeInForce,omitempty" validate:"omitempty"`
	ReduceOnly       bool                   `url:"reduceOnly,omitempty" validate:"omitempty"`
	PostOnly         bool                   `url:"postOnly,omitempty" validate:"omitempty"`
	NewOrderRespType utils.NewOrderRespType `url:"newOrderRespType,omitempty" validate:"omitempty"`
	ClientOrderID    string                 `url:"clientOrderId,omitempty" validate:"omitempty"`
	IsMmp            bool                   `url:"isMmp,omitempty" validate:"omitempty"`
}

type NewOrderParams struct {
	NewOrderParam
	bnutils.DefaultParam
}

type GetSingleOrderParam struct {
	Symbol        string `url:"symbol" validate:"required"`
	OrderID       int64  `url:"orderId,omitempty" validate:"omitempty"`
	ClientOrderID string `url:"clientOrderId,omitempty" validate:"omitempty"`
}

type GetSingleOrderParams struct {
	GetSingleOrderParam
	bnutils.DefaultParam
}

type CancelOrderParam struct {
	Symbol        string `url:"symbol" validate:"required"`
	OrderID       int64  `url:"orderId,omitempty" validate:"omitempty"`
	ClientOrderID string `url:"clientOrderId,omitempty" validate:"omitempty"`
}

type CancelOrderParams struct {
	CancelOrderParam
	bnutils.DefaultParam
}

type CancelAllOrdersParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type CancelAllOrdersParams struct {
	CancelAllOrdersParam
	bnutils.DefaultParam
}

type CancelAllOrdersByUnderlyingParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type CancelAllOrdersByUnderlyingParams struct {
	CancelAllOrdersByUnderlyingParam
	bnutils.DefaultParam
}

type GetCurrentOpenOrdersParam struct {
	Symbol    string `url:"symbol,omitempty" validate:"omitempty"`
	OrderID   int64  `url:"orderId,omitempty" validate:"omitempty"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type GetCurrentOpenOrdersParams struct {
	GetCurrentOpenOrdersParam
	bnutils.DefaultParam
}

type GetOrderHistoryParam struct {
	Symbol    string `url:"symbol" validate:"required"`
	OrderID   int64  `url:"orderId,omitempty" validate:"omitempty"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type GetOrderHistoryParams struct {
	GetOrderHistoryParam
	bnutils.DefaultParam
}
