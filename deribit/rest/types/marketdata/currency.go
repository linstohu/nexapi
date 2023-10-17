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

package marketdata

type Currency struct {
	CoinType                       string               `json:"coin_type"`
	Currency                       string               `json:"currency"`
	CurrencyLong                   string               `json:"currency_long"`
	DisabledDepositAddressCreation bool                 `json:"disabled_deposit_address_creation"`
	FeePrecision                   int                  `json:"fee_precision"`
	MinConfirmations               int                  `json:"min_confirmations"`
	MinWithdrawalFee               float64              `json:"min_withdrawal_fee"`
	WithdrawalFee                  float64              `json:"withdrawal_fee"`
	WithdrawalPriorities           []WithdrawalPriority `json:"withdrawal_priorities"`
}

type WithdrawalPriority struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}
