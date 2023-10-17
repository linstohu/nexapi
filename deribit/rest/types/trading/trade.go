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

type Trade struct {
	Advanced        string  `json:"advanced"`
	Amount          float64 `json:"amount"`
	API             bool    `json:"api"`
	BlockTradeId    string  `json:"block_trade_id"`
	Direction       string  `json:"direction"`
	Fee             float64 `json:"fee"`
	FeeCurrency     string  `json:"fee_currency"`
	IndexPrice      float64 `json:"index_price"`
	InstrumentName  string  `json:"instrument_name"`
	IV              float64 `json:"iv,omitempty"`
	Label           string  `json:"label"`
	Liquidation     string  `json:"liquidation"`
	Liquidity       string  `json:"liquidity"`
	MarkPrice       float64 `json:"mark_price"`
	MMP             bool    `json:"mmp"`
	OrderId         string  `json:"order_id"`
	OrderType       string  `json:"order_type"`
	Price           float64 `json:"price"`
	ProfitLoss      float64 `json:"profit_loss"`
	State           string  `json:"state"`
	Timestamp       int64   `json:"timestamp"`
	TradeID         string  `json:"trade_id"`
	TradeSeq        int64   `json:"trade_seq"`
	UnderlyingPrice float64 `json:"underlying_price,omitempty"`
}
