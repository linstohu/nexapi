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

package spotaccount

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/spot/spotaccount/types"
	spotutils "github.com/linstohu/nexapi/binance/spot/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
	"github.com/linstohu/nexapi/utils"
)

type SpotAccountClient struct {
	*spotutils.SpotClient

	// validate struct fields
	validate *validator.Validate
}

type SpotAccountClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewSpotAccountClient(cfg *SpotAccountClientCfg) (*SpotAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := spotutils.NewSpotClient(&spotutils.SpotClientCfg{
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

	return &SpotAccountClient{
		SpotClient: cli,
		validate:   validator,
	}, nil
}

func (s *SpotAccountClient) TestNewOrder(ctx context.Context, param types.NewOrderParam) error {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/order/test",
		Method:  http.MethodPost,
	}

	st := spotutils.TRADE

	{
		headers, err := s.GenHeaders(st)
		if err != nil {
			return err
		}
		req.Headers = headers
	}

	{
		body := types.NewOrderParams{
			NewOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: s.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := s.validate.Struct(body)
		if err != nil {
			return err
		}

		if need := s.NeedSignature(st); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return err
			}

			h := hmac.New(sha256.New, []byte(s.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	_, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *SpotAccountClient) NewOrder(ctx context.Context, param types.NewOrderParam) (*types.NewOrderResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/order",
		Method:  http.MethodPost,
	}

	st := spotutils.TRADE

	{
		headers, err := s.GenHeaders(st)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.NewOrderParams{
			NewOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: s.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := s.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := s.NeedSignature(st); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(s.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.NewOrderAPIResp
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.NewOrderResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (s *SpotAccountClient) CancelOrder(ctx context.Context, param types.CancelOrderParam) (*types.OrderInfoResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/order",
		Method:  http.MethodDelete,
	}

	st := spotutils.TRADE

	{
		headers, err := s.GenHeaders(st)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.CancelOrderParams{
			CancelOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: s.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := s.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := s.NeedSignature(st); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(s.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.OrderInfo
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.OrderInfoResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (s *SpotAccountClient) CancelOrdersOnOneSymbol(ctx context.Context, param types.CancelOrdersOnOneSymbolParam) (*types.CancelOrdersResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/openOrders",
		Method:  http.MethodDelete,
	}

	st := spotutils.TRADE

	{
		headers, err := s.GenHeaders(st)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.CancelOrdersOnOneSymbolParams{
			CancelOrdersOnOneSymbolParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: s.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := s.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := s.NeedSignature(st); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(s.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.OrderInfo
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.CancelOrdersResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (s *SpotAccountClient) QueryOrder(ctx context.Context, param types.QueryOrderParam) (*types.QueryOrderResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/order",
		Method:  http.MethodGet,
	}

	st := spotutils.USER_DATA

	{
		headers, err := s.GenHeaders(st)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.QueryOrderParams{
			QueryOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: s.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := s.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := s.NeedSignature(st); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(s.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.Order
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.QueryOrderResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (s *SpotAccountClient) GetOpenOrders(ctx context.Context, param types.GetOpenOrdersParam) (*types.GetOpenOrdersResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/openOrders",
		Method:  http.MethodGet,
	}

	st := spotutils.USER_DATA

	{
		headers, err := s.GenHeaders(st)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetOpenOrdersParams{
			GetOpenOrdersParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: s.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := s.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := s.NeedSignature(st); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(s.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.Order
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetOpenOrdersResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (s *SpotAccountClient) GetAllOrders(ctx context.Context, param types.GetAllOrdersParam) (*types.GetAllOrdersResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/allOrders",
		Method:  http.MethodGet,
	}

	st := spotutils.USER_DATA

	{
		headers, err := s.GenHeaders(st)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetAllOrdersParams{
			GetAllOrdersParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: s.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := s.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := s.NeedSignature(st); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(s.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.Order
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetAllOrdersResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (s *SpotAccountClient) GetAccountInfo(ctx context.Context) (*types.GetAccountInfoResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/account",
		Method:  http.MethodGet,
	}

	st := spotutils.USER_DATA

	{
		headers, err := s.GenHeaders(st)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := bnutils.DefaultParam{
			RecvWindow: s.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
		}

		err := s.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := s.NeedSignature(st); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(s.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := s.SendHTTPRequest(ctx, req)
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

func (s *SpotAccountClient) GetTradeList(ctx context.Context, param types.GetTradesParam) (*types.GetTradesResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/myTrades",
		Method:  http.MethodGet,
	}

	st := spotutils.USER_DATA

	{
		headers, err := s.GenHeaders(st)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetTradesParams{
			GetTradesParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: s.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := s.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := s.NeedSignature(st); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(s.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.Trade
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetTradesResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}
