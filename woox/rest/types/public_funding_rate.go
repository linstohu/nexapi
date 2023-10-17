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

type FundingRates struct {
	Success   bool                `json:"success"`
	Rows      []SymbolFundingRate `json:"rows"`
	Timestamp int64               `json:"timestamp"`
}

type FundingRate struct {
	Response
	SymbolFundingRate
	Timestamp int64 `json:"timestamp"`
}

type SymbolFundingRate struct {
	Symbol                   string  `json:"symbol"`
	EstFundingRate           float64 `json:"est_funding_rate"`
	EstFundingRateTimestamp  int64   `json:"est_funding_rate_timestamp"`
	LastFundingRate          float64 `json:"last_funding_rate"`
	LastFundingRateTimestamp int64   `json:"last_funding_rate_timestamp"`
	NextFundingTime          int64   `json:"next_funding_time"`
	LastFundingRateInterval  int     `json:"last_funding_rate_interval"`
	EstFundingRateInterval   int     `json:"est_funding_rate_interval"`
}
