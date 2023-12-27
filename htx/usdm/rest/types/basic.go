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

type GetContractInfoParam struct {
	ContractCode      string `url:"contract_code,omitempty" validate:"omitempty"`
	SupportMarginMode string `url:"support_margin_mode,omitempty" validate:"omitempty"`
	Pair              string `url:"pair,omitempty" validate:"omitempty"`
	ContractType      string `url:"contract_type,omitempty" validate:"omitempty"`
	BusinessType      string `url:"business_type,omitempty" validate:"omitempty"`
}

type GetContractInfoResp struct {
	DefaultResponse
	Data []ContractInfo `json:"data"`
}

type ContractInfo struct {
	Symbol            string  `json:"symbol,omitempty"`
	ContractCode      string  `json:"contract_code,omitempty"`
	ContractSize      float64 `json:"contract_size,omitempty"`
	PriceTick         float64 `json:"price_tick,omitempty"`
	DeliveryDate      string  `json:"delivery_date,omitempty"`
	DeliveryTime      string  `json:"delivery_time,omitempty"`
	CreateDate        string  `json:"create_date,omitempty"`
	ContractStatus    int     `json:"contract_status,omitempty"`
	SettlementDate    string  `json:"settlement_date,omitempty"`
	SupportMarginMode string  `json:"support_margin_mode,omitempty"`
	BusinessType      string  `json:"business_type,omitempty"`
	Pair              string  `json:"pair,omitempty"`
	ContractType      string  `json:"contract_type,omitempty"`
}

type GetFundingRateResp struct {
	DefaultResponse
	Data []FundingRate `json:"data,omitempty"`
}

type FundingRate struct {
	EstimatedRate   string `json:"estimated_rate"`
	FundingRate     string `json:"funding_rate"`
	ContractCode    string `json:"contract_code"`
	Symbol          string `json:"symbol"`
	FeeAsset        string `json:"fee_asset"`
	FundingTime     string `json:"funding_time"`
	NextFundingTime string `json:"next_funding_time"`
}

type GetMarketDepthParam struct {
	ContractCode string `url:"contract_code,omitempty" validate:"omitempty"`
	Type         string `url:"type" validate:"required"`
}

type GetMarketDepthResp struct {
	htxutils.V1Response
	Ts   int64       `json:"ts"`
	Tick MarketDepth `json:"tick,omitempty"`
}

type MarketDepth struct {
	Ch      string      `json:"ch,omitempty"`
	Ts      int64       `json:"ts,omitempty"`
	Version int64       `json:"version,omitempty"`
	Bids    [][]float64 `json:"bids,omitempty"`
	Asks    [][]float64 `json:"asks,omitempty"`
}

type GetMarketTickerParam struct {
	ContractCode string `url:"contract_code" validate:"required"`
}

type GetMarketTickerResp struct {
	htxutils.V1Response
	Ts   int64        `json:"ts"`
	Tick MarketTicker `json:"tick,omitempty"`
}

type GetMarketTickersParam struct {
	ContractCode string `url:"contract_code,omitempty" validate:"omitempty"`
	BusinessType string `url:"business_type,omitempty" validate:"omitempty"`
}

type GetMarketTickersResp struct {
	htxutils.V1Response
	Ts    int64          `json:"ts"`
	Ticks []MarketTicker `json:"ticks,omitempty"`
}

type MarketTicker struct {
	ID           int       `json:"id"`
	Ts           int64     `json:"ts"`
	Ask          []float64 `json:"ask"`
	Bid          []float64 `json:"bid"`
	BusinessType string    `json:"business_type"`
	ContractCode string    `json:"contract_code"`
	Open         string    `json:"open"`
	Close        string    `json:"close"`
	Low          string    `json:"low"`
	High         string    `json:"high"`
	Amount       string    `json:"amount"`
	Count        int       `json:"count"`
	Vol          string    `json:"vol"`
	NumberOf     string    `json:"number_of"`
}
