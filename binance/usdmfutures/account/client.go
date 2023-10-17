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
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/usdmfutures/account/types"
	umutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type UsdMFuturesAccountClient struct {
	*umutils.USDMarginedClient

	// validate struct fields
	validate *validator.Validate
}

type UsdMFuturesAccountClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewUsdMFuturesAccountClient(cfg *umutils.USDMarginedClientCfg) (*UsdMFuturesAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := umutils.NewUSDMarginedClient(&umutils.USDMarginedClientCfg{
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

	return &UsdMFuturesAccountClient{
		USDMarginedClient: cli,
		validate:          validator,
	}, nil
}

func (u *UsdMFuturesAccountClient) ChangePositionMode(ctx context.Context, param types.ChangePositionModeParam) (*types.Response, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/positionSide/dual",
		Method:       http.MethodPost,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.ChangePositionModeParams{
			ChangePositionModeParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: u.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := u.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Response
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (u *UsdMFuturesAccountClient) GetPositionMode(ctx context.Context) (*types.GetCurrentPositionModeResp, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/positionSide/dual",
		Method:       http.MethodGet,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := bnutils.DefaultParam{
			RecvWindow: u.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
		}

		err := u.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetCurrentPositionModeResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (u *UsdMFuturesAccountClient) ChangeMultiAssetsMode(ctx context.Context, param types.ChangeMultiAssetsModeParam) (*types.Response, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/multiAssetsMargin",
		Method:       http.MethodPost,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.ChangeMultiAssetsModeParams{
			ChangeMultiAssetsModeParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: u.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := u.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Response
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (u *UsdMFuturesAccountClient) GetMultiAssetsMode(ctx context.Context) (*types.GetCurrentMultiAssetsModeResp, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/multiAssetsMargin",
		Method:       http.MethodGet,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := bnutils.DefaultParam{
			RecvWindow: u.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
		}

		err := u.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetCurrentMultiAssetsModeResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (u *UsdMFuturesAccountClient) NewOrder(ctx context.Context, param types.NewOrderParam) (*types.Order, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/order",
		Method:       http.MethodPost,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.NewOrderParams{
			NewOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: u.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := u.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Order
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (u *UsdMFuturesAccountClient) QueryOrder(ctx context.Context, param types.GetOrderParam) (*types.Order, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/order",
		Method:       http.MethodGet,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetOrderParams{
			GetOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: u.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := u.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Order
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (u *UsdMFuturesAccountClient) QueryOpenOrder(ctx context.Context, param types.GetOrderParam) (*types.Order, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/openOrder",
		Method:       http.MethodGet,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetOrderParams{
			GetOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: u.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := u.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Order
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (u *UsdMFuturesAccountClient) QueryAllOpenOrders(ctx context.Context, param types.GetAllOpenOrdersParam) ([]*types.Order, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/openOrders",
		Method:       http.MethodGet,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetAllOpenOrdersParams{
			GetAllOpenOrdersParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: u.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := u.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.Order
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *UsdMFuturesAccountClient) CancelOrder(ctx context.Context, param types.GetOrderParam) (*types.Order, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/order",
		Method:       http.MethodDelete,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.GetOrderParams{
			GetOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: u.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := u.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Order
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (u *UsdMFuturesAccountClient) CancelAllOpenOrders(ctx context.Context, param types.CancelAllOpenOrdersParam) error {
	req := umutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/allOpenOrders",
		Method:       http.MethodDelete,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return err
		}
		req.Headers = headers
	}

	{
		body := types.CancelAllOpenOrdersParams{
			CancelAllOpenOrdersParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: u.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := u.validate.Struct(body)
		if err != nil {
			return err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	_, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (u *UsdMFuturesAccountClient) GetAllOrders(ctx context.Context, param types.GetAllOrdersParam) ([]*types.Order, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/allOrders",
		Method:       http.MethodGet,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetAllOrdersParams{
			GetAllOrdersParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: u.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := u.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.Order
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *UsdMFuturesAccountClient) GetBalance(ctx context.Context) ([]*types.Balance, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v2/balance",
		Method:       http.MethodGet,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := bnutils.DefaultParam{
			RecvWindow: u.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
		}

		err := u.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.Balance
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *UsdMFuturesAccountClient) GetAccountInformation(ctx context.Context) (*types.Account, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v2/account",
		Method:       http.MethodGet,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := bnutils.DefaultParam{
			RecvWindow: u.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
		}

		err := u.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Account
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (u *UsdMFuturesAccountClient) ChangeInitialLeverage(ctx context.Context, param types.ChangeLeverageParam) (*types.ChangeLeverageResp, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/leverage",
		Method:       http.MethodPost,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.ChangeLeverageParams{
			ChangeLeverageParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: u.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := u.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.ChangeLeverageResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (u *UsdMFuturesAccountClient) ChangeMarginType(ctx context.Context, param types.ChangeMarginTypeParam) error {
	req := umutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/marginType",
		Method:       http.MethodPost,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return err
		}
		req.Headers = headers
	}

	{
		body := types.ChangeMarginTypeParams{
			ChangeMarginTypeParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: u.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := u.validate.Struct(body)
		if err != nil {
			return err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	_, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (u *UsdMFuturesAccountClient) ModifyIsolatedPositionMargin(ctx context.Context, param types.ModifyIsolatedPositionMarginParam) (*types.ModifyIsolatedPositionMarginResp, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/positionMargin",
		Method:       http.MethodPost,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.ModifyIsolatedPositionMarginParams{
			ModifyIsolatedPositionMarginParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: u.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := u.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.ModifyIsolatedPositionMarginResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (u *UsdMFuturesAccountClient) GetPositionInformation(ctx context.Context, param types.GetPositionParam) ([]*types.Position, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v2/positionRisk",
		Method:       http.MethodGet,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetPositionParams{
			GetPositionParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: u.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := u.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.Position
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *UsdMFuturesAccountClient) GetAccountTradeList(ctx context.Context, param types.GetTradeListParam) ([]*types.Trade, error) {
	req := umutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/userTrades",
		Method:       http.MethodGet,
	}
	{
		headers, err := u.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetTradeListParams{
			GetTradeListParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: u.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := u.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := u.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(u.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.Trade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
