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
	"log/slog"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/coinmfutures/account/types"
	cmutils "github.com/linstohu/nexapi/binance/coinmfutures/utils"
	umutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
	"github.com/linstohu/nexapi/utils"
)

type CoinMFuturesAccountClient struct {
	*cmutils.CoinMarginedClient

	// validate struct fields
	validate *validator.Validate
}

type CoinMFuturesAccountClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewCoinMFuturesAccountClient(cfg *cmutils.CoinMarginedClientCfg) (*CoinMFuturesAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := cmutils.NewCoinMarginedClient(&cmutils.CoinMarginedClientCfg{
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

	return &CoinMFuturesAccountClient{
		CoinMarginedClient: cli,
		validate:           validator,
	}, nil
}

func (c *CoinMFuturesAccountClient) ChangePositionMode(ctx context.Context, param types.ChangePositionModeParam) (*types.DefaultResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/positionSide/dual",
		Method:  http.MethodPost,
	}

	securityType := umutils.TRADE

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.ChangePositionModeParams{
			ChangePositionModeParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: c.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := c.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.Response
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.DefaultResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (c *CoinMFuturesAccountClient) GetPositionMode(ctx context.Context) (*types.GetCurrentPositionModeResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/positionSide/dual",
		Method:  http.MethodGet,
	}

	securityType := umutils.TRADE

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := bnutils.DefaultParam{
			RecvWindow: c.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
		}

		err := c.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.GetCurrentPositionModeAPIResp
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetCurrentPositionModeResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (c *CoinMFuturesAccountClient) NewOrder(ctx context.Context, param types.NewOrderParam) (*types.OrderResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/order",
		Method:  http.MethodPost,
	}

	securityType := umutils.TRADE

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.NewOrderParams{
			NewOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: c.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := c.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := c.SendHTTPRequest(ctx, req)
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

func (c *CoinMFuturesAccountClient) QueryOrder(ctx context.Context, param types.GetOrderParam) (*types.OrderResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/order",
		Method:  http.MethodGet,
	}

	securityType := umutils.USER_DATA

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetOrderParams{
			GetOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: c.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := c.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := c.SendHTTPRequest(ctx, req)
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

func (c *CoinMFuturesAccountClient) QueryOpenOrder(ctx context.Context, param types.GetOrderParam) (*types.OrderResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/openOrder",
		Method:  http.MethodGet,
	}

	securityType := umutils.USER_DATA

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetOrderParams{
			GetOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: c.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := c.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := c.SendHTTPRequest(ctx, req)
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

func (c *CoinMFuturesAccountClient) QueryAllOpenOrders(ctx context.Context, param types.GetAllOpenOrdersParam) (*types.OrdersResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/openOrders",
		Method:  http.MethodGet,
	}

	securityType := umutils.USER_DATA

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetAllOpenOrdersParams{
			GetAllOpenOrdersParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: c.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := c.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := c.SendHTTPRequest(ctx, req)
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

func (c *CoinMFuturesAccountClient) CancelOrder(ctx context.Context, param types.GetOrderParam) (*types.OrderResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/order",
		Method:  http.MethodDelete,
	}

	securityType := umutils.TRADE

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.GetOrderParams{
			GetOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: c.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := c.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := c.SendHTTPRequest(ctx, req)
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

func (c *CoinMFuturesAccountClient) CancelAllOpenOrders(ctx context.Context, param types.CancelAllOpenOrdersParam) error {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/allOpenOrders",
		Method:  http.MethodDelete,
	}

	securityType := umutils.TRADE

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return err
		}
		req.Headers = headers
	}

	{
		body := types.CancelAllOpenOrdersParams{
			CancelAllOpenOrdersParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: c.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := c.validate.Struct(body)
		if err != nil {
			return err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	_, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (c *CoinMFuturesAccountClient) GetAllOrders(ctx context.Context, param types.GetAllOrdersParam) (*types.OrdersResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/allOrders",
		Method:  http.MethodGet,
	}

	securityType := umutils.USER_DATA

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetAllOrdersParams{
			GetAllOrdersParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: c.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := c.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := c.SendHTTPRequest(ctx, req)
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

func (c *CoinMFuturesAccountClient) GetBalance(ctx context.Context) (*types.GetBalanceResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/balance",
		Method:  http.MethodGet,
	}

	securityType := umutils.USER_DATA

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := bnutils.DefaultParam{
			RecvWindow: c.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
		}

		err := c.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.Balance
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetBalanceResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (c *CoinMFuturesAccountClient) GetAccountInformation(ctx context.Context) (*types.GetAccountInfoResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/account",
		Method:  http.MethodGet,
	}

	securityType := umutils.USER_DATA

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := bnutils.DefaultParam{
			RecvWindow: c.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
		}

		err := c.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.Account
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetAccountInfoResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (c *CoinMFuturesAccountClient) ChangeInitialLeverage(ctx context.Context, param types.ChangeLeverageParam) (*types.ChangeLeverageResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/leverage",
		Method:  http.MethodPost,
	}

	securityType := umutils.TRADE

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.ChangeLeverageParams{
			ChangeLeverageParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: c.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := c.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.ChangeLeverageAPIResp
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.ChangeLeverageResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (c *CoinMFuturesAccountClient) ChangeMarginType(ctx context.Context, param types.ChangeMarginTypeParam) error {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/marginType",
		Method:  http.MethodPost,
	}

	securityType := umutils.TRADE

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return err
		}
		req.Headers = headers
	}

	{
		body := types.ChangeMarginTypeParams{
			ChangeMarginTypeParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: c.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := c.validate.Struct(body)
		if err != nil {
			return err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	_, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (c *CoinMFuturesAccountClient) ModifyIsolatedPositionMargin(ctx context.Context, param types.ModifyIsolatedPositionMarginParam) (*types.ModifyIsolatedPositionMarginResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/positionMargin",
		Method:  http.MethodPost,
	}

	securityType := umutils.TRADE

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.ModifyIsolatedPositionMarginParams{
			ModifyIsolatedPositionMarginParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: c.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := c.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.ModifyIsolatedPositionMarginAPIResp
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.ModifyIsolatedPositionMarginResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (c *CoinMFuturesAccountClient) GetPositionInformation(ctx context.Context, param types.GetPositionParam) (*types.GetPositionResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/positionRisk",
		Method:  http.MethodGet,
	}

	securityType := umutils.USER_DATA

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetPositionParams{
			GetPositionParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: c.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := c.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.Position
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetPositionResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (c *CoinMFuturesAccountClient) GetAccountTradeList(ctx context.Context, param types.GetTradeListParam) (*types.GetTradeListResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/userTrades",
		Method:  http.MethodGet,
	}

	securityType := umutils.USER_DATA

	{
		headers, err := c.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetTradeListParams{
			GetTradeListParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: c.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := c.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := c.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(c.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.Trade
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetTradeListResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}
