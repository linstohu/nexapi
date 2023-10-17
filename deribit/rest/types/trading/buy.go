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

package trading

type BuyParams struct {
	InstrumentName string  `json:"instrument_name"`
	Amount         float64 `json:"amount"`
	Type           string  `json:"type,omitempty"`
	Label          string  `json:"label,omitempty"`
	Price          float64 `json:"price,omitempty"`
	TimeInForce    string  `json:"time_in_force,omitempty"`
	MaxShow        float64 `json:"max_show,omitempty"`
	PostOnly       bool    `json:"post_only,omitempty"`
	ReduceOnly     bool    `json:"reduce_only,omitempty"`
	Trigger        string  `json:"trigger,omitempty"`
	Advanced       string  `json:"advanced,omitempty"`
}

type BuyResponse struct {
	Trades []Trade `json:"trades"`
	Order  Order   `json:"order"`
}
