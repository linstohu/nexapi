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

type GetApiRestrictionsResp struct {
	Http *utils.ApiResponse
	Body *ApiRestrictions
}

type ApiRestrictions struct {
	IPRestrict                     bool  `json:"ipRestrict"`
	CreateTime                     int64 `json:"createTime"`
	EnableWithdrawals              bool  `json:"enableWithdrawals"`
	EnableInternalTransfer         bool  `json:"enableInternalTransfer"`
	PermitsUniversalTransfer       bool  `json:"permitsUniversalTransfer"`
	EnableVanillaOptions           bool  `json:"enableVanillaOptions"`
	EnableReading                  bool  `json:"enableReading"`
	EnableFutures                  bool  `json:"enableFutures"`
	EnableMargin                   bool  `json:"enableMargin"`
	EnableSpotAndMarginTrading     bool  `json:"enableSpotAndMarginTrading"`
	TradingAuthorityExpirationTime int64 `json:"tradingAuthorityExpirationTime"`
}
