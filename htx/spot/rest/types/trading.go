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

type NewOrderParam struct {
	AccountID        string `json:"account-id" validate:"required"`
	Symbol           string `json:"symbol" validate:"required"`
	Type             string `json:"type" validate:"required"`
	Amount           string `json:"amount" validate:"required"`
	Price            string `json:"price,omitempty" validate:"omitempty"`
	Source           string `json:"source,omitempty" validate:"omitempty"`
	ClientOrderID    string `json:"client-order-id,omitempty" validate:"omitempty"`
	SelfMatchPrevent int    `json:"self-match-prevent,omitempty" validate:"omitempty"`
	StopPrice        string `json:"stop-price,omitempty" validate:"omitempty"`
	Operator         string `json:"operator,omitempty" validate:"omitempty"`
}

type NewOrderResp struct {
	htxutils.V1Response
	Data string `json:"data,omitempty"`
}

type CancelOrderParams struct {
	OrderID string `json:"order-id" validate:"required"`
	Symbol  string `json:"symbol,omitempty" validate:"omitempty"`
}

type CancelOrderParam struct {
	Symbol string `json:"symbol,omitempty" validate:"omitempty"`
}

type CancelOrderResp struct {
	htxutils.V1Response
	Data string `json:"data,omitempty"`
}

type GetOpenOrdersParam struct {
	AccountID string `url:"account-id,omitempty" validate:"omitempty"`
	Symbol    string `url:"symbol,omitempty" validate:"omitempty"`
	Side      string `url:"side,omitempty" validate:"omitempty"`
	Types     string `url:"types,omitempty" validate:"omitempty"`
	From      string `url:"from,omitempty" validate:"omitempty"`
	Direct    string `url:"direct,omitempty" validate:"omitempty"`
	Size      string `url:"size,omitempty" validate:"omitempty"`
}

type GetOpenOrdersParams struct {
	GetOpenOrdersParam
	htxutils.DefaultAuthParam
}

type GetOpenOrdersResp struct {
	htxutils.V1Response
	Data []OpenOrder `json:"data,omitempty"`
}

type OpenOrder struct {
	Symbol           string `json:"symbol,omitempty"`
	Source           string `json:"source,omitempty"`
	Price            string `json:"price,omitempty"`
	CreatedAt        int64  `json:"created-at,omitempty"`
	Amount           string `json:"amount,omitempty"`
	AccountID        int    `json:"account-id,omitempty"`
	FilledCashAmount string `json:"filled-cash-amount,omitempty"`
	ClientOrderID    string `json:"client-order-id,omitempty"`
	FilledAmount     string `json:"filled-amount,omitempty"`
	FilledFees       string `json:"filled-fees,omitempty"`
	ID               int64  `json:"id,omitempty"`
	State            string `json:"state,omitempty"`
	Type             string `json:"type,omitempty"`
}

type CancelOrdersParam struct {
	AccountID string `json:"account-id,omitempty" validate:"omitempty"`
	Symbol    string `json:"symbol,omitempty" validate:"omitempty"`
	Types     string `json:"types,omitempty" validate:"omitempty"`
	Side      string `json:"side,omitempty" validate:"omitempty"`
	Size      int    `json:"size,omitempty" validate:"omitempty"`
}

type CancelOrdersResp struct {
	htxutils.V1Response
	Data struct {
		SuccessCount int   `json:"success-count,omitempty"`
		FailedCount  int   `json:"failed-count,omitempty"`
		NextID       int64 `json:"next-id,omitempty"`
	} `json:"data,omitempty"`
}

type SearchMatchResultsParams struct {
	SearchMatchResultsParam
	htxutils.DefaultAuthParam
}

type SearchMatchResultsParam struct {
	Symbol    string `url:"symbol" validate:"required"`
	Types     string `url:"types,omitempty" validate:"omitempty"`
	StartTime int64  `url:"start-time,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"end-time,omitempty" validate:"omitempty"`
	From      string `url:"from,omitempty" validate:"omitempty"`
	Direct    string `url:"direct,omitempty" validate:"omitempty"`
	Size      string `url:"size,omitempty" validate:"omitempty"`
}

type SearchMatchResultsResp struct {
	htxutils.V1Response
	Data []MatchResult `json:"data,omitempty"`
}

type MatchResult struct {
	Symbol            string `json:"symbol,omitempty"`
	FeeCurrency       string `json:"fee-currency,omitempty"`
	Source            string `json:"source,omitempty"`
	Price             string `json:"price,omitempty"`
	CreatedAt         int64  `json:"created-at,omitempty"`
	Role              string `json:"role,omitempty"`
	OrderID           int64  `json:"order-id,omitempty"`
	MatchID           int64  `json:"match-id,omitempty"`
	TradeID           int64  `json:"trade-id,omitempty"`
	FilledAmount      string `json:"filled-amount,omitempty"`
	FilledFees        string `json:"filled-fees,omitempty"`
	FilledPoints      string `json:"filled-points,omitempty"`
	FeeDeductCurrency string `json:"fee-deduct-currency,omitempty"`
	FeeDeductState    string `json:"fee-deduct-state,omitempty"`
	ID                int64  `json:"id,omitempty"`
	Type              string `json:"type,omitempty"`
}
