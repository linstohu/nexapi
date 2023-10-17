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

package trading

type Settlement struct {
	IndexPrice        float64 `json:"index_price"`
	InstrumentName    string  `json:"instrument_name"`
	MarkPrice         float64 `json:"mark_price"`
	Position          float64 `json:"position"`
	ProfitLoss        float64 `json:"profit_loss"`
	SessionBankrupcy  float64 `json:"session_bankrupcy"`
	SessionProfitLoss float64 `json:"session_profit_loss"`
	SessionTax        float64 `json:"session_tax"`
	SessionTaxRate    float64 `json:"session_tax_rate"`
	Socialized        float64 `json:"socialized"`
	Timestamp         int64   `json:"timestamp"`
	Type              string  `json:"type"`
}

type GetSettlementHistoryResponse struct {
	Settlements  []Settlement `json:"settlements"`
	Continuation string       `json:"continuation"`
}

type GetSettlementHistoryByInstrumentParams struct {
	InstrumentName string `json:"instrument_name"`
	Type           string `json:"type,omitempty"`
	Count          int    `json:"count,omitempty"`
	Continuation   string `json:"continuation,omitempty"`
}

type GetSettlementHistoryByCurrencyParams struct {
	Currency     string `json:"currency"`
	Type         string `json:"type,omitempty"`
	Count        int    `json:"count,omitempty"`
	Continuation string `json:"continuation,omitempty"`
}
