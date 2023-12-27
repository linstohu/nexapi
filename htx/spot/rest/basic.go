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

	"github.com/linstohu/nexapi/htx/spot/rest/types"
	"github.com/linstohu/nexapi/htx/utils"
)

func (scli *SpotClient) GetSymbols(ctx context.Context, param types.GetSymbolsParam) (*types.GetSymbolsResp, error) {
	req := utils.HTTPRequest{
		BaseURL: scli.cli.GetBaseURL(),
		Path:    "/v2/settings/common/symbols",
		Method:  http.MethodGet,
		Query:   param,
	}

	{
		headers, err := scli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := scli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetSymbolsResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (scli *SpotClient) GetSymbolInfo(ctx context.Context) (*types.GetSymbolInfoResp, error) {
	req := utils.HTTPRequest{
		BaseURL: scli.cli.GetBaseURL(),
		Path:    "/v1/settings/common/market-symbols",
		Method:  http.MethodGet,
	}

	{
		headers, err := scli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := scli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetSymbolInfoResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (scli *SpotClient) GetMarketDepth(ctx context.Context, param types.GetMarketDepthParam) (*types.GetMarketDepthResp, error) {
	req := utils.HTTPRequest{
		BaseURL: scli.cli.GetBaseURL(),
		Path:    "/market/depth",
		Method:  http.MethodGet,
		Query:   param,
	}

	{
		headers, err := scli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := scli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	var ret types.GetMarketDepthResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(err.Error())
	}

	return &ret, nil
}

func (scli *SpotClient) GetKline(ctx context.Context, param types.GetKlineParam) (*types.GetKlineResp, error) {
	req := utils.HTTPRequest{
		BaseURL: scli.cli.GetBaseURL(),
		Path:    "/market/history/kline",
		Method:  http.MethodGet,
		Query:   param,
	}

	{
		headers, err := scli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := scli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetKlineResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (scli *SpotClient) GetAllMarketTickers(ctx context.Context) (*types.GetMarketTickersResp, error) {
	req := utils.HTTPRequest{
		BaseURL: scli.cli.GetBaseURL(),
		Path:    "/market/tickers",
		Method:  http.MethodGet,
	}

	{
		headers, err := scli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := scli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetMarketTickersResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}
