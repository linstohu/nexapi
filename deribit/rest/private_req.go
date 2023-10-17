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

package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/linstohu/nexapi/deribit/rest/types"
	"github.com/linstohu/nexapi/deribit/rest/types/account"
	"github.com/linstohu/nexapi/deribit/rest/types/auth"
	"github.com/linstohu/nexapi/deribit/rest/types/trading"
)

func (d *DeribitRestClient) Auth(ctx context.Context, param auth.AuthParams) (*auth.AuthResponse, error) {
	err := d.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/public/auth",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "public/auth",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret auth.AuthResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

var ErrAuth = errors.New("auth error, you should reinitialize client using key and secret")

func (d *DeribitRestClient) checkAuth() error {
	if d.auth.token == "" {
		return ErrAuth
	}

	now := time.Now().Unix()
	if now > d.auth.expiresAt {
		token, err := d.Auth(context.TODO(), auth.AuthParams{
			GrantType:    "client_credentials",
			ClientID:     d.key,
			ClientSecret: d.secret,
		})
		if err != nil {
			return fmt.Errorf("auth failed, error: %v", err)
		}

		now := time.Now().Unix()
		d.auth.token = token.AccessToken
		d.auth.expiresAt = now + token.ExpiresIn - 5
	}

	return nil
}

func (d *DeribitRestClient) genAuthHeaders() map[string]string {
	ret := DefaultContentType
	ret["Authorization"] = fmt.Sprintf("Bearer %s", d.auth.token)
	return ret
}

func (d *DeribitRestClient) GetAccountSummary(ctx context.Context, param account.GetAccountSummaryParams) (*account.AccountSummary, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/get_account_summary",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/get_account_summary",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret account.AccountSummary
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetPositions(ctx context.Context, param account.GetPositionsParams) ([]*account.Position, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/get_positions",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/get_positions",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret []*account.Position
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (d *DeribitRestClient) GetTransactionLog(ctx context.Context, param account.GetTransactionLogParams) (*account.GetTransactionLogResponse, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/get_transaction_log",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/get_transaction_log",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret account.GetTransactionLogResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) Buy(ctx context.Context, param trading.BuyParams) (*trading.BuyResponse, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/buy",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/buy",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret trading.BuyResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) Sell(ctx context.Context, param trading.SellParams) (*trading.SellResponse, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/sell",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/sell",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret trading.SellResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) Cancel(ctx context.Context, param trading.CancelParams) (*trading.Order, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/cancel",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/cancel",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret trading.Order
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) CancelAll(ctx context.Context) (int, error) {
	if err := d.checkAuth(); err != nil {
		return 0, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/cancel_all",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/cancel_all",
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return 0, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return 0, err
	}

	var ret int
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return 0, err
	}

	return ret, nil
}

func (d *DeribitRestClient) CancelAllByInstrument(ctx context.Context, param trading.CancelAllByInstrumentParams) (int, error) {
	if err := d.checkAuth(); err != nil {
		return 0, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/cancel_all_by_instrument",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/cancel_all_by_instrument",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return 0, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return 0, err
	}

	var ret int
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return 0, err
	}

	return ret, nil
}

func (d *DeribitRestClient) ClosePosition(ctx context.Context, param trading.ClosePositionParams) (*trading.ClosePositionResponse, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/close_position",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/close_position",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret trading.ClosePositionResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetOpenOrdersByCurrency(ctx context.Context, param trading.GetOpenOrdersByCurrencyParams) ([]*trading.Order, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/get_open_orders_by_currency",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/get_open_orders_by_currency",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret []*trading.Order
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (d *DeribitRestClient) GetOpenOrdersByInstrument(ctx context.Context, param trading.GetOpenOrdersByInstrumentParams) ([]*trading.Order, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/get_open_orders_by_instrument",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/get_open_orders_by_instrument",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret []*trading.Order
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (d *DeribitRestClient) GetOrderState(ctx context.Context, param trading.GetOrderStateParams) (*trading.Order, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/get_order_state",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/get_order_state",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret trading.Order
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetUserTradesByCurrency(ctx context.Context, param trading.GetUserTradesByCurrencyParams) (*trading.GetUserTradesResponse, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/get_user_trades_by_currency",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/get_user_trades_by_currency",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret trading.GetUserTradesResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetUserTradesByCurrencyAndTime(ctx context.Context, param trading.GetUserTradesByCurrencyAndTimeParams) (*trading.GetUserTradesResponse, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/get_user_trades_by_currency_and_time",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/get_user_trades_by_currency_and_time",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret trading.GetUserTradesResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetUserTradesByInstrument(ctx context.Context, param trading.GetUserTradesByInstrumentParams) (*trading.GetUserTradesResponse, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/get_user_trades_by_instrument",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/get_user_trades_by_instrument",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret trading.GetUserTradesResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetUserTradesByInstrumentAndTime(ctx context.Context, param trading.GetUserTradesByInstrumentAndTimeParams) (*trading.GetUserTradesResponse, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/get_user_trades_by_instrument_and_time",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/get_user_trades_by_instrument_and_time",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret trading.GetUserTradesResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetSettlementHistoryByInstrument(ctx context.Context, param trading.GetSettlementHistoryByInstrumentParams) (*trading.GetSettlementHistoryResponse, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/get_settlement_history_by_instrument",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/get_settlement_history_by_instrument",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret trading.GetSettlementHistoryResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetSettlementHistoryByCurrency(ctx context.Context, param trading.GetSettlementHistoryByCurrencyParams) (*trading.GetSettlementHistoryResponse, error) {
	if err := d.checkAuth(); err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/get_settlement_history_by_currency",
		Method:  http.MethodPost,
		Headers: d.genAuthHeaders(),
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/get_settlement_history_by_currency",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret trading.GetSettlementHistoryResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
