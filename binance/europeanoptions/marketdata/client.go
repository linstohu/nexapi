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
	"net/http"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/europeanoptions/marketdata/types"
	eoutils "github.com/linstohu/nexapi/binance/europeanoptions/utils"
	spottypes "github.com/linstohu/nexapi/binance/spot/marketdata/types"
	usdmtypes "github.com/linstohu/nexapi/binance/usdmfutures/marketdata/types"
	usdmutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	"github.com/linstohu/nexapi/utils"
)

type OptionsMarketDataClient struct {
	*eoutils.OptionsClient

	// validate struct fields
	validate *validator.Validate
}

func NewOptionsMarketDataClient(cfg *eoutils.OptionsClientCfg) (*OptionsMarketDataClient, error) {
	cli, err := eoutils.NewOptionsClient(cfg)
	if err != nil {
		return nil, err
	}

	validator := validator.New()

	return &OptionsMarketDataClient{
		OptionsClient: cli,
		validate:      validator,
	}, nil
}

func (o *OptionsMarketDataClient) Ping(ctx context.Context) error {
	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/ping",
		Method:  http.MethodGet,
	}

	headers, err := o.GenHeaders(usdmutils.NONE)
	if err != nil {
		return err
	}
	req.Headers = headers

	_, err = o.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (o *OptionsMarketDataClient) GetServerTime(ctx context.Context) (*spottypes.ServerTimeResp, error) {
	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/time",
		Method:  http.MethodGet,
	}

	headers, err := o.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body spottypes.ServerTime
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &spottypes.ServerTimeResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (o *OptionsMarketDataClient) GetExchangeInfo(ctx context.Context) (*types.GetExchangeInfoResp, error) {
	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/exchangeInfo",
		Method:  http.MethodGet,
	}

	headers, err := o.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.ExchangeInfo

	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetExchangeInfoResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (o *OptionsMarketDataClient) GetOrderbook(ctx context.Context, param types.GetOrderbookParams) (*types.GetOrderbookResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/depth",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := o.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.Orderbook

	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetOrderbookResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (o *OptionsMarketDataClient) GetRecentTradesList(ctx context.Context, param types.GetTradeParams) (*types.GetTradeResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/trades",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := o.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.Trade
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetTradeResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (o *OptionsMarketDataClient) GetKlines(ctx context.Context, param usdmtypes.GetKlineParam) (*types.GetKlineResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/klines",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := o.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.Kline
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetKlineResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (o *OptionsMarketDataClient) GetMarkPrice(ctx context.Context, param types.GetMarkPriceParam) (*types.GetMarkPriceResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/mark",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := o.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.MarkPrice
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetMarkPriceResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (o *OptionsMarketDataClient) GetTickerPrice(ctx context.Context, param types.GetTickerPriceParam) (*types.GetTickerPriceResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/ticker",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := o.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.TickerPrice
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetTickerPriceResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (o *OptionsMarketDataClient) GetUnderlyingIndexPrice(ctx context.Context, param types.GetUnderlyingIndexPriceParams) (*types.GetUnderlyingIndexPriceResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/index",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := o.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.UnderlyingIndexPrice
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetUnderlyingIndexPriceResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (o *OptionsMarketDataClient) GetOpenInterest(ctx context.Context, param types.GetOpenInterestParam) (*types.GetOpenInterestResp, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   o.GetDebug(),
		BaseURL: o.GetBaseURL(),
		Path:    "/eapi/v1/openInterest",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := o.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.OpenInterest
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetOpenInterestResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}
