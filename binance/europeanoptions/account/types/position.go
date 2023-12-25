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
	bnutils "github.com/linstohu/nexapi/binance/utils"
	"github.com/linstohu/nexapi/utils"
)

type GetPositionInfoParam struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
}

type GetPositionInfoParams struct {
	GetPositionInfoParam
	bnutils.DefaultParam
}

type GetPositionInfoResp struct {
	Http *utils.ApiResponse
	Body []*Position
}

type Position struct {
	EntryPrice    string `json:"entryPrice"`
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Quantity      string `json:"quantity"`
	ReducibleQty  string `json:"reducibleQty"`
	MarkValue     string `json:"markValue"`
	Ror           string `json:"ror"`
	UnrealizedPNL string `json:"unrealizedPNL"`
	MarkPrice     string `json:"markPrice"`
	StrikePrice   string `json:"strikePrice"`
	PositionCost  string `json:"positionCost"`
	ExpiryDate    int64  `json:"expiryDate"`
	PriceScale    int    `json:"priceScale"`
	QuantityScale int    `json:"quantityScale"`
	OptionSide    string `json:"optionSide"`
	QuoteAsset    string `json:"quoteAsset"`
}
