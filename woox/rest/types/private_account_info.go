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

type GetAccountInfo struct {
	Response
	Data struct {
		ApplicationID          string  `json:"applicationId"`
		Account                string  `json:"account"`
		Alias                  string  `json:"alias"`
		AccountMode            string  `json:"accountMode"`
		Leverage               int     `json:"leverage"`
		TakerFeeRate           float64 `json:"takerFeeRate"`
		MakerFeeRate           float64 `json:"makerFeeRate"`
		InterestRate           float64 `json:"interestRate"`
		FuturesTakerFeeRate    float64 `json:"futuresTakerFeeRate"`
		FuturesMakerFeeRate    float64 `json:"futuresMakerFeeRate"`
		Otpauth                bool    `json:"otpauth"`
		MarginRatio            float64 `json:"marginRatio"`
		OpenMarginRatio        float64 `json:"openMarginRatio"`
		InitialMarginRatio     float64 `json:"initialMarginRatio"`
		MaintenanceMarginRatio float64 `json:"maintenanceMarginRatio"`
		TotalCollateral        float64 `json:"totalCollateral"`
		FreeCollateral         float64 `json:"freeCollateral"`
		TotalAccountValue      float64 `json:"totalAccountValue"`
		TotalVaultValue        float64 `json:"totalVaultValue"`
		TotalStakingValue      float64 `json:"totalStakingValue"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}
