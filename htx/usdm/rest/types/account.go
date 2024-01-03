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

type GetAssetValuationParam struct {
	ValuationAsset string `json:"valuation_asset,omitempty" validate:"omitempty"`
}

type GetAssetValuationResp struct {
	DefaultResponse
	Data []AssetValuation `json:"data"`
}

type AssetValuation struct {
	ValuationAsset string `json:"valuation_asset"`
	Balance        string `json:"balance"`
}

type GetIsolatedAccountsParam struct {
	ContractCode string `json:"contract_code,omitempty" validate:"omitempty"`
}

type GetIsolatedAccountsResp struct {
	DefaultResponse
	Data []IsolatedAccount `json:"data"`
}

type IsolatedAccount struct {
	Symbol            string  `json:"symbol,omitempty"`
	ContractCode      string  `json:"contract_code,omitempty"`
	MarginAsset       string  `json:"margin_asset,omitempty"`
	MarginBalance     float64 `json:"margin_balance,omitempty"`
	MarginStatic      float64 `json:"margin_static,omitempty"`
	MarginPosition    float64 `json:"margin_position,omitempty"`
	MarginFrozen      float64 `json:"margin_frozen,omitempty"`
	MarginAvailable   float64 `json:"margin_available,omitempty"`
	ProfitUnreal      float64 `json:"profit_unreal,omitempty"`
	RiskRate          float64 `json:"risk_rate,omitempty"`
	NewRiskRate       float64 `json:"new_risk_rate,omitempty"`
	ProfitReal        float64 `json:"profit_real,omitempty"`
	TradePartition    string  `json:"trade_partition,omitempty"`
	LiquidationPrice  float64 `json:"liquidation_price,omitempty"`
	WithdrawAvailable float64 `json:"withdraw_available,omitempty"`
	LeverRate         float64 `json:"lever_rate,omitempty"`
	AdjustFactor      float64 `json:"adjust_factor,omitempty"`
	MarginMode        string  `json:"margin_mode,omitempty"`
	MarginAccount     string  `json:"margin_account,omitempty"`
	PositionMode      string  `json:"position_mode,omitempty"`
}

type GetCrossAccountsParam struct {
	MarginAccount string `json:"margin_account,omitempty" validate:"omitempty"`
}

type GetCrossAccountsResp struct {
	DefaultResponse
	Data []CrossAccount `json:"data"`
}

type CrossAccount struct {
	FuturesContractDetail []FuturesContractDetail `json:"futures_contract_detail,omitempty"`
	MarginMode            string                  `json:"margin_mode,omitempty"`
	MarginAccount         string                  `json:"margin_account,omitempty"`
	MarginAsset           string                  `json:"margin_asset,omitempty"`
	MarginBalance         float64                 `json:"margin_balance,omitempty"`
	MarginStatic          float64                 `json:"margin_static,omitempty"`
	MarginPosition        float64                 `json:"margin_position,omitempty"`
	MarginFrozen          float64                 `json:"margin_frozen,omitempty"`
	ProfitUnreal          float64                 `json:"profit_unreal,omitempty"`
	WithdrawAvailable     float64                 `json:"withdraw_available,omitempty"`
	RiskRate              float64                 `json:"risk_rate,omitempty"`
	MoneyIn               float64                 `json:"money_in,omitempty"`
	MoneyOut              float64                 `json:"money_out,omitempty"`
	NewRiskRate           float64                 `json:"new_risk_rate,omitempty"`
	PositionMode          string                  `json:"position_mode,omitempty"`
	ContractDetail        []ContractDetail        `json:"contract_detail,omitempty"`
}

type FuturesContractDetail struct {
	Symbol            string  `json:"symbol,omitempty"`
	ContractCode      string  `json:"contract_code,omitempty"`
	MarginPosition    float64 `json:"margin_position,omitempty"`
	MarginFrozen      float64 `json:"margin_frozen,omitempty"`
	MarginAvailable   float64 `json:"margin_available,omitempty"`
	ProfitUnreal      float64 `json:"profit_unreal,omitempty"`
	LiquidationPrice  float64 `json:"liquidation_price,omitempty"`
	LeverRate         float64 `json:"lever_rate,omitempty"`
	AdjustFactor      float64 `json:"adjust_factor,omitempty"`
	ContractType      string  `json:"contract_type,omitempty"`
	CrossMaxAvailable float64 `json:"cross_max_available,omitempty"`
	TradePartition    string  `json:"trade_partition,omitempty"`
	Pair              string  `json:"pair,omitempty"`
	BusinessType      string  `json:"business_type,omitempty"`
}

type ContractDetail struct {
	Symbol            string  `json:"symbol,omitempty"`
	ContractCode      string  `json:"contract_code,omitempty"`
	MarginPosition    float64 `json:"margin_position,omitempty"`
	MarginFrozen      float64 `json:"margin_frozen,omitempty"`
	MarginAvailable   float64 `json:"margin_available,omitempty"`
	ProfitUnreal      float64 `json:"profit_unreal,omitempty"`
	LiquidationPrice  float64 `json:"liquidation_price,omitempty"`
	LeverRate         float64 `json:"lever_rate,omitempty"`
	AdjustFactor      float64 `json:"adjust_factor,omitempty"`
	ContractType      string  `json:"contract_type,omitempty"`
	CrossMaxAvailable float64 `json:"cross_max_available,omitempty"`
	TradePartition    string  `json:"trade_partition,omitempty"`
	Pair              string  `json:"pair,omitempty"`
	BusinessType      string  `json:"business_type,omitempty"`
}

type GetUnifiedAccountsParam struct {
	ContractCode string `json:"contract_code,omitempty" validate:"omitempty"`
}

type GetUnifiedAccountsResp struct {
	htxutils.V2Response
	Data []UnifiedAccount `json:"data"`
}

type CrossFuture struct {
	BusinessType      string  `json:"business_type,omitempty"`
	ContractCode      string  `json:"contract_code,omitempty"`
	ContractType      string  `json:"contract_type,omitempty"`
	CrossMaxAvailable float64 `json:"cross_max_available,omitempty"`
	LeverRate         float64 `json:"lever_rate,omitempty"`
	MarginAvailable   float64 `json:"margin_available,omitempty"`
	MarginMode        string  `json:"margin_mode,omitempty"`
	Symbol            string  `json:"symbol,omitempty"`
}

type CrossSwap struct {
	BusinessType      string  `json:"business_type,omitempty"`
	ContractCode      string  `json:"contract_code,omitempty"`
	ContractType      string  `json:"contract_type,omitempty"`
	CrossMaxAvailable float64 `json:"cross_max_available,omitempty"`
	LeverRate         float64 `json:"lever_rate,omitempty"`
	MarginAvailable   float64 `json:"margin_available,omitempty"`
	MarginMode        string  `json:"margin_mode,omitempty"`
	Symbol            string  `json:"symbol,omitempty"`
}

type UnifiedAccount struct {
	CrossMarginStatic float64       `json:"cross_margin_static,omitempty"`
	CrossProfitUnreal float64       `json:"cross_profit_unreal,omitempty"`
	CrossRiskRate     float64       `json:"cross_risk_rate,omitempty"`
	MarginAsset       string        `json:"margin_asset,omitempty"`
	MarginBalance     float64       `json:"margin_balance,omitempty"`
	MarginFrozen      float64       `json:"margin_frozen,omitempty"`
	MarginStatic      float64       `json:"margin_static,omitempty"`
	UserID            float64       `json:"userId,omitempty"`
	WithdrawAvailable float64       `json:"withdraw_available,omitempty"`
	CrossFuture       []CrossFuture `json:"cross_future,omitempty"`
	CrossSwap         []CrossSwap   `json:"cross_swap,omitempty"`
}
