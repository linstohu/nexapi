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

type GetWalletBalanceParam struct {
	AccountType AccountType `url:"accountType" validate:"required,oneof=UNIFIED SPOT OPTION CONTRACT FUND"` //
}

// AccountType
// doc: https://bybit-exchange.github.io/docs/v5/account/wallet-balance
//   - Unified account: UNIFIED (trade spot/linear/options), CONTRACT(trade inverse)
//   - Classic account: CONTRACT, SPOT
//
// AccountType Enum doc: https://bybit-exchange.github.io/docs/v5/enum#accounttype
type AccountType = string

const (
	UNIFIED  = "UNIFIED"
	SPOT     = "SPOT"
	OPTION   = "OPTION"
	CONTRACT = "CONTRACT"
	FUND     = "FUND"
)

type GetWalletBalanceResp struct {
	Http *utils.ApiResponse
	Body *GetWalletBalanceAPIResp
}

type GetWalletBalanceAPIResp struct {
	Response `json:",inline"`
	Result   WalletBalanceResult `json:"result"`
}

type WalletBalanceResult struct {
	List []WalletBalanceList `json:"list"`
}

type WalletBalanceList struct {
	TotalEquity            string              `json:"totalEquity"`
	AccountIMRate          string              `json:"accountIMRate"`
	TotalMarginBalance     string              `json:"totalMarginBalance"`
	TotalInitialMargin     string              `json:"totalInitialMargin"`
	AccountType            string              `json:"accountType"`
	TotalAvailableBalance  string              `json:"totalAvailableBalance"`
	AccountMMRate          string              `json:"accountMMRate"`
	TotalPerpUPL           string              `json:"totalPerpUPL"`
	TotalWalletBalance     string              `json:"totalWalletBalance"`
	TotalMaintenanceMargin string              `json:"totalMaintenanceMargin"`
	Coin                   []WalletBalanceCoin `json:"coin"`
}

type WalletBalanceCoin struct {
	AvailableToBorrow   string `json:"availableToBorrow"`
	AccruedInterest     string `json:"accruedInterest"`
	AvailableToWithdraw string `json:"availableToWithdraw"`
	TotalOrderIM        string `json:"totalOrderIM"`
	Equity              string `json:"equity"`
	TotalPositionMM     string `json:"totalPositionMM"`
	UsdValue            string `json:"usdValue"`
	UnrealisedPnl       string `json:"unrealisedPnl"`
	BorrowAmount        string `json:"borrowAmount"`
	TotalPositionIM     string `json:"totalPositionIM"`
	WalletBalance       string `json:"walletBalance"`
	CumRealisedPnl      string `json:"cumRealisedPnl"`
	Coin                string `json:"coin"`
}

type GetAccountBalanceParam struct {
	AccountType AccountType `url:"accountType" validate:"required,oneof=UNIFIED SPOT OPTION CONTRACT FUND"`
	WithBonus   string      `url:"accountType,omitempty"`
}

type GetAccountBalanceResp struct {
	Http *utils.ApiResponse
	Body *GetAccountBalanceAPIResp
}

type GetAccountBalanceAPIResp struct {
	Response `json:",inline"`
	Result   GetAccountBalanceResult `json:"result"`
}

type GetAccountBalanceResult struct {
	MemberID    string           `json:"memberId"`
	AccountType string           `json:"accountType"`
	Balance     []AccountBalance `json:"balance"`
}

type AccountBalance struct {
	Coin            string `json:"coin"`
	TransferBalance string `json:"transferBalance"`
	WalletBalance   string `json:"walletBalance"`
	Bonus           string `json:"bonus"`
}
