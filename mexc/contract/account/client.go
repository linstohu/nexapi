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
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/mexc/contract/account/types"
	ctutils "github.com/linstohu/nexapi/mexc/contract/utils"
)

type ContractAccountClient struct {
	*ctutils.ContractClient

	// validate struct fields
	validate *validator.Validate
}

type ContractAccountClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewContractAccountClient(cfg *ctutils.ContractClientCfg) (*ContractAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := ctutils.NewContractClient(&ctutils.ContractClientCfg{
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

	return &ContractAccountClient{
		ContractClient: cli,
		validate:       validator,
	}, nil
}

func (c *ContractAccountClient) GetAccountAsset(ctx context.Context) (*types.GetAccountAsset, error) {
	req := ctutils.HTTPRequest{
		BaseURL: c.GetBaseURL(),
		Path:    "/api/v1/private/account/assets",
		Method:  http.MethodGet,
	}

	{
		headers, err := c.GenAuthHeaders(req)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetAccountAsset
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *ContractAccountClient) GetOpenPositions(ctx context.Context, param types.GetOpenPositionsParams) (*types.GetOpenPositions, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := ctutils.HTTPRequest{
		BaseURL: c.GetBaseURL(),
		Path:    "/api/v1/private/position/open_positions",
		Method:  http.MethodGet,
		Query:   param,
	}

	{
		headers, err := c.GenAuthHeaders(req)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetOpenPositions
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *ContractAccountClient) GetPositionLeverage(ctx context.Context, param types.GetLeverageParams) (*types.GetLeverageResp, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := ctutils.HTTPRequest{
		BaseURL: c.GetBaseURL(),
		Path:    "/api/v1/private/position/leverage",
		Method:  http.MethodGet,
		Query:   param,
	}

	{
		headers, err := c.GenAuthHeaders(req)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetLeverageResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *ContractAccountClient) SetPositionLeverage(ctx context.Context, param types.SetLeverageParams) (*types.SetLeverageResp, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := ctutils.HTTPRequest{
		BaseURL: c.GetBaseURL(),
		Path:    "/api/v1/private/position/change_leverage",
		Method:  http.MethodPost,
		Body:    param,
	}

	{
		headers, err := c.GenAuthHeaders(req)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.SetLeverageResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
