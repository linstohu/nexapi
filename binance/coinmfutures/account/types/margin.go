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
	umutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
	"github.com/linstohu/nexapi/utils"
)

type ChangeMarginTypeParam struct {
	Symbol     string             `url:"symbol" validate:"required"`
	MarginType umutils.MarginType `url:"marginType,omitempty" validate:"omitempty"`
}

type ChangeMarginTypeParams struct {
	ChangeMarginTypeParam
	bnutils.DefaultParam
}

type ModifyIsolatedPositionMarginParam struct {
	Symbol       string               `url:"symbol" validate:"required"`
	PositionSide umutils.PositionSide `url:"positionSide,omitempty" validate:"omitempty"`
	Amount       float64              `url:"amount" validate:"required"`
	Type         int                  `url:"type,omitempty" validate:"required"`
}

type ModifyIsolatedPositionMarginParams struct {
	ModifyIsolatedPositionMarginParam
	bnutils.DefaultParam
}

type ModifyIsolatedPositionMarginResp struct {
	Http *utils.ApiResponse
	Body *ModifyIsolatedPositionMarginAPIResp
}

type ModifyIsolatedPositionMarginAPIResp struct {
	Amount float64 `json:"amount"`
	Code   int     `json:"code"`
	Msg    string  `json:"msg"`
	Type   int     `json:"type"`
}
