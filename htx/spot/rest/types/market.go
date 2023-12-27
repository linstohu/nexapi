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
	htxutils "github.com/linstohu/nexapi/htx/utils"
)

type GetMergedMarketTickerParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type GetMergedMarketTickerResp struct {
	htxutils.V1Response
	Ts   int64 `json:"ts"`
	Tick Tick  `json:"tick,omitempty"`
}

type Tick struct {
	ID      int64     `json:"id,omitempty"`
	Version int64     `json:"version,omitempty"`
	Open    float64   `json:"open,omitempty"`
	Close   float64   `json:"close,omitempty"`
	Low     float64   `json:"low,omitempty"`
	High    float64   `json:"high,omitempty"`
	Amount  float64   `json:"amount,omitempty"`
	Vol     float64   `json:"vol,omitempty"`
	Count   int       `json:"count,omitempty"`
	Bid     []float64 `json:"bid,omitempty"`
	Ask     []float64 `json:"ask,omitempty"`
}

type GetMarketDepthParam struct {
	Symbol string `url:"symbol" validate:"required"`
	Depth  int    `url:"depth,omitempty" validate:"omitempty"`
	Type   string `url:"type" validate:"required"`
}

type GetMarketDepthResp struct {
	htxutils.V1Response
	Ts   int64       `json:"ts"`
	Tick MarketDepth `json:"tick,omitempty"`
}

type MarketDepth struct {
	Ts      int64       `json:"ts,omitempty"`
	Version int64       `json:"version,omitempty"`
	Bids    [][]float64 `json:"bids,omitempty"`
	Asks    [][]float64 `json:"asks,omitempty"`
}

type GetMarketTickerParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type GetMarketTickerResp struct {
	htxutils.V1Response
	Ts   int64              `json:"ts"`
	Tick MarketTickerMerged `json:"tick,omitempty"`
}

type MarketTickerMerged struct {
	ID      int64     `json:"id"`
	Version int64     `json:"version"`
	Open    float64   `json:"open"`
	Close   float64   `json:"close"`
	Low     float64   `json:"low"`
	High    float64   `json:"high"`
	Amount  float64   `json:"amount"`
	Vol     float64   `json:"vol"`
	Count   int       `json:"count"`
	Bid     []float64 `json:"bid"`
	Ask     []float64 `json:"ask"`
}

type GetMarketTickersResp struct {
	htxutils.V1Response
	Ts   int64          `json:"ts"`
	Data []MarketTicker `json:"data,omitempty"`
}

type MarketTicker struct {
	Symbol  string  `json:"symbol"`  // 交易对，例如 btcusdt, ethbtc
	Open    float64 `json:"open"`    // 开盘价（以新加坡时间自然日计）
	High    float64 `json:"high"`    // 最高价（以新加坡时间自然日计）
	Low     float64 `json:"low"`     // 最低价（以新加坡时间自然日计）
	Close   float64 `json:"close"`   // 最新价（以新加坡时间自然日计）
	Amount  float64 `json:"amount"`  // 以基础币种计量的交易量（以滚动24小时计）
	Vol     float64 `json:"vol"`     // 以报价币种计量的交易量（以滚动24小时计）
	Count   int     `json:"count"`   // 交易笔数（以滚动24小时计）
	Bid     float64 `json:"bid"`     // 买一价
	BidSize float64 `json:"bidSize"` // 买一量
	Ask     float64 `json:"ask"`     // 卖一价
	AskSize float64 `json:"askSize"` // 卖一量
}
