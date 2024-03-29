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

type GetTickerPriceForSymbolParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type GetTickerPriceForSymbolsParam struct {
	Symbols []string `url:"symbols" validate:"required"`
}

type TickerPriceParams struct {
	Symbol  string `url:"symbol,omitempty" validate:"omitempty"`
	Symbols string `url:"symbols,omitempty" validate:"omitempty"`
}

type GetTickerPriceForSymbolResp struct {
	Http *utils.ApiResponse
	Body *TickerPrice
}

type GetTickerPriceForSymbolsResp struct {
	Http *utils.ApiResponse
	Body []*TickerPrice
}

type TickerPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
