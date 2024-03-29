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

type GetBalanceParam struct {
	Asset string `url:"asset" validate:"required"`
}

type GetBalanceParams struct {
	GetBalanceParam
	bnutils.DefaultParam
}

type GetBalanceResp struct {
	Http *utils.ApiResponse
	Body *Balance
}

type GetAllBalanceParam struct {
	Asset string `url:"asset"`
}

type GetAllBalanceParams struct {
	GetAllBalanceParam
	bnutils.DefaultParam
}

type GetAllBalanceResp struct {
	Http *utils.ApiResponse
	Body []*Balance
}

type Balance struct {
	Asset               string `json:"asset"`
	TotalWalletBalance  string `json:"totalWalletBalance"`
	CrossMarginAsset    string `json:"crossMarginAsset"`
	CrossMarginBorrowed string `json:"crossMarginBorrowed"`
	CrossMarginFree     string `json:"crossMarginFree"`
	CrossMarginInterest string `json:"crossMarginInterest"`
	CrossMarginLocked   string `json:"crossMarginLocked"`
	CrossMarginNetAsset string `json:"crossMarginNetAsset"`
	UmWalletBalance     string `json:"umWalletBalance"`
	UmUnrealizedPNL     string `json:"umUnrealizedPNL"`
	CmWalletBalance     string `json:"cmWalletBalance"`
	CmUnrealizedPNL     string `json:"cmUnrealizedPNL"`
	UpdateTime          int64  `json:"updateTime"`
}
