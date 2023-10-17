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

type SendOrderReq struct {
	Symbol          string  `url:"symbol" validate:"required"`
	ClientOrderID   int64   `url:"client_order_id,omitempty"`
	OrderTag        string  `url:"order_tag,omitempty"`
	OrderType       string  `url:"order_type" validate:"required,oneof=LIMIT MARKET IOC FOK POST_ONLY ASK BID"`
	OrderPrice      float64 `url:"order_price,omitempty"`
	OrderQuantity   float64 `url:"order_quantity,omitempty"`
	OrderAmount     float64 `url:"order_amount,omitempty"`
	ReduceOnly      bool    `url:"reduce_only,omitempty"`
	VisibleQuantity float64 `url:"visible_quantity,omitempty"`
	Side            string  `url:"side" validate:"required,oneof=SELL BUY"`
}

type SendOrderResp struct {
	Response
	OrderID       int     `json:"order_id"`
	ClientOrderID int     `json:"client_order_id"`
	OrderType     string  `json:"order_type"`
	OrderPrice    float64 `json:"order_price"`
	OrderQuantity float64 `json:"order_quantity"`
	OrderAmount   float64 `json:"order_amount"`
	ReduceOnly    bool    `json:"reduce_only"`
	Timestamp     string  `json:"timestamp"`
}
