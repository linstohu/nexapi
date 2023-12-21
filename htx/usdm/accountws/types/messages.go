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

type IsoAccount struct {
	Op    string    `json:"op,omitempty"`
	Topic string    `json:"topic,omitempty"`
	Ts    int64     `json:"ts,omitempty"`
	Event string    `json:"event,omitempty"`
	Data  []IsoData `json:"data,omitempty"`
	UID   string    `json:"uid,omitempty"`
}

type IsoData struct {
	Symbol            string  `json:"symbol,omitempty"`
	ContractCode      string  `json:"contract_code,omitempty"`
	MarginBalance     float64 `json:"margin_balance,omitempty"`
	MarginStatic      float64 `json:"margin_static,omitempty"`
	MarginPosition    float64 `json:"margin_position,omitempty"`
	MarginFrozen      float64 `json:"margin_frozen,omitempty"`
	MarginAvailable   float64 `json:"margin_available,omitempty"`
	ProfitReal        float64 `json:"profit_real,omitempty"`
	ProfitUnreal      float64 `json:"profit_unreal,omitempty"`
	WithdrawAvailable float64 `json:"withdraw_available,omitempty"`
	RiskRate          float64 `json:"risk_rate,omitempty"`
	LiquidationPrice  float64 `json:"liquidation_price,omitempty"`
	LeverRate         int     `json:"lever_rate,omitempty"`
	AdjustFactor      float64 `json:"adjust_factor,omitempty"`
	MarginAsset       string  `json:"margin_asset,omitempty"`
	MarginMode        string  `json:"margin_mode,omitempty"`
	MarginAccount     string  `json:"margin_account,omitempty"`
	PositionMode      string  `json:"position_mode,omitempty"`
}

type CrossAccount struct {
	Op    string      `json:"op,omitempty"`
	Topic string      `json:"topic,omitempty"`
	Ts    int64       `json:"ts,omitempty"`
	Event string      `json:"event,omitempty"`
	Data  []CrossData `json:"data,omitempty"`
	UID   string      `json:"uid,omitempty"`
}

type ContractDetail struct {
	Symbol           string  `json:"symbol,omitempty"`
	ContractCode     string  `json:"contract_code,omitempty"`
	MarginPosition   float64 `json:"margin_position,omitempty"`
	MarginFrozen     float64 `json:"margin_frozen,omitempty"`
	MarginAvailable  float64 `json:"margin_available,omitempty"`
	ProfitUnreal     float64 `json:"profit_unreal,omitempty"`
	LiquidationPrice float64 `json:"liquidation_price,omitempty"`
	LeverRate        int     `json:"lever_rate,omitempty"`
	AdjustFactor     float64 `json:"adjust_factor,omitempty"`
	ContractType     string  `json:"contract_type,omitempty"`
	Pair             string  `json:"pair,omitempty"`
	BusinessType     string  `json:"business_type,omitempty"`
}

type FuturesContractDetail struct {
	Symbol           string  `json:"symbol,omitempty"`
	ContractCode     string  `json:"contract_code,omitempty"`
	MarginPosition   float64 `json:"margin_position,omitempty"`
	MarginFrozen     float64 `json:"margin_frozen,omitempty"`
	MarginAvailable  float64 `json:"margin_available,omitempty"`
	ProfitUnreal     float64 `json:"profit_unreal,omitempty"`
	LiquidationPrice float64 `json:"liquidation_price,omitempty"`
	LeverRate        int     `json:"lever_rate,omitempty"`
	AdjustFactor     float64 `json:"adjust_factor,omitempty"`
	ContractType     string  `json:"contract_type,omitempty"`
	Pair             string  `json:"pair,omitempty"`
	BusinessType     string  `json:"business_type,omitempty"`
}

type CrossData struct {
	MarginMode            string                  `json:"margin_mode,omitempty"`
	MarginAccount         string                  `json:"margin_account,omitempty"`
	MarginAsset           string                  `json:"margin_asset,omitempty"`
	MarginBalance         float64                 `json:"margin_balance,omitempty"`
	MarginStatic          float64                 `json:"margin_static,omitempty"`
	MarginPosition        float64                 `json:"margin_position,omitempty"`
	MarginFrozen          float64                 `json:"margin_frozen,omitempty"`
	ProfitReal            float64                 `json:"profit_real,omitempty"`
	ProfitUnreal          float64                 `json:"profit_unreal,omitempty"`
	WithdrawAvailable     float64                 `json:"withdraw_available,omitempty"`
	RiskRate              float64                 `json:"risk_rate,omitempty"`
	PositionMode          string                  `json:"position_mode,omitempty"`
	ContractDetail        []ContractDetail        `json:"contract_detail,omitempty"`
	FuturesContractDetail []FuturesContractDetail `json:"futures_contract_detail,omitempty"`
}

type UnifyAccount struct {
	Op    string             `json:"op,omitempty"`
	Topic string             `json:"topic,omitempty"`
	Ts    int64              `json:"ts,omitempty"`
	Event string             `json:"event,omitempty"`
	Data  []UnifyAccountData `json:"data,omitempty"`
	UID   string             `json:"uid,omitempty"`
}

type CrossSwap struct {
	Symbol            string  `json:"symbol,omitempty"`
	ContractCode      string  `json:"contract_code,omitempty"`
	MarginMode        string  `json:"margin_mode,omitempty"`
	MarginAvailable   float64 `json:"margin_available,omitempty"`
	LeverRate         float64 `json:"lever_rate,omitempty"`
	ContractType      string  `json:"contract_type,omitempty"`
	BusinessType      string  `json:"business_type,omitempty"`
	CrossMaxAvailable float64 `json:"cross_max_available,omitempty"`
}

type CrossFuture struct {
	Symbol            string  `json:"symbol,omitempty"`
	ContractCode      string  `json:"contract_code,omitempty"`
	MarginMode        string  `json:"margin_mode,omitempty"`
	MarginAvailable   float64 `json:"margin_available,omitempty"`
	LeverRate         float64 `json:"lever_rate,omitempty"`
	ContractType      string  `json:"contract_type,omitempty"`
	BusinessType      string  `json:"business_type,omitempty"`
	CrossMaxAvailable float64 `json:"cross_max_available,omitempty"`
}

type IsolatedSwap struct {
	Symbol            string  `json:"symbol,omitempty"`
	ContractCode      string  `json:"contract_code,omitempty"`
	MarginMode        string  `json:"margin_mode,omitempty"`
	MarginAvailable   float64 `json:"margin_available,omitempty"`
	WithdrawAvailable float64 `json:"withdraw_available,omitempty"`
	LeverRate         int     `json:"lever_rate,omitempty"`
	PositionMode      string  `json:"position_mode,omitempty"`
}

type UnifyAccountData struct {
	MarginAsset       string         `json:"margin_asset,omitempty"`
	MarginStatic      float64        `json:"margin_static,omitempty"`
	CrossMarginStatic float64        `json:"cross_margin_static,omitempty"`
	MarginBalance     float64        `json:"margin_balance,omitempty"`
	CrossProfitUnreal float64        `json:"cross_profit_unreal,omitempty"`
	MarginFrozen      float64        `json:"margin_frozen,omitempty"`
	WithdrawAvailable float64        `json:"withdraw_available,omitempty"`
	CrossRiskRate     float64        `json:"cross_risk_rate,omitempty"`
	CrossSwap         []CrossSwap    `json:"cross_swap,omitempty"`
	CrossFuture       []CrossFuture  `json:"cross_future,omitempty"`
	IsolatedSwap      []IsolatedSwap `json:"isolated_swap,omitempty"`
}
