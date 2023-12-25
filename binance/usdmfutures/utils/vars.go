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

package utils

var (
	USDMarginedBaseURL = "https://fapi.binance.com"
)

type SecurityType = string

var (
	NONE        SecurityType = "NONE"
	TRADE       SecurityType = "TRADE"
	USER_DATA   SecurityType = "USER_DATA"
	USER_STREAM SecurityType = "USER_STREAM"
	MARKET_DATA SecurityType = "MARKET_DATA"
)

type ContractType = string

var (
	Perpetual           ContractType = "PERPETUAL"
	CurrentMonth        ContractType = "CURRENT_MONTH"
	NextMonth           ContractType = "NEXT_MONTH"
	CurrentQuarter      ContractType = "CURRENT_QUARTER"
	NextQuarter         ContractType = "NEXT_QUARTER"
	PerpetualDelivering ContractType = "PERPETUAL_DELIVERING"
)

type OrderStatus = string

var (
	New             OrderStatus = "NEW"
	PartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	Filled          OrderStatus = "FILLED"
	Canceled        OrderStatus = "CANCELED"
	Rejected        OrderStatus = "REJECTED"
	Expired         OrderStatus = "EXPIRED"
)

type OrderType = string

var (
	LimitOrder              OrderType = "LIMIT"
	MarketOrder             OrderType = "MARKET"
	StopOrder               OrderType = "STOP"
	StopMarketOrder         OrderType = "STOP_MARKET"
	TakeProfitOrder         OrderType = "TAKE_PROFIT"
	TakeProfitMarketOrder   OrderType = "TAKE_PROFIT_MARKET"
	TrailingStopMarketOrder OrderType = "TRAILING_STOP_MARKET"
)

type OrderSide = string

var (
	BuySide  OrderSide = "BUY"
	SellSide OrderSide = "SELL"
)

type PositionSide = string

var (
	Both  PositionSide = "BOTH"
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
	// GTX - Good Till Crossing (Post Only)
	GTX TimeInForce = "GTX"
)

type NewOrderRespType = string

var (
	ACK    NewOrderRespType = "ACK"
	RESULT NewOrderRespType = "RESULT"
)

type KlineInterval = string

var (
	Minute1  KlineInterval = "1m"
	Minute3  KlineInterval = "3m"
	Minute5  KlineInterval = "5m"
	Minute15 KlineInterval = "15m"
	Minute30 KlineInterval = "30m"
	Hour1    KlineInterval = "1h"
	Hour2    KlineInterval = "2h"
	Hour4    KlineInterval = "4h"
	Hour6    KlineInterval = "6h"
	Hour8    KlineInterval = "8h"
	Hour12   KlineInterval = "12h"
	Day1     KlineInterval = "1d"
	Day3     KlineInterval = "3d"
	Week1    KlineInterval = "1w"
	Month1   KlineInterval = "1M"
)

type MarginType = string

var (
	ISOLATED MarginType = "ISOLATED"
	CROSSED  MarginType = "CROSSED"
)
