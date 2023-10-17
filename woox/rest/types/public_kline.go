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

type GetKlineParam struct {
	Symbol string `url:"symbol" validate:"required"`
	Type   string `url:"type" validate:"required,oneof=1m 5m 15m 30m 1h 4h 12h 1d 1w 1mon 1y"`
	Limit  int    `url:"limit,omitempty" validate:"max=1000"`
}

type Kline struct {
	Response
	Rows []struct {
		Open           float64 `json:"open"`
		Close          float64 `json:"close"`
		Low            float64 `json:"low"`
		High           float64 `json:"high"`
		Volume         float64 `json:"volume"`
		Amount         float64 `json:"amount"`
		Symbol         string  `json:"symbol"`
		Type           string  `json:"type"`
		StartTimestamp int64   `json:"start_timestamp"`
		EndTimestamp   int64   `json:"end_timestamp"`
	} `json:"rows"`
}
