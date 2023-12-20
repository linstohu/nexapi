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

type Kline struct {
	ID            int64   `json:"id,omitempty"`
	MrID          int64   `json:"mrid,omitempty"`
	Open          float64 `json:"open,omitempty"`
	Close         float64 `json:"close,omitempty"`
	Low           float64 `json:"low,omitempty"`
	High          float64 `json:"high,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
	Vol           float64 `json:"vol,omitempty"`
	Count         int     `json:"count,omitempty"`
	TradeTurnover float64 `json:"trade_turnover,omitempty"`
}

type Depth struct {
	MrID    int64       `json:"mrid,omitempty"`
	ID      int64       `json:"id,omitempty"`
	Bids    [][]float64 `json:"bids,omitempty"`
	Asks    [][]float64 `json:"asks,omitempty"`
	Ts      int64       `json:"ts,omitempty"`
	Version int64       `json:"version,omitempty"`
	Ch      string      `json:"ch,omitempty"`
}

type BBO struct {
	MrID    int64     `json:"mrid,omitempty"`
	ID      int64     `json:"id,omitempty"`
	Bid     []float64 `json:"bid,omitempty"`
	Ask     []float64 `json:"ask,omitempty"`
	Ts      int64     `json:"ts,omitempty"`
	Version int64     `json:"version,omitempty"`
	Ch      string    `json:"ch,omitempty"`
}

type MarketTradeMsg struct {
	ID   int64         `json:"id,omitempty"`
	Ts   int64         `json:"ts,omitempty"`
	Data []MarketTrade `json:"data,omitempty"`
}

type MarketTrade struct {
	ID            int64   `json:"id,omitempty"`
	Ts            int64   `json:"ts,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
	Price         float64 `json:"price,omitempty"`
	Direction     string  `json:"direction,omitempty"`
	Quantity      float64 `json:"quantity,omitempty"`
	TradeTurnover float64 `json:"trade_turnover,omitempty"`
}
