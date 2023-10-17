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

type Order struct {
	Side                 string  `json:"side"`
	Status               string  `json:"status"`
	Symbol               string  `json:"symbol"`
	ClientOrderID        int     `json:"client_order_id"`
	ReduceOnly           bool    `json:"reduce_only"`
	OrderID              int     `json:"order_id"`
	OrderTag             string  `json:"order_tag"`
	Type                 string  `json:"type"`
	Price                float64 `json:"price"`
	Quantity             float64 `json:"quantity"`
	Amount               float64 `json:"amount"`
	Visible              float64 `json:"visible"`
	Executed             float64 `json:"executed"`
	TotalFee             float64 `json:"total_fee"`
	FeeAsset             string  `json:"fee_asset"`
	CreatedTime          string  `json:"created_time"`
	UpdatedTime          string  `json:"updated_time"`
	AverageExecutedPrice float64 `json:"average_executed_price"`
	RealizedPnl          float64 `json:"realized_pnl"`
}

type GetOrder struct {
	Response
	Order
}

type GetOrdersParam struct {
	Symbol    string `url:"symbol,omitempty"`
	Side      string `url:"side,omitempty" validate:"omitempty,oneof=SELL BUY"`
	Size      int64  `url:"size,omitempty" validate:"omitempty,max=500"`
	OrderType string `url:"order_type,omitempty" validate:"omitempty,oneof=LIMIT MARKET IOC FOK POST_ONLY LIQUIDATE"`
	OrderTag  string `url:"order_tag,omitempty" validate:"omitempty"`
	Status    string `url:"status,omitempty" validate:"omitempty,oneof=NEW CANCELLED PARTIAL_FILLED FILLED REJECTED INCOMPLETE COMPLETED"`
	StartTime int64  `url:"start_t,omitempty" validate:"omitempty,gt=999999999999"`
	EndTime   int64  `url:"end_t,omitempty" validate:"omitempty,gt=999999999999"`
	Page      int64  `url:"page,omitempty" validate:"omitempty"`
}

type GetOrders struct {
	Response
	Meta struct {
		Total          int `json:"total"`
		RecordsPerPage int `json:"records_per_page"`
		CurrentPage    int `json:"current_page"`
	} `json:"meta"`
	Rows []Order `json:"rows"`
}
