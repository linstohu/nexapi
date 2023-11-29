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

import okxutils "github.com/linstohu/nexapi/okx/utils"

type GetMarketTickersParam struct {
	InstType   InstrumentType `url:"instType,omitempty" validate:"omitempty,oneof=SPOT SWAP FUTURES OPTION"`
	Uly        string         `url:"uly,omitempty"`
	InstFamily string         `url:"instFamily,omitempty"`
}

type InstrumentType = string

const (
	Spot    = "SPOT"
	Swap    = "SWAP"
	Futures = "FUTURES"
	Option  = "OPTION"
)

type GetMarketTickersResp struct {
	okxutils.Response
	Data []*MarketTicker `json:"data"`
}

type MarketTicker struct {
	InstType  string `json:"instType,string"`
	InstID    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24h   string `json:"open24h"`
	High24h   string `json:"high24h"`
	Low24h    string `json:"low24h"`
	VolCcy24h string `json:"volCcy24h"`
	Vol24h    string `json:"vol24h"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
	TS        string `json:"ts"`
}
