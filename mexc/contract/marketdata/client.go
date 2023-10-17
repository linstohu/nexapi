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

package marketdata

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/mexc/contract/marketdata/types"
	ctutils "github.com/linstohu/nexapi/mexc/contract/utils"
)

type ContractMarketDataClient struct {
	*ctutils.ContractClient

	// validate struct fields
	validate *validator.Validate
}

func NewContractMarketDataClient(cfg *ctutils.ContractClientCfg) (*ContractMarketDataClient, error) {
	cli, err := ctutils.NewContractClient(cfg)
	if err != nil {
		return nil, err
	}

	validator := validator.New()

	return &ContractMarketDataClient{
		ContractClient: cli,
		validate:       validator,
	}, nil
}

func (s *ContractMarketDataClient) GetServerTime(ctx context.Context) (*types.ServerTime, error) {
	req := ctutils.HTTPRequest{
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v1/contract/ping",
		Method:  http.MethodGet,
	}

	headers, err := s.GenPubHeaders()
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.ServerTime
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *ContractMarketDataClient) GetContractDetails(ctx context.Context, param types.GetContractDetailsParams) (*types.GetContractDetailsResp, error) {
	req := ctutils.HTTPRequest{
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v1/contract/detail",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := s.GenPubHeaders()
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetContractDetailsResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *ContractMarketDataClient) GetTickerForSymbol(ctx context.Context, param types.GetTickerForSymbolParam) (*types.GetTickerResp, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := ctutils.HTTPRequest{
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v1/contract/ticker",
		Method:  http.MethodGet,
	}

	query := types.GetTickerForSymbolParam{
		Symbol: param.Symbol,
	}
	req.Query = query

	headers, err := s.GenPubHeaders()
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetTickerResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *ContractMarketDataClient) GetTickerForAllSymbols(ctx context.Context) (*types.GetAllTickersResp, error) {
	req := ctutils.HTTPRequest{
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v1/contract/ticker",
		Method:  http.MethodGet,
	}

	query := types.TickerParams{}
	req.Query = query

	headers, err := s.GenPubHeaders()
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetAllTickersResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
