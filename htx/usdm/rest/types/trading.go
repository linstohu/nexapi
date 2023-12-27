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

type PlaceIsolatedOrderParam struct {
	ContractCode     string  `json:"contract_code" validate:"required"`
	ReduceOnly       int     `json:"reduce_only,omitempty" validate:"omitempty"`
	ClientOrderID    int64   `json:"client_order_id,omitempty" validate:"omitempty"`
	Price            float64 `json:"price,omitempty" validate:"omitempty"`
	Volume           int64   `json:"volume" validate:"required"`
	Direction        string  `json:"direction" validate:"required"`
	Offset           string  `json:"offset,omitempty" validate:"omitempty"`
	LeverRate        int     `json:"lever_rate" validate:"required"`
	OrderPriceType   string  `json:"order_price_type" validate:"required"`
	TpTriggerPrice   float64 `json:"tp_trigger_price,omitempty" validate:"omitempty"`
	TpOrderPrice     float64 `json:"tp_order_price,omitempty" validate:"omitempty"`
	TpOrderPriceType string  `json:"tp_order_price_type,omitempty" validate:"omitempty"`
	SlTriggerPrice   float64 `json:"sl_trigger_price,omitempty" validate:"omitempty"`
	SlOrderPrice     float64 `json:"sl_order_price,omitempty" validate:"omitempty"`
	SlOrderPriceType string  `json:"sl_order_price_type,omitempty" validate:"omitempty"`
}

type PlaceOrderResp struct {
	DefaultResponse
	Data struct {
		OrderID       int64  `json:"order_id,omitempty"`
		ClientOrderID int64  `json:"client_order_id,omitempty"`
		OrderIDStr    string `json:"order_id_str,omitempty"`
	} `json:"data"`
}

type PlaceCrossOrderParam struct {
	ContractCode     string  `json:"contract_code" validate:"required"`
	Pair             string  `json:"pair,omitempty" validate:"omitempty"`
	ContractType     string  `json:"contract_type,omitempty" validate:"omitempty"`
	ReduceOnly       int     `json:"reduce_only,omitempty" validate:"omitempty"`
	ClientOrderID    int64   `json:"client_order_id,omitempty" validate:"omitempty"`
	Price            float64 `json:"price,omitempty" validate:"omitempty"`
	Volume           int64   `json:"volume" validate:"required"`
	Direction        string  `json:"direction" validate:"required"`
	Offset           string  `json:"offset,omitempty" validate:"omitempty"`
	LeverRate        int     `json:"lever_rate" validate:"required"`
	OrderPriceType   string  `json:"order_price_type" validate:"required"`
	TpTriggerPrice   float64 `json:"tp_trigger_price,omitempty" validate:"omitempty"`
	TpOrderPrice     float64 `json:"tp_order_price,omitempty" validate:"omitempty"`
	TpOrderPriceType string  `json:"tp_order_price_type,omitempty" validate:"omitempty"`
	SlTriggerPrice   float64 `json:"sl_trigger_price,omitempty" validate:"omitempty"`
	SlOrderPrice     float64 `json:"sl_order_price,omitempty" validate:"omitempty"`
	SlOrderPriceType string  `json:"sl_order_price_type,omitempty" validate:"omitempty"`
}

type CancelIsolatedOrderParam struct {
	OrderID       string `json:"order_id,omitempty" validate:"omitempty"`
	ClientOrderID string `json:"client_order_id,omitempty" validate:"omitempty"`
	ContractCode  string `json:"contract_code" validate:"required"`
}

type CancelOrderResp struct {
	DefaultResponse
	Data struct {
		Errors []struct {
			OrderID string `json:"order_id,omitempty"`
			ErrCode int    `json:"err_code,omitempty"`
			ErrMsg  string `json:"err_msg,omitempty"`
		} `json:"errors,omitempty"`
		Successes string `json:"successes,omitempty"`
	} `json:"data"`
}

type CancelCrossOrderParam struct {
	OrderID       string `json:"order_id,omitempty" validate:"omitempty"`
	ClientOrderID string `json:"client_order_id,omitempty" validate:"omitempty"`
	ContractCode  string `json:"contract_code,omitempty" validate:"omitempty"`
	Pair          string `json:"pair,omitempty" validate:"omitempty"`
	ContractType  string `json:"contract_type,omitempty" validate:"omitempty"`
}

type CancelAllIsolatedOrdersParam struct {
	ContractCode string `json:"contract_code" validate:"required"`
	Direction    string `json:"direction,omitempty" validate:"omitempty"`
	Offset       string `json:"offset,omitempty" validate:"omitempty"`
}

type CancelAllCrossOrdersParam struct {
	ContractCode string `json:"contract_code,omitempty" validate:"omitempty"`
	Pair         string `json:"pair,omitempty" validate:"omitempty"`
	ContractType string `json:"contract_type,omitempty" validate:"omitempty"`
	Direction    string `json:"direction,omitempty" validate:"omitempty"`
	Offset       string `json:"offset,omitempty" validate:"omitempty"`
}

type GetIsolatedOpenOrdersParam struct {
	ContractCode string `json:"contract_code" validate:"required"`
	PageIndex    int    `json:"page_index,omitempty" validate:"omitempty"`
	PageSize     int    `json:"page_size,omitempty" validate:"omitempty"`
	SortBy       string `json:"sort_by,omitempty" validate:"omitempty"`
	TradeType    int    `json:"trade_type,omitempty" validate:"omitempty"`
}

type GetIsolatedOpenOrdersResp struct {
	DefaultResponse
	Data IsolatedOpenOrders `json:"data"`
}

type IsolatedOrder struct {
	Symbol          string  `json:"symbol,omitempty"`
	ContractCode    string  `json:"contract_code,omitempty"`
	Volume          float64 `json:"volume,omitempty"`
	Price           float64 `json:"price,omitempty"`
	OrderPriceType  string  `json:"order_price_type,omitempty"`
	OrderType       int     `json:"order_type,omitempty"`
	Direction       string  `json:"direction,omitempty"`
	Offset          string  `json:"offset,omitempty"`
	LeverRate       int     `json:"lever_rate,omitempty"`
	OrderID         int64   `json:"order_id,omitempty"`
	ClientOrderID   int64   `json:"client_order_id,omitempty"`
	CreatedAt       int64   `json:"created_at,omitempty"`
	TradeVolume     float64 `json:"trade_volume,omitempty"`
	TradeTurnover   float64 `json:"trade_turnover,omitempty"`
	Fee             float64 `json:"fee,omitempty"`
	TradeAvgPrice   float64 `json:"trade_avg_price,omitempty"`
	MarginFrozen    float64 `json:"margin_frozen,omitempty"`
	Profit          float64 `json:"profit,omitempty"`
	Status          int     `json:"status,omitempty"`
	OrderSource     string  `json:"order_source,omitempty"`
	OrderIDStr      string  `json:"order_id_str,omitempty"`
	FeeAsset        string  `json:"fee_asset,omitempty"`
	LiquidationType string  `json:"liquidation_type,omitempty"`
	CanceledAt      int64   `json:"canceled_at,omitempty"`
	MarginAsset     string  `json:"margin_asset,omitempty"`
	MarginMode      string  `json:"margin_mode,omitempty"`
	MarginAccount   string  `json:"margin_account,omitempty"`
	IsTpsl          int     `json:"is_tpsl,omitempty"`
	UpdateTime      int64   `json:"update_time,omitempty"`
	RealProfit      float64 `json:"real_profit,omitempty"`
	ReduceOnly      int     `json:"reduce_only,omitempty"`
}

type IsolatedOpenOrders struct {
	Orders      []IsolatedOrder `json:"orders,omitempty"`
	TotalPage   int             `json:"total_page,omitempty"`
	CurrentPage int             `json:"current_page,omitempty"`
	TotalSize   int             `json:"total_size,omitempty"`
}

type GetCrossOpenOrdersParam struct {
	ContractCode string `json:"contract_code,omitempty" validate:"omitempty"`
	Pair         string `json:"pair,omitempty" validate:"omitempty"`
	PageIndex    int    `json:"page_index,omitempty" validate:"omitempty"`
	PageSize     int    `json:"page_size,omitempty" validate:"omitempty"`
	SortBy       string `json:"sort_by,omitempty" validate:"omitempty"`
	TradeType    int    `json:"trade_type,omitempty" validate:"omitempty"`
}

type GetCrossOpenOrdersResp struct {
	DefaultResponse
	Data CrossOpenOrders `json:"data"`
}

type CrossOpenOrders struct {
	Orders      []CrossOpenOrder `json:"orders,omitempty"`
	TotalPage   int              `json:"total_page,omitempty"`
	CurrentPage int              `json:"current_page,omitempty"`
	TotalSize   int              `json:"total_size,omitempty"`
}

type CrossOpenOrder struct {
	UpdateTime      int64   `json:"update_time,omitempty"`
	BusinessType    string  `json:"business_type,omitempty"`
	ContractType    string  `json:"contract_type,omitempty"`
	Pair            string  `json:"pair,omitempty"`
	Symbol          string  `json:"symbol,omitempty"`
	ContractCode    string  `json:"contract_code,omitempty"`
	Volume          float64 `json:"volume,omitempty"`
	Price           float64 `json:"price,omitempty"`
	OrderPriceType  string  `json:"order_price_type,omitempty"`
	OrderType       int     `json:"order_type,omitempty"`
	Direction       string  `json:"direction,omitempty"`
	Offset          string  `json:"offset,omitempty"`
	LeverRate       int     `json:"lever_rate,omitempty"`
	OrderID         int64   `json:"order_id,omitempty"`
	ClientOrderID   int64   `json:"client_order_id,omitempty"`
	CreatedAt       int64   `json:"created_at,omitempty"`
	TradeVolume     float64 `json:"trade_volume,omitempty"`
	TradeTurnover   float64 `json:"trade_turnover,omitempty"`
	Fee             float64 `json:"fee,omitempty"`
	TradeAvgPrice   float64 `json:"trade_avg_price,omitempty"`
	MarginFrozen    float64 `json:"margin_frozen,omitempty"`
	Profit          float64 `json:"profit,omitempty"`
	Status          int     `json:"status,omitempty"`
	OrderSource     string  `json:"order_source,omitempty"`
	OrderIDStr      string  `json:"order_id_str,omitempty"`
	FeeAsset        string  `json:"fee_asset,omitempty"`
	LiquidationType string  `json:"liquidation_type,omitempty"`
	CanceledAt      int64   `json:"canceled_at,omitempty"`
	MarginAsset     string  `json:"margin_asset,omitempty"`
	MarginAccount   string  `json:"margin_account,omitempty"`
	MarginMode      string  `json:"margin_mode,omitempty"`
	IsTpsl          int     `json:"is_tpsl,omitempty"`
	RealProfit      string  `json:"real_profit,omitempty"`
	ReduceOnly      int     `json:"reduce_only,omitempty"`
}

type GetIsolatedHistoryMatchResultsParam struct {
	Contract  string `json:"contract,omitempty" validate:"omitempty"`
	Pair      string `json:"pair,omitempty" validate:"omitempty"`
	TradeType int    `json:"trade_type" validate:"required"`
	StartTime int64  `json:"start_time,omitempty"`
	EndTime   int64  `json:"end_time,omitempty"`
	Direct    string `json:"direct,omitempty" validate:"omitempty"`
	FromID    int64  `json:"from_id,omitempty"`
}

type GetCrossHistoryMatchResultsParam struct {
	Contract  string `json:"contract" validate:"required"`
	Pair      string `json:"pair,omitempty" validate:"omitempty"`
	TradeType int    `json:"trade_type" validate:"required"`
	StartTime int64  `json:"start_time,omitempty"`
	EndTime   int64  `json:"end_time,omitempty"`
	Direct    string `json:"direct,omitempty" validate:"omitempty"`
	FromID    int64  `json:"from_id,omitempty"`
}

type HistoryMatchResultsResp struct {
	DefaultResponse
	Data []HistoryMatchResult `json:"data"`
}

type HistoryMatchResult struct {
	ID               string  `json:"id,omitempty"`
	QueryID          int64   `json:"query_id,omitempty"`
	MatchID          int64   `json:"match_id,omitempty"`
	OrderID          int64   `json:"order_id,omitempty"`
	OrderIDStr       string  `json:"order_id_str,omitempty"`
	Symbol           string  `json:"symbol,omitempty"`
	ContractCode     string  `json:"contract_code,omitempty"`
	MarginMode       string  `json:"margin_mode,omitempty"`
	MarginAccount    string  `json:"margin_account,omitempty"`
	Direction        string  `json:"direction,omitempty"`
	Offset           string  `json:"offset,omitempty"`
	TradeVolume      float64 `json:"trade_volume,omitempty"`
	TradePrice       float64 `json:"trade_price,omitempty"`
	TradeTurnover    float64 `json:"trade_turnover,omitempty"`
	CreateDate       int64   `json:"create_date,omitempty"`
	OffsetProfitloss float64 `json:"offset_profitloss,omitempty"`
	RealProfit       float64 `json:"real_profit,omitempty"`
	TradeFee         float64 `json:"trade_fee,omitempty"`
	Role             string  `json:"role,omitempty"`
	FeeAsset         string  `json:"fee_asset,omitempty"`
	OrderSource      string  `json:"order_source,omitempty"`
	ContractType     string  `json:"contract_type,omitempty"`
	Pair             string  `json:"pair,omitempty"`
	BusinessType     string  `json:"business_type,omitempty"`
	ReduceOnly       int     `json:"reduce_only,omitempty"`
}
