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

import "strconv"

type Order struct {
	RejectPostOnly        bool    `json:"reject_post_only,omitempty"`
	Label                 string  `json:"label"`
	OrderState            string  `json:"order_state"`
	Usd                   float64 `json:"usd,omitempty"`
	Implv                 float64 `json:"implv,omitempty"`
	TriggerReferencePrice float64 `json:"trigger_reference_price"`
	OriginalOrderType     string  `json:"original_order_type"`
	BlockTrade            bool    `json:"block_trade,omitempty"`
	TriggerPrice          float64 `json:"trigger_price,omitempty"`
	API                   bool    `json:"api"`
	MMP                   bool    `json:"mmp"`
	TriggerOrderId        string  `json:"trigger_order_id"`
	CancelReason          string  `json:"cancel_reason"`
	RiskReducing          bool    `json:"risk_reducing"`
	FilledAmount          float64 `json:"filled_amount"`
	InstrumentName        string  `json:"instrument_name"`
	MaxShow               float64 `json:"max_show"`
	AppName               string  `json:"app_name"`
	MMPCancelled          bool    `json:"mmp_cancelled"`
	Direction             string  `json:"direction"`
	LastUpdateTimestamp   int64   `json:"last_update_timestamp"`
	TriggerOffset         float64 `json:"trigger_offset"`
	Price                 Price   `json:"price"`
	IsLiquidation         bool    `json:"is_liquidation,omitempty"`
	ReduceOnly            bool    `json:"reduce_only,omitempty"`
	Amount                float64 `json:"amount"`
	PostOnly              bool    `json:"post_only,omitempty"`
	Mobile                bool    `json:"mobile"`
	Triggered             bool    `json:"triggered"`
	OrderId               string  `json:"order_id"`
	Replaced              bool    `json:"replaced"`
	OrderType             string  `json:"order_type"`
	TimeInForce           string  `json:"time_in_force"`
	AutoReplaced          bool    `json:"auto_replaced"`
	Trigger               string  `json:"trigger"`
	Web                   bool    `json:"web"`
	CreationTimestamp     int64   `json:"creation_timestamp"`
	AveragePrice          float64 `json:"average_price"`
	Advanced              string  `json:"advanced"`
}

type Price float64

func (p *Price) UnmarshalJSON(data []byte) error {
	if string(data) == `"market_price"` {
		*p = 0
		return nil
	}
	var f float64
	f, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return err
	}

	*p = Price(f)

	return nil
}

func (p *Price) ToFloat64() float64 {
	return float64(*p)
}
