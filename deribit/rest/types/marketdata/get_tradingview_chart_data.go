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

package marketdata

type GetTradingviewChartDataParams struct {
	InstrumentName string `url:"instrument_name,omitempty" json:"instrument_name"`
	StartTimestamp int64  `url:"start_timestamp,omitempty" json:"start_timestamp"`
	EndTimestamp   int64  `url:"end_timestamp,omitempty" json:"end_timestamp"`
	Resolution     string `url:"resolution,omitempty" json:"resolution"`
}

type GetTradingviewChartDataResponse struct {
	Close  []float64 `json:"close"`
	Cost   []float64 `json:"cost"`
	High   []float64 `json:"high"`
	Low    []float64 `json:"low"`
	Open   []float64 `json:"open"`
	Volume []float64 `json:"volume"`
	Ticks  []int64   `json:"ticks"`
	Status string    `json:"status"`
}
