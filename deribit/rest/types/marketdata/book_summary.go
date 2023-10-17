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

type BookSummary struct {
	Volume            float64 `json:"volume"`
	UnderlyingPrice   float64 `json:"underlying_price"`
	UnderlyingIndex   string  `json:"underlying_index"`
	QuoteCurrency     string  `json:"quote_currency"`
	OpenInterest      float64 `json:"open_interest"`
	MidPrice          float64 `json:"mid_price"`
	MarkPrice         float64 `json:"mark_price"`
	Low               float64 `json:"low"`
	Last              float64 `json:"last"`
	InterestRate      float64 `json:"interest_rate"`
	InstrumentName    string  `json:"instrument_name"`
	High              float64 `json:"high"`
	CreationTimestamp int64   `json:"creation_timestamp"`
	BidPrice          float64 `json:"bid_price"`
	BaseCurrency      string  `json:"base_currency"`
	AskPrice          float64 `json:"ask_price"`
}

type GetBookSummaryByCurrencyParams struct {
	Currency string `json:"currency"`
	Kind     string `json:"kind,omitempty"`
}

type GetBookSummaryByInstrumentParams struct {
	InstrumentName string `json:"instrument_name"`
}
