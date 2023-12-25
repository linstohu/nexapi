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

package tradingaccount

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/okx/tradingaccount/types"
	okxutils "github.com/linstohu/nexapi/okx/utils"
	"github.com/linstohu/nexapi/utils"
)

type TradingAccountClient struct {
	*okxutils.OKXRestClient

	// validate struct fields
	validate *validator.Validate
}

type TradingAccountClientCfg struct {
	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	Passphrase string `validate:"required"`
	Debug      bool
	// Logger
	Logger *slog.Logger
}

func NewTradingAccountClient(cfg *TradingAccountClientCfg) (*TradingAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := okxutils.NewOKXRestClient(&okxutils.OKXRestClientCfg{
		Debug:      cfg.Debug,
		Logger:     cfg.Logger,
		BaseURL:    cfg.BaseURL,
		Key:        cfg.Key,
		Secret:     cfg.Secret,
		Passphrase: cfg.Passphrase,
	})
	if err != nil {
		return nil, err
	}

	return &TradingAccountClient{
		OKXRestClient: cli,
		validate:      validator,
	}, nil
}

func (t *TradingAccountClient) GetBalance(ctx context.Context, param types.GetBalanceParam) (*types.GetBalanceResp, error) {
	err := t.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   t.GetDebug(),
		BaseURL: t.GetBaseURL(),
		Path:    "/api/v5/account/balance",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := t.GenAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := t.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.GetBalanceResp
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (t *TradingAccountClient) GetPositions(ctx context.Context, param types.GetPositionsParam) (*types.GetPositionsResp, error) {
	err := t.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   t.GetDebug(),
		BaseURL: t.GetBaseURL(),
		Path:    "/api/v5/account/positions",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := t.GenAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := t.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.GetPositionsResp
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	return &body, nil
}
