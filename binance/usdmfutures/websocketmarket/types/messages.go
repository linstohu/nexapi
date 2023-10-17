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
)

type AggregateTrade struct {
	EventType        string `json:"e"`
	EventTime        int64  `json:"E"`
	Symbol           string `json:"s"`
	AggregateTradeID int64  `json:"a"`
	Price            string `json:"p"`
	Quantity         string `json:"q"`
	FirstTradeID     int64  `json:"f"`
	LastTradeID      int64  `json:"l"`
	TradeTime        int64  `json:"T"`
	IsBuyerMaker     bool   `json:"m"`
}

type MarkPrice struct {
	EventType            string `json:"e"`
	EventTime            int64  `json:"E"`
	Symbol               string `json:"s"`
	MarkPrice            string `json:"p"`
	IndexPrice           string `json:"i"`
	EstimatedSettlePrice string `json:"P"`
	FundingRate          string `json:"r"`
	NextFundingTime      int64  `json:"T"`
}

type Ticker struct {
	EventType                   string `json:"e"`
	EventTime                   int64  `json:"E"`
	Symbol                      string `json:"s"`
	PriceChange                 string `json:"p"`
	PriceChangePercent          string `json:"P"`
	WeightedAveragePrice        string `json:"w"`
	LastPrice                   string `json:"c"`
	LastQuantity                string `json:"Q"`
	OpenPrice                   string `json:"o"`
	HighPrice                   string `json:"h"`
	LowPrice                    string `json:"l"`
	TotalTradedBaseAssetVolume  string `json:"v"`
	TotalTradedQuoteAssetVolume string `json:"q"`
	StatisticsOpenTime          int64  `json:"O"`
	StatisticsCloseTime         int64  `json:"C"`
	FirstTradeID                int64  `json:"F"`
	LastTradeID                 int64  `json:"L"`
	TradesNum                   int64  `json:"n"`
}

type BookTicker struct {
	EventType       string `json:"e"`
	UpdateID        int64  `json:"u"`
	EventTime       int64  `json:"E"`
	TransactionTime int64  `json:"T"`
	Symbol          string `json:"s"`
	BestBidPrice    string `json:"b"`
	BestBidQty      string `json:"B"`
	BestAskPrice    string `json:"a"`
	BestAskQty      string `json:"A"`
}

type LiquidationOrder struct {
	EventType string `json:"e"`
	EventTime int64  `json:"E"`
	Order     struct {
		Symbol                    string              `json:"s"`
		Side                      string              `json:"S"`
		OrderType                 umutils.OrderType   `json:"o"`
		TimeInForce               umutils.TimeInForce `json:"f"`
		OriginalQuantity          string              `json:"q"`
		Price                     string              `json:"p"`
		AveragePrice              string              `json:"ap"`
		OrderStatus               umutils.OrderStatus `json:"X"`
		LastFilledQuantity        string              `json:"l"`
		FilledAccumulatedQuantity string              `json:"z"`
		TradeTime                 int64               `json:"T"`
	} `json:"o"`
}

type OrderbookDepth struct {
	EventType               string     `json:"e"`
	EventTime               int64      `json:"E"`
	TransactionTime         int64      `json:"T"`
	Symbol                  string     `json:"s"`
	FirstUpdateID           int64      `json:"U"`
	FinalUpdateID           int64      `json:"u"`
	FinalUpdateIDLastStream int64      `json:"pu"`
	Bids                    [][]string `json:"b"`
	Asks                    [][]string `json:"a"`
}
