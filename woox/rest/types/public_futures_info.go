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

type FuturesInfo struct {
	Symbol          string  `json:"symbol"`
	IndexPrice      float64 `json:"index_price"`
	MarkPrice       float64 `json:"mark_price"`
	EstFundingRate  float64 `json:"est_funding_rate"`
	LastFundingRate float64 `json:"last_funding_rate"`
	NextFundingTime int64   `json:"next_funding_time"`
	OpenInterest    float64 `json:"open_interest"`
	Two4HOpen       float64 `json:"24h_open"`
	Two4HClose      float64 `json:"24h_close"`
	Two4HHigh       float64 `json:"24h_high"`
	Two4HLow        float64 `json:"24h_low"`
	Two4HVolume     float64 `json:"24h_volume"`
	Two4HAmount     float64 `json:"24h_amount"`
}

type AllFuturesInfo struct {
	Success   bool          `json:"success"`
	Rows      []FuturesInfo `json:"rows"`
	Timestamp int64         `json:"timestamp"`
}

type OneFuturesInfo struct {
	Response
	Info      FuturesInfo `json:"info"`
	Timestamp int64       `json:"timestamp"`
}
