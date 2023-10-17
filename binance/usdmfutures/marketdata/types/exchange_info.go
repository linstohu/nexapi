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

type ExchangeInfo struct {
	Timezone    string `json:"timezone"`
	ServerTime  int64  `json:"serverTime"`
	FuturesType string `json:"futuresType"`
	RateLimits  []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
	} `json:"rateLimits"`
	Assets []struct {
		Asset             string `json:"asset"`
		MarginAvailable   bool   `json:"marginAvailable"`
		AutoAssetExchange string `json:"autoAssetExchange"`
	} `json:"assets"`
	Symbols []struct {
		Symbol                string               `json:"symbol"`
		Pair                  string               `json:"pair"`
		ContractType          umutils.ContractType `json:"contractType"`
		DeliveryDate          int64                `json:"deliveryDate"`
		OnboardDate           int64                `json:"onboardDate"`
		Status                string               `json:"status"`
		MaintMarginPercent    string               `json:"maintMarginPercent"`
		RequiredMarginPercent string               `json:"requiredMarginPercent"`
		BaseAsset             string               `json:"baseAsset"`
		QuoteAsset            string               `json:"quoteAsset"`
		MarginAsset           string               `json:"marginAsset"`
		PricePrecision        int                  `json:"pricePrecision"`
		QuantityPrecision     int                  `json:"quantityPrecision"`
		BaseAssetPrecision    int                  `json:"baseAssetPrecision"`
		QuotePrecision        int                  `json:"quotePrecision"`
		UnderlyingType        string               `json:"underlyingType"`
		UnderlyingSubType     []string             `json:"underlyingSubType"`
		SettlePlan            int                  `json:"settlePlan"`
		TriggerProtect        string               `json:"triggerProtect"`
		LiquidationFee        string               `json:"liquidationFee"`
		MarketTakeBound       string               `json:"marketTakeBound"`
		MaxMoveOrderLimit     int                  `json:"maxMoveOrderLimit"`
		Filters               []struct {
			MinPrice          string `json:"minPrice,omitempty"`
			MaxPrice          string `json:"maxPrice,omitempty"`
			FilterType        string `json:"filterType"`
			TickSize          string `json:"tickSize,omitempty"`
			StepSize          string `json:"stepSize,omitempty"`
			MaxQty            string `json:"maxQty,omitempty"`
			MinQty            string `json:"minQty,omitempty"`
			Limit             int    `json:"limit,omitempty"`
			Notional          string `json:"notional,omitempty"`
			MultiplierDown    string `json:"multiplierDown,omitempty"`
			MultiplierUp      string `json:"multiplierUp,omitempty"`
			MultiplierDecimal string `json:"multiplierDecimal,omitempty"`
		} `json:"filters"`
		OrderTypes  []umutils.OrderType   `json:"orderTypes"`
		TimeInForce []umutils.TimeInForce `json:"timeInForce"`
	} `json:"symbols"`
}
