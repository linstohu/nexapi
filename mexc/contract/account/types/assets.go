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

type GetAccountAsset struct {
	Response
	Data []*ContractAsset `json:"data"`
}

type ContractAsset struct {
	Currency         string  `json:"currency"`
	PositionMargin   float64 `json:"positionMargin"`
	AvailableBalance float64 `json:"availableBalance"`
	CashBalance      float64 `json:"cachBalance"`
	FrozenBalance    float64 `json:"frozenBalance"`
	Equity           float64 `json:"equity"`
	Unrealized       float64 `json:"unrealized"`
	Bonus            float64 `json:"bonus"`
}
