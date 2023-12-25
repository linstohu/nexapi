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

import "github.com/linstohu/nexapi/utils"

type GetExchangeInfoResp struct {
	Http *utils.ApiResponse
	Body *ExchangeInfo
}

type ExchangeInfo struct {
	Timezone        string `json:"timezone"`
	ServerTime      int64  `json:"serverTime"`
	OptionContracts []struct {
		ID          int    `json:"id"`
		BaseAsset   string `json:"baseAsset"`
		QuoteAsset  string `json:"quoteAsset"`
		Underlying  string `json:"underlying"`
		SettleAsset string `json:"settleAsset"`
	} `json:"optionContracts"`
	OptionAssets []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"optionAssets"`
	OptionSymbols []struct {
		ContractID int64 `json:"contractId"`
		ExpiryDate int64 `json:"expiryDate"`
		Filters    []struct {
			FilterType string `json:"filterType"`
			MinPrice   string `json:"minPrice,omitempty"`
			MaxPrice   string `json:"maxPrice,omitempty"`
			TickSize   string `json:"tickSize,omitempty"`
			MinQty     string `json:"minQty,omitempty"`
			MaxQty     string `json:"maxQty,omitempty"`
			StepSize   string `json:"stepSize,omitempty"`
		} `json:"filters"`
		ID                   int64  `json:"id"`
		Symbol               string `json:"symbol"`
		Side                 string `json:"side"`
		StrikePrice          string `json:"strikePrice"`
		Underlying           string `json:"underlying"`
		Unit                 int64  `json:"unit"`
		MakerFeeRate         string `json:"makerFeeRate"`
		TakerFeeRate         string `json:"takerFeeRate"`
		MinQty               string `json:"minQty"`
		MaxQty               string `json:"maxQty"`
		InitialMargin        string `json:"initialMargin"`
		MaintenanceMargin    string `json:"maintenanceMargin"`
		MinInitialMargin     string `json:"minInitialMargin"`
		MinMaintenanceMargin string `json:"minMaintenanceMargin"`
		PriceScale           int    `json:"priceScale"`
		QuantityScale        int    `json:"quantityScale"`
		QuoteAsset           string `json:"quoteAsset"`
	} `json:"optionSymbols"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
	} `json:"rateLimits"`
}
