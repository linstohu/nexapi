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
	htxutils "github.com/linstohu/nexapi/htx/utils"
)

type GetMergedMarketTickerParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type GetMergedMarketTickerResp struct {
	htxutils.V1Response
	Ts   int64 `json:"ts"`
	Tick Tick  `json:"tick,omitempty"`
}

type Tick struct {
	ID      int64     `json:"id,omitempty"`
	Version int64     `json:"version,omitempty"`
	Open    float64   `json:"open,omitempty"`
	Close   float64   `json:"close,omitempty"`
	Low     float64   `json:"low,omitempty"`
	High    float64   `json:"high,omitempty"`
	Amount  float64   `json:"amount,omitempty"`
	Vol     float64   `json:"vol,omitempty"`
	Count   int       `json:"count,omitempty"`
	Bid     []float64 `json:"bid,omitempty"`
	Ask     []float64 `json:"ask,omitempty"`
}
