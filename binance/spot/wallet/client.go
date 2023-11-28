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

package wallet

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	spotutils "github.com/linstohu/nexapi/binance/spot/utils"
	"github.com/linstohu/nexapi/binance/spot/wallet/types"
	bnutils "github.com/linstohu/nexapi/binance/utils"
	"github.com/linstohu/nexapi/utils"
)

type SpotWalletClient struct {
	*spotutils.SpotClient

	// validate struct fields
	validate *validator.Validate
}

type SpotWalletClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewSpotWalletClient(cfg *SpotWalletClientCfg) (*SpotWalletClient, error) {
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

	return &SpotWalletClient{
		SpotClient: cli,
		validate:   validator,
	}, nil
}

func (s *SpotWalletClient) GetSystemStatus(ctx context.Context) (*types.GetSystemStatusResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/sapi/v1/system/status",
		Method:  http.MethodGet,
	}

	headers, err := s.GenHeaders(spotutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.SystemStatus
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetSystemStatusResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (s *SpotWalletClient) GetAllCoinsInfo(ctx context.Context) (*types.GetCoinInfoResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/sapi/v1/capital/config/getall",
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

	var body []*types.CoinInfo
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetCoinInfoResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (s *SpotWalletClient) GetAssetDetail(ctx context.Context, param types.GetAssetDetailParam) (*types.AssetDetailResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/sapi/v1/asset/assetDetail",
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
		query := types.AssetDetailParams{
			GetAssetDetailParam: param,
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

	body := make(map[string]*types.AssetDetail)
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.AssetDetailResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (s *SpotWalletClient) GetTradeFee(ctx context.Context, param types.GetTradeFeeParam) (*types.GetTradeFeeResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/sapi/v1/asset/tradeFee",
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
		query := types.TradeFeeParams{
			GetTradeFeeParam: param,
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

	var body []*types.TradeFee
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetTradeFeeResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (s *SpotWalletClient) UniversalTransfer(ctx context.Context, param types.UniversalTransferParam) (*types.UniversalTransferResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/sapi/v1/asset/transfer",
		Method:  http.MethodPost,
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
		body := types.UniversalTransferParams{
			UniversalTransferParam: param,
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

	var body types.UniversalTransferAPIResp
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.UniversalTransferResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (s *SpotWalletClient) GetUniversalTransferHistory(ctx context.Context, param types.GetUniversalTransferHistoryParam) (*types.GetUniversalTransferHistoryResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/sapi/v1/asset/transfer",
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
		query := types.GetUniversalTransferHistoryParams{
			GetUniversalTransferHistoryParam: param,
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

	var body types.GetUniversalTransferHistory
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetUniversalTransferHistoryResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (s *SpotWalletClient) GetFundingAsset(ctx context.Context, param types.GetFundingAssetParam) (*types.GetFundingAssetResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/sapi/v1/asset/get-funding-asset",
		Method:  http.MethodPost,
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
		body := types.GetFundingAssetParams{
			GetFundingAssetParam: param,
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

	var body []*types.FundingAsset
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetFundingAssetResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (s *SpotWalletClient) GetUserAsset(ctx context.Context, param types.GetUserAssetParam) (*types.GetUserAssetResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/sapi/v3/asset/getUserAsset",
		Method:  http.MethodPost,
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
		body := types.GetUserAssetParams{
			GetUserAssetParam: param,
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

	var body []*types.UserAsset
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetUserAssetResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (s *SpotWalletClient) GetApiRestrictions(ctx context.Context) (*types.GetApiRestrictionsResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/sapi/v1/account/apiRestrictions",
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

	var body types.ApiRestrictions
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetApiRestrictionsResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (s *SpotWalletClient) GetWalletBalance(ctx context.Context) (*types.GetWalletBalanceResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/sapi/v1/asset/wallet/balance",
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

	var body []*types.Balance
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetWalletBalanceResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (s *SpotWalletClient) GetEarnAccount(ctx context.Context) (*types.GetEarnAccountResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/sapi/v1/simple-earn/account",
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

	var body types.EarnAccount
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetEarnAccountResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}
