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

package rest

import (
	"context"
	"errors"
	"net/http"

	"github.com/linstohu/nexapi/htx/usdm/rest/types"
	"github.com/linstohu/nexapi/htx/utils"
)

func (ucli *UsdmClient) GetContractInfo(ctx context.Context, param types.GetContractInfoParam) (*types.GetContractInfoResp, error) {
	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-api/v1/swap_contract_info",
		Method:  http.MethodGet,
		Query:   param,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetContractInfoResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (ucli *UsdmClient) GetFundingRate(ctx context.Context) (*types.GetFundingRateResp, error) {
	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-api/v1/swap_batch_funding_rate",
		Method:  http.MethodGet,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetFundingRateResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (ucli *UsdmClient) GetMarketDepth(ctx context.Context, param types.GetMarketDepthParam) (*types.GetMarketDepthResp, error) {
	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-ex/market/depth",
		Method:  http.MethodGet,
		Query:   param,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetMarketDepthResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (ucli *UsdmClient) GetKline(ctx context.Context, param types.GetKlineParam) (*types.GetKlineResp, error) {
	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-ex/market/history/kline",
		Method:  http.MethodGet,
		Query:   param,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetKlineResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (ucli *UsdmClient) GetMarketTicker(ctx context.Context, param types.GetMarketTickerParam) (*types.GetMarketTickerResp, error) {
	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-ex/market/detail/merged",
		Method:  http.MethodGet,
		Query:   param,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetMarketTickerResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (ucli *UsdmClient) GetMarketTickers(ctx context.Context, param types.GetMarketTickersParam) (*types.GetMarketTickersResp, error) {
	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/v2/linear-swap-ex/market/detail/batch_merged",
		Method:  http.MethodGet,
		Query:   param,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetMarketTickersResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}
