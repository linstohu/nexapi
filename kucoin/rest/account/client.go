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
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/kucoin/rest/account/types"
	"github.com/linstohu/nexapi/kucoin/rest/utils"
)

type AccountClient struct {
	cli *utils.KucoinClient

	// validate struct fields
	validate *validator.Validate
}

type AccountClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	KeyVersion string `validate:"required"`
	Secret     string `validate:"required"`
	Passphrase string `validate:"required"`
}

func NewAccountClient(cfg *AccountClientCfg) (*AccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := utils.NewKucoinRestClient(&utils.KucoinClientCfg{
		Debug:      cfg.Debug,
		Logger:     cfg.Logger,
		BaseURL:    cfg.BaseURL,
		Key:        cfg.Key,
		KeyVersion: cfg.KeyVersion,
		Secret:     cfg.Secret,
		Passphrase: cfg.Passphrase,
	})
	if err != nil {
		return nil, err
	}

	return &AccountClient{
		cli:      cli,
		validate: validator,
	}, nil
}

func (a *AccountClient) GetAccountList(ctx context.Context, param types.GetAccountListParam) ([]*types.AccountModel, error) {
	req := utils.HTTPRequest{
		BaseURL: a.cli.GetBaseURL(),
		Path:    "/api/v1/accounts",
		Method:  http.MethodGet,
		Query:   param,
	}

	{
		headers, err := a.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		err := a.validate.Struct(param)
		if err != nil {
			return nil, err
		}

		h, err := a.cli.GenSignature(req)
		if err != nil {
			return nil, err
		}
		for k, v := range h {
			req.Headers[k] = v
		}
	}

	resp, err := a.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	ar := &utils.ApiResponse{Resp: resp}
	if err := resp.ReadJsonBody(ar); err != nil {
		return nil, errors.New(resp.Error())
	}

	var ret []*types.AccountModel
	if err := ar.ReadData(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}
