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
)

type GetFundingFlowParam struct {
	Currency  string `url:"currency" validate:"required"`
	RecordID  int64  `url:"recordId,omitempty" validate:"omitempty"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type GetFundingFlowParams struct {
	GetFundingFlowParam
	bnutils.DefaultParam
}

type FundingFlow struct {
	ID         string `json:"id"`
	Asset      string `json:"asset"`
	Amount     string `json:"amount"`
	Type       string `json:"type"`
	CreateDate int64  `json:"createDate"`
}
