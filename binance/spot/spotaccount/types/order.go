/*
 * Copyright (c) 2023, LinstoHu
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package types

import (
	"github.com/linstohu/nexapi/binance/utils"
)

type SideType string

var (
	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"
)

type OrderType string

var (
	Limit           OrderType = "LIMIT"
	Market          OrderType = "MARKET"
	StopLoss        OrderType = "STOP_LOSS"
	StopLossLimit   OrderType = "STOP_LOSS_LIMIT"
	TakeProfit      OrderType = "TAKE_PROFIT"
	TakeProfitLimit OrderType = "TAKE_PROFIT_LIMIT"
	LimitMaker      OrderType = "LIMIT_MAKER"
)

type TimeInForceType string

var (
	GTC TimeInForceType = "GTC"
	IOC TimeInForceType = "IOC"
	FOK TimeInForceType = "FOK"
)

type NewOrderRespType string

var (
	ACK    NewOrderRespType = "ACK"
	RESULT NewOrderRespType = "RESULT"
	FULL   NewOrderRespType = "FULL"
)

type OrderInfo struct {
	Symbol                  string `json:"symbol"`
	OrigClientOrderID       string `json:"origClientOrderId"`
	OrderID                 int64  `json:"orderId"`
	OrderListID             int    `json:"orderListId"`
	ClientOrderID           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
}

type Order struct {
	OrderInfo
	StopPrice         string `json:"stopPrice"`
	IcebergQty        string `json:"icebergQty"`
	Time              int64  `json:"time"`
	UpdateTime        int64  `json:"updateTime"`
	IsWorking         bool   `json:"isWorking"`
	WorkingTime       int64  `json:"workingTime"`
	OrigQuoteOrderQty string `json:"origQuoteOrderQty"`
}

type NewOrderParam struct {
	Symbol                  string           `url:"symbol" validate:"required"`
	Side                    SideType         `url:"side" validate:"required,oneof=BUY SELL"`
	Type                    OrderType        `url:"type" validate:"required"`
	TimeInForce             TimeInForceType  `url:"timeInForce,omitempty" validate:"omitempty"`
	Quantity                float64          `url:"quantity,omitempty" validate:"omitempty"`
	QuoteOrderQty           float64          `url:"quoteOrderQty,omitempty" validate:"omitempty"`
	Price                   float64          `url:"price,omitempty" validate:"omitempty"`
	NewClientOrderId        string           `url:"newClientOrderId,omitempty" validate:"omitempty"`
	StrategyID              int              `url:"strategyId,omitempty" validate:"omitempty"`
	StrategyType            int              `url:"strategyType,omitempty" validate:"omitempty"`
	StopPrice               float64          `url:"stopPrice,omitempty" validate:"omitempty"`
	TrailingDelta           int64            `url:"trailingDelta,omitempty" validate:"omitempty"`
	IcebergQty              float64          `url:"icebergQty,omitempty" validate:"omitempty"`
	NewOrderRespType        NewOrderRespType `url:"newOrderRespType,omitempty" validate:"omitempty"`
	SelfTradePreventionMode string           `url:"selfTradePreventionMode,omitempty" validate:"omitempty"`
}

type NewOrderParams struct {
	NewOrderParam
	utils.DefaultParam
}

type NewOrderResp struct {
	OrderInfo
	TransactTime int64 `json:"transactTime"`
	WorkingTime  int64 `json:"workingTime"`
}

type CancelOrderParam struct {
	Symbol             string `url:"symbol" validate:"required"`
	OrderID            int64  `url:"orderId,omitempty" validate:"omitempty"`
	OrigClientOrderId  string `url:"origClientOrderId,omitempty" validate:"omitempty"`
	NewClientOrderId   string `url:"newClientOrderId,omitempty" validate:"omitempty"`
	CancelRestrictions string `url:"cancelRestrictions,omitempty" validate:"omitempty"`
}

type CancelOrderParams struct {
	CancelOrderParam
	utils.DefaultParam
}

type CancelOrdersOnOneSymbolParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type CancelOrdersOnOneSymbolParams struct {
	CancelOrdersOnOneSymbolParam
	utils.DefaultParam
}

type QueryOrderParam struct {
	Symbol            string `url:"symbol" validate:"required"`
	OrderID           int64  `url:"orderId,omitempty" validate:"omitempty"`
	OrigClientOrderId string `url:"origClientOrderId,omitempty" validate:"omitempty"`
}

type QueryOrderParams struct {
	QueryOrderParam
	utils.DefaultParam
}

type GetOpenOrdersParam struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
}

type GetOpenOrdersParams struct {
	GetOpenOrdersParam
	utils.DefaultParam
}

type GetAllOrdersParam struct {
	Symbol    string `url:"symbol" validate:"required"`
	OrderID   int64  `url:"orderId,omitempty" validate:"omitempty"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type GetAllOrdersParams struct {
	GetAllOrdersParam
	utils.DefaultParam
}
