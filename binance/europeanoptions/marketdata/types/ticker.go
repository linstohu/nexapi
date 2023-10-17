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

type GetTickerPriceParam struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
}

type TickerPrice struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	Open               string `json:"open"`
	High               string `json:"high"`
	Low                string `json:"low"`
	Volume             string `json:"volume"`
	Amount             string `json:"amount"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstTradeID       int64  `json:"firstTradeId"`
	TradeCount         int64  `json:"tradeCount"`
	StrikePrice        string `json:"strikePrice"`
	ExercisePrice      string `json:"exercisePrice"`
}
