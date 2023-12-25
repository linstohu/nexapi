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
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/coinmfutures/marketdata/types"
	cmutils "github.com/linstohu/nexapi/binance/coinmfutures/utils"
	spottypes "github.com/linstohu/nexapi/binance/spot/marketdata/types"
	usdmtypes "github.com/linstohu/nexapi/binance/usdmfutures/marketdata/types"
	usdmutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	"github.com/linstohu/nexapi/utils"
	"github.com/valyala/fastjson"
)

type CoinMFuturesMarketDataClient struct {
	*cmutils.CoinMarginedClient

	// validate struct fields
	validate *validator.Validate
}

func NewCoinMFuturesMarketDataClient(cfg *cmutils.CoinMarginedClientCfg) (*CoinMFuturesMarketDataClient, error) {
	cli, err := cmutils.NewCoinMarginedClient(cfg)
	if err != nil {
		return nil, err
	}

	validator := validator.New()

	return &CoinMFuturesMarketDataClient{
		CoinMarginedClient: cli,
		validate:           validator,
	}, nil
}

func (c *CoinMFuturesMarketDataClient) Ping(ctx context.Context) error {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/ping",
		Method:  http.MethodGet,
	}

	headers, err := c.GenHeaders(usdmutils.NONE)
	if err != nil {
		return err
	}
	req.Headers = headers

	_, err = c.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (c *CoinMFuturesMarketDataClient) GetServerTime(ctx context.Context) (*spottypes.ServerTimeResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/time",
		Method:  http.MethodGet,
	}

	headers, err := c.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
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

func (c *CoinMFuturesMarketDataClient) GetExchangeInfo(ctx context.Context) (*types.GetExchangeInfoResp, error) {
	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/exchangeInfo",
		Method:  http.MethodGet,
	}

	headers, err := c.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
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

func (c *CoinMFuturesMarketDataClient) GetOrderbook(ctx context.Context, param types.GetOrderbookParams) (*types.GetOrderbookResp, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/depth",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := c.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
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

func (c *CoinMFuturesMarketDataClient) GetRecentTradeList(ctx context.Context, param types.GetTradeParams) (*types.GetTradeResp, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/trades",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := c.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
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

func (c *CoinMFuturesMarketDataClient) GetAggTrades(ctx context.Context, param usdmtypes.GetAggTradesParam) (*usdmtypes.GetAggTradesResp, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/aggTrades",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := c.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*usdmtypes.AggTrade
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &usdmtypes.GetAggTradesResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (c *CoinMFuturesMarketDataClient) GetMarkPrice(ctx context.Context, param types.GetMarkPriceParam) (*types.GetMarkPriceResp, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/premiumIndex",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := c.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
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

func (c *CoinMFuturesMarketDataClient) GetFundingRateHistory(ctx context.Context, param types.GetFundingRateParam) (*types.GetFundingRateResp, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/fundingRate",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := c.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.FundingRate
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetFundingRateResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (c *CoinMFuturesMarketDataClient) GetKlines(ctx context.Context, param usdmtypes.GetKlineParam) (*types.GetKlineResp, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/klines",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := c.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	body, err := resp.ReadBody()
	if err != nil {
		return nil, err
	}

	var p fastjson.Parser
	js, err := p.ParseBytes(body)
	if err != nil {
		return nil, err
	}

	arr, err := js.Array()
	if err != nil {
		return nil, err
	}

	var ret []*types.Kline
	for _, v := range arr {
		kline, err := v.Array()
		if err != nil {
			return nil, err
		}

		if len(kline) != 12 {
			return nil, fmt.Errorf("unknown kline value: %s", v.String())
		}

		ret = append(ret, &types.Kline{
			OpenTime:                kline[0].GetInt64(),
			OpenPrice:               kline[1].String(),
			HighPrice:               kline[2].String(),
			LowPrice:                kline[3].String(),
			ClosePrice:              kline[4].String(),
			Volume:                  kline[5].String(),
			CloseTime:               kline[6].GetInt64(),
			BaseAssetVolume:         kline[7].String(),
			NumberOfTrades:          kline[8].GetInt64(),
			TakerBuyVolume:          kline[9].String(),
			TakerBuyBaseAssetVolume: kline[10].String(),
		})
	}

	data := &types.GetKlineResp{
		Http: resp,
		Body: ret,
	}

	return data, nil
}

func (c *CoinMFuturesMarketDataClient) GetTickerPrice(ctx context.Context, param types.GetPriceTickerParam) (*types.GetPriceTickerResp, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/ticker/price",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := c.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.PriceTicker
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetPriceTickerResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (c *CoinMFuturesMarketDataClient) GetBookTicker(ctx context.Context, param types.GetBookTickerParam) (*types.GetBookTickerResp, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/ticker/bookTicker",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := c.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.BookTicker
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetBookTickerResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (c *CoinMFuturesMarketDataClient) GetOpenInterest(ctx context.Context, param types.GetOpenInterestParam) (*types.GetOpenInterestResp, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/dapi/v1/openInterest",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := c.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.OpenInterest

	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetOpenInterestResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (c *CoinMFuturesMarketDataClient) GetOpenInterestHistory(ctx context.Context, param types.GetOpenInterestHistParam) (*types.GetOpenInterestHistResp, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   c.GetDebug(),
		BaseURL: c.GetBaseURL(),
		Path:    "/futures/data/openInterestHist",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := c.GenHeaders(usdmutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.OpenInterestHist
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetOpenInterestHistResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}
