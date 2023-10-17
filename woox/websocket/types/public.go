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

type Orderbook struct {
	Topic string `json:"topic"`
	Ts    int64  `json:"ts"`
	Data  struct {
		Symbol string      `json:"symbol"`
		Asks   [][]float64 `json:"asks"`
		Bids   [][]float64 `json:"bids"`
	} `json:"data"`
}

type Trade struct {
	Topic string `json:"topic"`
	Ts    int64  `json:"ts"`
	Data  struct {
		Symbol string  `json:"symbol"`
		Price  float64 `json:"price"`
		Size   float64 `json:"size"`
		Side   string  `json:"side"`
		Source int     `json:"source"`
	} `json:"data"`
}

type Ticker struct {
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	Close  float64 `json:"close"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Volume float64 `json:"volume"`
	Amount float64 `json:"amount"`
	Count  int     `json:"count"`
}

type Ticker24H struct {
	Topic string `json:"topic"`
	Ts    int64  `json:"ts"`
	Data  Ticker `json:"data"`
}

type Tickers struct {
	Topic string   `json:"topic"`
	Ts    int64    `json:"ts"`
	Data  []Ticker `json:"data"`
}

type BestBidOffer struct {
	Symbol  string  `json:"symbol"`
	Ask     float64 `json:"ask"`
	AskSize float64 `json:"askSize"`
	Bid     float64 `json:"bid"`
	BidSize float64 `json:"bidSize"`
}

type BBO struct {
	Topic string       `json:"topic"`
	Ts    int64        `json:"ts"`
	Data  BestBidOffer `json:"data"`
}

type AllBBO struct {
	Topic string         `json:"topic"`
	Ts    int64          `json:"ts"`
	Data  []BestBidOffer `json:"data"`
}

type Kline struct {
	Topic string `json:"topic"`
	Ts    int64  `json:"ts"`
	Data  struct {
		Symbol    string  `json:"symbol"`
		Type      string  `json:"type"`
		Open      float64 `json:"open"`
		Close     float64 `json:"close"`
		High      float64 `json:"high"`
		Low       float64 `json:"low"`
		Volume    float64 `json:"volume"`
		Amount    float64 `json:"amount"`
		StartTime int64   `json:"startTime"`
		EndTime   int64   `json:"endTime"`
	} `json:"data"`
}

type SymbolPrice struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

type IndexPrice struct {
	Topic string      `json:"topic"`
	Ts    int64       `json:"ts"`
	Data  SymbolPrice `json:"data"`
}

type MarkPrice struct {
	Topic string      `json:"topic"`
	Ts    int64       `json:"ts"`
	Data  SymbolPrice `json:"data"`
}

type MarkPrices struct {
	Topic string        `json:"topic"`
	Ts    int64         `json:"ts"`
	Data  []SymbolPrice `json:"data"`
}

type OpenInterest struct {
	Topic string `json:"topic"`
	Ts    int64  `json:"ts"`
	Data  struct {
		Symbol       string  `json:"symbol"`
		OpenInterest float64 `json:"openInterest"`
	} `json:"data"`
}

type EstFundingRate struct {
	Topic string `json:"topic"`
	Ts    int64  `json:"ts"`
	Data  struct {
		Symbol      string  `json:"symbol"`
		FundingRate float64 `json:"fundingRate"`
		FundingTs   int64   `json:"fundingTs"`
	} `json:"data"`
}
