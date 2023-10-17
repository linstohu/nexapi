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

type Trade struct {
	ID                int     `json:"id"`
	Symbol            string  `json:"symbol"`
	OrderID           int     `json:"order_id"`
	OrderTag          string  `json:"order_tag"`
	ExecutedPrice     float64 `json:"executed_price"`
	ExecutedQuantity  float64 `json:"executed_quantity"`
	IsMaker           int     `json:"is_maker"`
	Side              string  `json:"side"`
	Fee               float64 `json:"fee"`
	FeeAsset          string  `json:"fee_asset"`
	ExecutedTimestamp string  `json:"executed_timestamp"`
}

type GetTrade struct {
	Response
	Trade
}

type GetTrades struct {
	Response
	Rows []Trade `json:"rows"`
}

type GetTradeHistoryParam struct {
	Symbol    string `url:"symbol,omitempty"`
	OrderTag  string `url:"order_tag,omitempty" validate:"omitempty"`
	StartTime int64  `url:"start_t,omitempty" validate:"omitempty,gt=999999999999"`
	EndTime   int64  `url:"end_t,omitempty" validate:"omitempty,gt=999999999999"`
	Page      int64  `url:"page,omitempty" validate:"omitempty"`
	Size      int64  `url:"size,omitempty" validate:"omitempty"`
}

type GetTradeHistory struct {
	Response
	Meta struct {
		Total          int `json:"total"`
		RecordsPerPage int `json:"records_per_page"`
		CurrentPage    int `json:"current_page"`
	} `json:"meta"`
	Rows []Trade `json:"rows"`
}
