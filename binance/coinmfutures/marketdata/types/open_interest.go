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
	cmutils "github.com/linstohu/nexapi/binance/coinmfutures/utils"
	"github.com/linstohu/nexapi/utils"
)

type GetOpenInterestParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type GetOpenInterestResp struct {
	Http *utils.ApiResponse
	Body *OpenInterest
}

type OpenInterest struct {
	Symbol       string `json:"symbol"`
	Pair         string `json:"pair"`
	OpenInterest string `json:"openInterest"`
	ContractType string `json:"contractType"`
	Time         int64  `json:"time"`
}

type GetOpenInterestHistParam struct {
	Pair         string               `url:"pair" validate:"required"`
	ContractType cmutils.ContractType `url:"contractType" validate:"required,oneof=ALL CURRENT_QUARTER NEXT_QUARTER PERPETUAL"`
	Period       string               `url:"period" validate:"required,oneof=5m 15m 30m 1h 2h 4h 6h 12h 1d"`
	StartTime    int64                `url:"startTime,omitempty" validate:"omitempty"`
	EndTime      int64                `url:"endTime,omitempty" validate:"omitempty"`
	Limit        int                  `url:"limit,omitempty" validate:"omitempty,max=500"`
}

type GetOpenInterestHistResp struct {
	Http *utils.ApiResponse
	Body []*OpenInterestHist
}

type OpenInterestHist struct {
	Pair                 string `json:"pair"`
	ContractType         string `json:"contractType"`
	SumOpenInterest      string `json:"sumOpenInterest"`
	SumOpenInterestValue string `json:"sumOpenInterestValue"`
	Timestamp            int64  `json:"timestamp"`
}
