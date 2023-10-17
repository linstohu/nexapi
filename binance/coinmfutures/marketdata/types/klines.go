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

type Kline struct {
	OpenTime                int64  `json:"openTime"`
	OpenPrice               string `json:"openPrice"`
	HighPrice               string `json:"highPrice"`
	LowPrice                string `json:"lowPrice"`
	ClosePrice              string `json:"closePrice"`
	Volume                  string `json:"volume"`
	CloseTime               int64  `json:"closeTime"`
	BaseAssetVolume         string `json:"baseAssetVolume"`
	NumberOfTrades          int64  `json:"numberOfTrades"`
	TakerBuyVolume          string `json:"takerBuyVolume"`
	TakerBuyBaseAssetVolume string `json:"takerBuyBaseAssetVolume"`
}
