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
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type GetTradeListParam struct {
	Symbol    string `url:"symbol,omitempty" validate:"omitempty"`
	FromID    int64  `url:"fromId,omitempty" validate:"omitempty"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type GetTradeListParams struct {
	GetTradeListParam
	bnutils.DefaultParam
}

type UserTrade struct {
	ID             int64  `json:"id"`
	TradeID        int64  `json:"tradeId"`
	OrderID        int64  `json:"orderId"`
	Symbol         string `json:"symbol"`
	Price          string `json:"price"`
	Quantity       string `json:"quantity"`
	Fee            string `json:"fee"`
	RealizedProfit string `json:"realizedProfit"`
	Side           string `json:"side"`
	Type           string `json:"type"`
	Volatility     string `json:"volatility"`
	Liquidity      string `json:"liquidity"`
	QuoteAsset     string `json:"quoteAsset"`
	Time           int64  `json:"time"`
	PriceScale     int    `json:"priceScale"`
	QuantityScale  int    `json:"quantityScale"`
	OptionSide     string `json:"optionSide"`
}
