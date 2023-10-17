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

type V1Position struct {
	Symbol           string  `json:"symbol"`
	Holding          float64 `json:"holding"`
	PendingLongQty   float64 `json:"pending_long_qty"`
	PendingShortQty  float64 `json:"pending_short_qty"`
	SettlePrice      float64 `json:"settle_price"`
	AverageOpenPrice float64 `json:"average_open_price"`
	Pnl24H           float64 `json:"pnl_24_h"`
	Fee24H           float64 `json:"fee_24_h"`
	MarkPrice        float64 `json:"mark_price"`
	EstLiqPrice      float64 `json:"est_liq_price"`
	Timestamp        string  `json:"timestamp"`
}

type GetOnePositionInfo struct {
	Response
	V1Position
}

type V3Position struct {
	Symbol           string  `json:"symbol"`
	Holding          float64 `json:"holding"`
	PendingLongQty   float64 `json:"pending_long_qty"`
	PendingShortQty  float64 `json:"pending_short_qty"`
	SettlePrice      float64 `json:"settle_price"`
	AverageOpenPrice float64 `json:"average_open_price"`
	Pnl24H           float64 `json:"pnl_24_h"`
	Fee24H           float64 `json:"fee_24_h"`
	MarkPrice        float64 `json:"mark_price"`
	EstLiqPrice      float64 `json:"est_liq_price"`
	Timestamp        float64 `json:"timestamp"`
}

type GetAllV3PositionInfo struct {
	Response
	Data struct {
		Positions []V3Position `json:"positions"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}
