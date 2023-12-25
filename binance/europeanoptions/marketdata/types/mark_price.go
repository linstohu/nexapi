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

import "github.com/linstohu/nexapi/utils"

type GetMarkPriceParam struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
}

type GetMarkPriceResp struct {
	Http *utils.ApiResponse
	Body []*MarkPrice
}

type MarkPrice struct {
	Symbol         string `json:"symbol"`
	MarkPrice      string `json:"markPrice"`
	BidIV          string `json:"bidIV"`
	AskIV          string `json:"askIV"`
	MarkIV         string `json:"markIV"`
	Delta          string `json:"delta"`
	Theta          string `json:"theta"`
	Gamma          string `json:"gamma"`
	Vega           string `json:"vega"`
	HighPriceLimit string `json:"highPriceLimit"`
	LowPriceLimit  string `json:"lowPriceLimit"`
}
