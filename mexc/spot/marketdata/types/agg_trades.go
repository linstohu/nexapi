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

type GetAggTradesParam struct {
	Symbol    string `url:"symbol" validate:"required"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type AggTrade struct {
	A  int    `json:"a"` // Aggregate tradeId
	F  int    `json:"f"` // First tradeId
	L  int    `json:"l"` // Last tradeId
	P  string `json:"p"` // Price
	Q  string `json:"q"` // Quantity
	T  int64  `json:"T"` // Timestamp
	M  bool   `json:"m"` // Was the buyer the maker?
	Ma bool   `json:"M"` // Was the trade the best price match?
}
