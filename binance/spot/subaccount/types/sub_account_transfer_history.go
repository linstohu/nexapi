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

type GetSubAccountTransferHistoryParam struct {
	Asset     string `url:"asset,omitempty" validate:"omitempty"`
	Type      int    `url:"type,omitempty" validate:"omitempty"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty"`
}

type GetSubAccountTransferHistoryParams struct {
	GetSubAccountTransferHistoryParam
	bnutils.DefaultParam
}

type GetSubAccountTransferHistoryResp struct {
	Http *utils.ApiResponse
	Body []*SubAccountTransferHistory
}

type SubAccountTransferHistory struct {
	CounterParty    string `json:"counterParty"`
	Email           string `json:"email"`
	Type            int    `json:"type"`
	Asset           string `json:"asset"`
	Qty             string `json:"qty"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Status          string `json:"status"`
	TranID          int64  `json:"tranId"`
	Time            int64  `json:"time"`
}
