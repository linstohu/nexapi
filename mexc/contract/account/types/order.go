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
