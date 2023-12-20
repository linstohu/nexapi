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
	htxutils "github.com/linstohu/nexapi/htx/utils"
)

type GetAccountInfoResponse struct {
	Status string        `json:"status"`
	Data   []AccountInfo `json:"data"`
}
type AccountInfo struct {
	Id      int64  `json:"id"`
	Type    string `json:"type"`
	Subtype string `json:"subtype"`
	State   string `json:"state"`
}

type GetAccountValuationParam struct {
	AccountType       string `url:"accountType,omitempty" validate:"omitempty"`
	ValuationCurrency string `url:"valuationCurrency,omitempty" validate:"omitempty"`
}

type GetAccountValuationParams struct {
	GetAccountValuationParam
	htxutils.DefaultAuthParam
}

type GetAccountValuationResp struct {
	htxutils.V2Response
	Data AccountValue `json:"data"`
}

type Updated struct {
	Success bool  `json:"success,omitempty"`
	Time    int64 `json:"time,omitempty"`
}

type ProfitAccountBalanceList struct {
	DistributionType string  `json:"distributionType,omitempty"`
	Balance          float64 `json:"balance,omitempty"`
	Success          bool    `json:"success,omitempty"`
	AccountBalance   string  `json:"accountBalance,omitempty"`
}

type AccountValue struct {
	Updated                  Updated                    `json:"updated,omitempty"`
	TodayProfitRate          string                     `json:"todayProfitRate,omitempty"`
	TotalBalance             string                     `json:"totalBalance,omitempty"`
	TodayProfit              string                     `json:"todayProfit,omitempty"`
	ProfitAccountBalanceList []ProfitAccountBalanceList `json:"profitAccountBalanceList,omitempty"`
}
