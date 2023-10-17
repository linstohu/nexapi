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
	"github.com/linstohu/nexapi/binance/europeanoptions/marketdata/types"
	eoutils "github.com/linstohu/nexapi/binance/europeanoptions/utils"
	spottypes "github.com/linstohu/nexapi/binance/spot/marketdata/types"
	usdmtypes "github.com/linstohu/nexapi/binance/usdmfutures/marketdata/types"
	usdmutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
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
	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/ping",
		Method:       http.MethodGet,
	}

	headers, err := o.GenHeaders(req.SecurityType)
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

func (o *OptionsMarketDataClient) GetServerTime(ctx context.Context) (*spottypes.ServerTime, error) {
	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/time",
		Method:       http.MethodGet,
	}

	headers, err := o.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret spottypes.ServerTime
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (o *OptionsMarketDataClient) GetExchangeInfo(ctx context.Context) (*types.ExchangeInfo, error) {
	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/exchangeInfo",
		Method:       http.MethodGet,
	}

	headers, err := o.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.ExchangeInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (o *OptionsMarketDataClient) GetOrderbook(ctx context.Context, param types.GetOrderbookParams) (*types.Orderbook, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/depth",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := o.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Orderbook
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (o *OptionsMarketDataClient) GetRecentTradesList(ctx context.Context, param types.GetTradeParams) ([]*types.Trade, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/trades",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := o.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.Trade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (o *OptionsMarketDataClient) GetKlines(ctx context.Context, param usdmtypes.GetKlineParam) ([]*types.Kline, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/klines",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := o.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.Kline
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (o *OptionsMarketDataClient) GetMarkPrice(ctx context.Context, param types.GetMarkPriceParam) ([]*types.MarkPrice, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/mark",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := o.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.MarkPrice
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (o *OptionsMarketDataClient) GetTickerPrice(ctx context.Context, param types.GetTickerPriceParam) ([]*types.TickerPrice, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/ticker",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := o.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.TickerPrice
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (o *OptionsMarketDataClient) GetUnderlyingIndexPrice(ctx context.Context, param types.GetUnderlyingIndexPriceParams) (*types.UnderlyingIndexPrice, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/index",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := o.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.UnderlyingIndexPrice
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (o *OptionsMarketDataClient) GetOpenInterest(ctx context.Context, param types.GetOpenInterestParam) ([]*types.OpenInterest, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/openInterest",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := o.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.OpenInterest
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
