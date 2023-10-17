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

type SymbolInfo struct {
	Response
	Info struct {
		Symbol      string  `json:"symbol"`
		QuoteMin    float64 `json:"quote_min"`
		QuoteMax    float64 `json:"quote_max"`
		QuoteTick   float64 `json:"quote_tick"`
		BaseMin     float64 `json:"base_min"`
		BaseMax     float64 `json:"base_max"`
		BaseTick    float64 `json:"base_tick"`
		MinNotional float64 `json:"min_notional"`
		PriceRange  float64 `json:"price_range"`
		PriceScope  float64 `json:"price_scope"`
		CreatedTime string  `json:"created_time"`
		UpdatedTime string  `json:"updated_time"`
		IsStable    int     `json:"is_stable"`
	} `json:"info"`
}

type AvailableSymbols struct {
	Response
	Rows []struct {
		Symbol      string  `json:"symbol"`
		QuoteMin    float64 `json:"quote_min"`
		QuoteMax    float64 `json:"quote_max"`
		QuoteTick   float64 `json:"quote_tick"`
		BaseMin     float64 `json:"base_min"`
		BaseMax     float64 `json:"base_max"`
		BaseTick    float64 `json:"base_tick"`
		MinNotional float64 `json:"min_notional"`
		PriceRange  float64 `json:"price_range"`
		PriceScope  float64 `json:"price_scope"`
		CreatedTime string  `json:"created_time"`
		UpdatedTime string  `json:"updated_time"`
		IsStable    int     `json:"is_stable"`
	} `json:"rows"`
}
