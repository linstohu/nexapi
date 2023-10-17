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
	umutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type Order struct {
	ClientOrderID string `json:"clientOrderId"`
	CumQty        string `json:"cumQty"`
	CumBase       string `json:"cumBase"`
	ExecutedQty   string `json:"executedQty"`
	OrderID       int64  `json:"orderId"`
	AvgPrice      string `json:"avgPrice"`
	OrigQty       string `json:"origQty"`
	Price         string `json:"price"`
	ReduceOnly    bool   `json:"reduceOnly"`
	Side          string `json:"side"`
	PositionSide  string `json:"positionSide"`
	Status        string `json:"status"`
	StopPrice     string `json:"stopPrice"`
	ClosePosition bool   `json:"closePosition"`
	Symbol        string `json:"symbol"`
	Pair          string `json:"pair"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	OrigType      string `json:"origType"`
	ActivatePrice string `json:"activatePrice"`
	PriceRate     string `json:"priceRate"`
	UpdateTime    int64  `json:"updateTime"`
	WorkingType   string `json:"workingType"`
	PriceProtect  bool   `json:"priceProtect"`
}

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

type GetOrderParam struct {
	Symbol            string `url:"symbol" validate:"required"`
	OrderID           int64  `url:"orderId,omitempty" validate:"omitempty"`
	OrigClientOrderId string `url:"origClientOrderId,omitempty" validate:"omitempty"`
}

type GetOrderParams struct {
	GetOrderParam
	bnutils.DefaultParam
}

type CancelAllOpenOrdersParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type CancelAllOpenOrdersParams struct {
	CancelAllOpenOrdersParam
	bnutils.DefaultParam
}

type GetAllOpenOrdersParam struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
	Pair   string `url:"pair,omitempty" validate:"omitempty"`
}

type GetAllOpenOrdersParams struct {
	GetAllOpenOrdersParam
	bnutils.DefaultParam
}

type GetAllOrdersParam struct {
	Symbol    string `url:"symbol" validate:"required"`
	Pair      string `url:"pair,omitempty" validate:"omitempty"`
	OrderID   int64  `url:"orderId,omitempty" validate:"omitempty"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type GetAllOrdersParams struct {
	GetAllOrdersParam
	bnutils.DefaultParam
}
