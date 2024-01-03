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

package account

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/europeanoptions/account/types"
	eoutils "github.com/linstohu/nexapi/binance/europeanoptions/utils"
	usdmutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
	"github.com/linstohu/nexapi/utils"
)

type OptionsAccountClient struct {
	*eoutils.OptionsClient

	// validate struct fields
	validate *validator.Validate
}

type OptionsAccountClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewOptionsAccountClient(cfg *eoutils.OptionsClientCfg) (*OptionsAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := eoutils.NewOptionsClient(&eoutils.OptionsClientCfg{
		Debug:      cfg.Debug,
		Logger:     cfg.Logger,
		BaseURL:    cfg.BaseURL,
		Key:        cfg.Key,
		Secret:     cfg.Secret,
		RecvWindow: cfg.RecvWindow,
	})
	if err != nil {
		return nil, err
	}

	return &OptionsAccountClient{
		OptionsClient: cli,
		validate:      validator,
	}, nil
}

func (o *OptionsAccountClient) GetAccountInfo(ctx context.Context) (*types.GetAccountInfoResp, error) {
	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/account",
		Method:  http.MethodGet,
	}

	securityType := usdmutils.TRADE

	{
		headers, err := o.GenGetHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := bnutils.DefaultParam{
			RecvWindow: o.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
		}

		err := o.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := o.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.AccountInfo
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetAccountInfoResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (o *OptionsAccountClient) NewOrder(ctx context.Context, param types.NewOrderParam) (*types.OrderResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/order",
		Method:  http.MethodPost,
	}

	securityType := usdmutils.TRADE

	{
		headers, err := o.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.NewOrderParams{
			NewOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := o.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.Order
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.OrderResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (o *OptionsAccountClient) GetSingleOrder(ctx context.Context, param types.GetSingleOrderParam) (*types.OrderResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	if param.OrderID == 0 && param.ClientOrderID == "" {
		return nil, fmt.Errorf("either orderId or clientOrderId must be sent")
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/order",
		Method:  http.MethodGet,
	}

	securityType := usdmutils.TRADE

	{
		headers, err := o.GenGetHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetSingleOrderParams{
			GetSingleOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := o.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.Order
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.OrderResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (o *OptionsAccountClient) CancelOrder(ctx context.Context, param types.CancelOrderParam) (*types.OrderResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	if param.OrderID == 0 && param.ClientOrderID == "" {
		return nil, fmt.Errorf("either orderId or clientOrderId must be sent")
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/order",
		Method:  http.MethodDelete,
	}

	securityType := usdmutils.TRADE

	{
		headers, err := o.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.CancelOrderParams{
			CancelOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := o.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.Order
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.OrderResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (o *OptionsAccountClient) CancelAllOrdersBySymbol(ctx context.Context, param types.CancelAllOrdersParam) error {
	err := o.validate.Struct(param)
	if err != nil {
		return err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/allOpenOrders",
		Method:  http.MethodDelete,
	}

	securityType := usdmutils.TRADE

	{
		headers, err := o.GenHeaders(securityType)
		if err != nil {
			return err
		}
		req.Headers = headers
	}

	{
		query := types.CancelAllOrdersParams{
			CancelAllOrdersParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(query)
		if err != nil {
			return err
		}

		if need := o.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	_, err = o.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (o *OptionsAccountClient) CancelAllOrdersByUnderlying(ctx context.Context, param types.CancelAllOrdersByUnderlyingParam) error {
	err := o.validate.Struct(param)
	if err != nil {
		return err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/allOpenOrdersByUnderlying",
		Method:  http.MethodDelete,
	}

	securityType := usdmutils.TRADE

	{
		headers, err := o.GenHeaders(securityType)
		if err != nil {
			return err
		}
		req.Headers = headers
	}

	{
		query := types.CancelAllOrdersByUnderlyingParams{
			CancelAllOrdersByUnderlyingParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(query)
		if err != nil {
			return err
		}

		if need := o.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	_, err = o.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (o *OptionsAccountClient) GetOpenOrders(ctx context.Context, param types.GetCurrentOpenOrdersParam) (*types.OrdersResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/openOrders",
		Method:  http.MethodGet,
	}

	securityType := usdmutils.TRADE

	{
		headers, err := o.GenGetHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetCurrentOpenOrdersParams{
			GetCurrentOpenOrdersParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := o.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.Order
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.OrdersResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (o *OptionsAccountClient) GetOrderHistory(ctx context.Context, param types.GetOrderHistoryParam) (*types.OrdersResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/historyOrders",
		Method:  http.MethodGet,
	}

	securityType := usdmutils.TRADE

	{
		headers, err := o.GenGetHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetOrderHistoryParams{
			GetOrderHistoryParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := o.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.Order
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.OrdersResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (o *OptionsAccountClient) GetPositionInfo(ctx context.Context, param types.GetPositionInfoParam) (*types.GetPositionInfoResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/position",
		Method:  http.MethodGet,
	}

	securityType := usdmutils.TRADE

	{
		headers, err := o.GenGetHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetPositionInfoParams{
			GetPositionInfoParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := o.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.Position
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetPositionInfoResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (o *OptionsAccountClient) GetTradeList(ctx context.Context, param types.GetTradeListParam) (*types.GetTradeListResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/userTrades",
		Method:  http.MethodGet,
	}

	securityType := usdmutils.TRADE

	{
		headers, err := o.GenGetHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetTradeListParams{
			GetTradeListParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := o.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.UserTrade
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetTradeListResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (o *OptionsAccountClient) GetExerciseRecord(ctx context.Context, param types.GetExerciseRecordParam) (*types.GetTradeListResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/exerciseRecord",
		Method:  http.MethodGet,
	}

	securityType := usdmutils.TRADE

	{
		headers, err := o.GenGetHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetExerciseRecordParams{
			GetExerciseRecordParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := o.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.UserTrade
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetTradeListResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (o *OptionsAccountClient) GetFundingFlow(ctx context.Context, param types.GetFundingFlowParam) (*types.GetFundingFlowResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/bill",
		Method:  http.MethodGet,
	}

	securityType := usdmutils.TRADE

	{
		headers, err := o.GenGetHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetFundingFlowParams{
			GetFundingFlowParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := o.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.FundingFlow
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetFundingFlowResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}
