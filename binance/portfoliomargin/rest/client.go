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

package pmrest

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/portfoliomargin/rest/types"
	pmutils "github.com/linstohu/nexapi/binance/portfoliomargin/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
	"github.com/linstohu/nexapi/utils"
)

type PortfolioMarginAccountClient struct {
	*pmutils.PortfolioMarginClient

	// validate struct fields
	validate *validator.Validate
}

type PortfolioMarginAccountClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewPortfolioMarginAccountClient(cfg *pmutils.PortfolioMarginClientCfg) (*PortfolioMarginAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := pmutils.NewPortfolioMarginClient(&pmutils.PortfolioMarginClientCfg{
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

	return &PortfolioMarginAccountClient{
		PortfolioMarginClient: cli,
		validate:              validator,
	}, nil
}

func (p *PortfolioMarginAccountClient) GetAssetBalance(ctx context.Context, param types.GetBalanceParam) (*types.GetBalanceResp, error) {
	err := p.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   p.GetDebug(),
		BaseURL: p.GetBaseURL(),
		Path:    "/papi/v1/balance",
		Method:  http.MethodGet,
	}

	securityType := pmutils.USER_DATA

	{
		headers, err := p.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetBalanceParams{
			GetBalanceParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: p.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := p.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := p.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(p.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := p.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.Balance
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetBalanceResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (p *PortfolioMarginAccountClient) GetBalance(ctx context.Context) (*types.GetAllBalanceResp, error) {
	req := utils.HTTPRequest{
		Debug:   p.GetDebug(),
		BaseURL: p.GetBaseURL(),
		Path:    "/papi/v1/balance",
		Method:  http.MethodGet,
	}

	securityType := pmutils.USER_DATA

	{
		headers, err := p.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetAllBalanceParams{
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: p.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := p.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := p.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(p.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := p.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.Balance
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetAllBalanceResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (p *PortfolioMarginAccountClient) GetAccountInformation(ctx context.Context) (*types.GetAccountInfoResp, error) {
	req := utils.HTTPRequest{
		Debug:   p.GetDebug(),
		BaseURL: p.GetBaseURL(),
		Path:    "/papi/v1/account",
		Method:  http.MethodGet,
	}

	securityType := pmutils.USER_DATA

	{
		headers, err := p.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := bnutils.DefaultParam{
			RecvWindow: p.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
		}

		err := p.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := p.NeedSignature(securityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(p.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := p.SendHTTPRequest(ctx, req)
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
