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

type GetExchangeInfoParam struct {
	Symbol string
}

type GetExchangeInfoParams struct {
	Symbol string `url:"symbols,omitempty" validate:"omitempty"`
}

type ExchangeInfo struct {
	Timezone   string `json:"timezone"`
	ServerTime int64  `json:"serverTime"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
	} `json:"rateLimits"`
	Symbols []struct {
		Symbol                   string   `json:"symbol"`
		Status                   string   `json:"status"`
		BaseAsset                string   `json:"baseAsset"`
		BaseAssetPrecision       int      `json:"baseAssetPrecision"`
		QuoteAsset               string   `json:"quoteAsset"`
		QuotePrecision           int      `json:"quotePrecision"`
		QuoteAssetPrecision      int      `json:"quoteAssetPrecision"`
		BaseCommissionPrecision  int      `json:"baseCommissionPrecision"`
		QuoteCommissionPrecision int      `json:"quoteCommissionPrecision"`
		OrderTypes               []string `json:"orderTypes"`
		IsSpotTradingAllowed     bool     `json:"isSpotTradingAllowed"`
		IsMarginTradingAllowed   bool     `json:"isMarginTradingAllowed"`
		QuoteAmountPrecision     string   `json:"quoteAmountPrecision"`
		BaseSizePrecision        string   `json:"baseSizePrecision"`
		Permissions              []string `json:"permissions"`
		Filters                  []struct {
			FilterType            string `json:"filterType"`
			MinPrice              string `json:"minPrice,omitempty"`
			MaxPrice              string `json:"maxPrice,omitempty"`
			TickSize              string `json:"tickSize,omitempty"`
			MinQty                string `json:"minQty,omitempty"`
			MaxQty                string `json:"maxQty,omitempty"`
			StepSize              string `json:"stepSize,omitempty"`
			Limit                 int    `json:"limit,omitempty"`
			MinTrailingAboveDelta int    `json:"minTrailingAboveDelta,omitempty"`
			MaxTrailingAboveDelta int    `json:"maxTrailingAboveDelta,omitempty"`
			MinTrailingBelowDelta int    `json:"minTrailingBelowDelta,omitempty"`
			MaxTrailingBelowDelta int    `json:"maxTrailingBelowDelta,omitempty"`
			BidMultiplierUp       string `json:"bidMultiplierUp,omitempty"`
			BidMultiplierDown     string `json:"bidMultiplierDown,omitempty"`
			AskMultiplierUp       string `json:"askMultiplierUp,omitempty"`
			AskMultiplierDown     string `json:"askMultiplierDown,omitempty"`
			AvgPriceMins          int    `json:"avgPriceMins,omitempty"`
			MinNotional           string `json:"minNotional,omitempty"`
			ApplyMinToMarket      bool   `json:"applyMinToMarket,omitempty"`
			MaxNotional           string `json:"maxNotional,omitempty"`
			ApplyMaxToMarket      bool   `json:"applyMaxToMarket,omitempty"`
			MaxNumOrders          int    `json:"maxNumOrders,omitempty"`
			MaxNumAlgoOrders      int    `json:"maxNumAlgoOrders,omitempty"`
		} `json:"filters"`
		MaxQuoteAmount             string `json:"maxQuoteAmount"`
		MakerCommission            string `json:"makerCommission"`
		TakerCommission            string `json:"takerCommission"`
		QuoteAmountPrecisionMarket string `json:"quoteAmountPrecisionMarket"`
		MaxQuoteAmountMarket       string `json:"maxQuoteAmountMarket"`
	} `json:"symbols"`
}
