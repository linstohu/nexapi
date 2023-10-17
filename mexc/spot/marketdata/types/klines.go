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
	spotutils "github.com/linstohu/nexapi/mexc/spot/utils"
)

type GetKlineParam struct {
	Symbol    string                  `url:"symbol" validate:"required"`
	Interval  spotutils.KlineInterval `url:"interval" validate:"required,oneof=1m 5m 15m 30m 60m 4h 1d 1M"`
	StartTime int64                   `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64                   `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int                     `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type Kline struct {
	OpenTime         int64  `json:"openTime"`
	OpenPrice        string `json:"openPrice"`
	HighPrice        string `json:"highPrice"`
	LowPrice         string `json:"lowPrice"`
	ClosePrice       string `json:"closePrice"`
	Volume           string `json:"volume"`
	CloseTime        int64  `json:"closeTime"`
	QuoteAssetVolume string `json:"quoteAssetVolume"`
}
