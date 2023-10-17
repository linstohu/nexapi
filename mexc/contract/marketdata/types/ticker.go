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

type GetTickerForSymbolParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type TickerParams struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
}

type GetTickerResp struct {
	Response
	Data *Ticker `json:"data"`
}

type GetAllTickersResp struct {
	Response
	Data []*Ticker `json:"data"`
}

type Ticker struct {
	ContractID    int     `json:"contractId"`
	Symbol        string  `json:"symbol"`
	LastPrice     float64 `json:"lastPrice"`
	Bid1          float64 `json:"bid1"`
	Ask1          float64 `json:"ask1"`
	Volume24      float64 `json:"volume24"`
	Amount24      float64 `json:"amount24"`
	HoldVol       float64 `json:"holdVol"`
	Lower24Price  float64 `json:"lower24Price"`
	High24Price   float64 `json:"high24Price"`
	RiseFallRate  float64 `json:"riseFallRate"`
	RiseFallValue float64 `json:"riseFallValue"`
	IndexPrice    float64 `json:"indexPrice"`
	FairPrice     float64 `json:"fairPrice"`
	FundingRate   float64 `json:"fundingRate"`
	MaxBidPrice   float64 `json:"maxBidPrice"`
	MinAskPrice   float64 `json:"minAskPrice"`
	Timestamp     int64   `json:"timestamp"`
	RiseFallRates struct {
		Zone string  `json:"zone"`
		R    float64 `json:"r"`
		V    float64 `json:"v"`
		R7   float64 `json:"r7"`
		R30  float64 `json:"r30"`
		R90  float64 `json:"r90"`
		R180 float64 `json:"r180"`
		R365 float64 `json:"r365"`
	} `json:"riseFallRates"`
	RiseFallRatesOfTimezone []float64 `json:"riseFallRatesOfTimezone"`
}
