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

type UniversalTransferParam struct {
	Type       TransferType `url:"type" validate:"required"`
	Asset      string       `url:"asset" validate:"required"`
	Amount     float64      `url:"amount" validate:"required"`
	FromSymbol string       `url:"fromSymbol,omitempty" validate:"omitempty"`
	ToSymbol   string       `url:"toSymbol,omitempty" validate:"omitempty"`
}

type TransferType string

const (
	MAIN_UMFUTURE                 TransferType = "MAIN_UMFUTURE"
	MAIN_CMFUTURE                 TransferType = "MAIN_CMFUTURE"
	MAIN_MARGIN                   TransferType = "MAIN_MARGIN"
	UMFUTURE_MAIN                 TransferType = "UMFUTURE_MAIN"
	UMFUTURE_MARGIN               TransferType = "UMFUTURE_MARGIN"
	CMFUTURE_MAIN                 TransferType = "CMFUTURE_MAIN"
	CMFUTURE_MARGIN               TransferType = "CMFUTURE_MARGIN"
	MARGIN_MAIN                   TransferType = "MARGIN_MAIN"
	MARGIN_UMFUTURE               TransferType = "MARGIN_UMFUTURE"
	MARGIN_CMFUTURE               TransferType = "MARGIN_CMFUTURE"
	ISOLATEDMARGIN_MARGIN         TransferType = "ISOLATEDMARGIN_MARGIN"
	MARGIN_ISOLATEDMARGIN         TransferType = "MARGIN_ISOLATEDMARGIN"
	ISOLATEDMARGIN_ISOLATEDMARGIN TransferType = "ISOLATEDMARGIN_ISOLATEDMARGIN"
	MAIN_FUNDING                  TransferType = "MAIN_FUNDING"
	FUNDING_MAIN                  TransferType = "FUNDING_MAIN"
	FUNDING_UMFUTURE              TransferType = "FUNDING_UMFUTURE"
	UMFUTURE_FUNDING              TransferType = "UMFUTURE_FUNDING"
	MARGIN_FUNDING                TransferType = "MARGIN_FUNDING"
	FUNDING_MARGIN                TransferType = "FUNDING_MARGIN"
	FUNDING_CMFUTURE              TransferType = "FUNDING_CMFUTURE"
	CMFUTURE_FUNDING              TransferType = "CMFUTURE_FUNDING"
	MAIN_OPTION                   TransferType = "MAIN_OPTION"
	OPTION_MAIN                   TransferType = "OPTION_MAIN"
	UMFUTURE_OPTION               TransferType = "UMFUTURE_OPTION"
	OPTION_UMFUTURE               TransferType = "OPTION_UMFUTURE"
	MARGIN_OPTION                 TransferType = "MARGIN_OPTION"
	OPTION_MARGIN                 TransferType = "OPTION_MARGIN"
	FUNDING_OPTION                TransferType = "FUNDING_OPTION"
	OPTION_FUNDING                TransferType = "OPTION_FUNDING"
	MAIN_PORTFOLIO_MARGIN         TransferType = "MAIN_PORTFOLIO_MARGIN"
	PORTFOLIO_MARGIN_MAIN         TransferType = "PORTFOLIO_MARGIN_MAIN"
)

type UniversalTransferParams struct {
	UniversalTransferParam
	bnutils.DefaultParam
}

type UniversalTransferResp struct {
	Http *utils.ApiResponse
	Body *UniversalTransferAPIResp
}

type UniversalTransferAPIResp struct {
	TranID int64 `json:"tranId"`
}

type GetUniversalTransferHistoryParam struct {
	Type       TransferType `url:"type" validate:"required"`
	StartTime  int64        `url:"startTime,omitempty" validate:"omitempty"`
	EndTime    int64        `url:"endTime,omitempty" validate:"omitempty"`
	Current    int          `url:"current,omitempty" validate:"omitempty"`
	Size       int          `url:"size,omitempty" validate:"omitempty"`
	FromSymbol string       `url:"fromSymbol,omitempty" validate:"omitempty"`
	ToSymbol   string       `url:"toSymbol,omitempty" validate:"omitempty"`
}

type GetUniversalTransferHistoryParams struct {
	GetUniversalTransferHistoryParam
	bnutils.DefaultParam
}

type GetUniversalTransferHistoryResp struct {
	Http *utils.ApiResponse
	Body *GetUniversalTransferHistory
}

type GetUniversalTransferHistory struct {
	Total int `json:"total"`
	Rows  []struct {
		Asset     string `json:"asset"`
		Amount    string `json:"amount"`
		Type      string `json:"type"`
		Status    string `json:"status"`
		TranID    int64  `json:"tranId"`
		Timestamp int64  `json:"timestamp"`
	} `json:"rows"`
}
