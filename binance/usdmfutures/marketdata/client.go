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
	spottypes "github.com/linstohu/nexapi/binance/spot/marketdata/types"
	"github.com/linstohu/nexapi/binance/usdmfutures/marketdata/types"
	umutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	"github.com/linstohu/nexapi/utils"
	"github.com/valyala/fastjson"
)

type USDMFuturesMarketDataClient struct {
	*umutils.USDMarginedClient

	// validate struct fields
	validate *validator.Validate
}

func NewUSDMFuturesMarketDataClient(cfg *umutils.USDMarginedClientCfg) (*USDMFuturesMarketDataClient, error) {
	cli, err := umutils.NewUSDMarginedClient(cfg)
	if err != nil {
		return nil, err
	}

	validator := validator.New()

	return &USDMFuturesMarketDataClient{
		USDMarginedClient: cli,
		validate:          validator,
	}, nil
}

func (u *USDMFuturesMarketDataClient) Ping(ctx context.Context) error {
	req := utils.HTTPRequest{
		Debug:   u.GetDebug(),
		BaseURL: u.GetBaseURL(),
		Path:    "/fapi/v1/ping",
		Method:  http.MethodGet,
	}

	headers, err := u.GenHeaders(umutils.NONE)
	if err != nil {
		return err
	}
	req.Headers = headers

	_, err = u.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (u *USDMFuturesMarketDataClient) GetServerTime(ctx context.Context) (*spottypes.ServerTimeResp, error) {
	req := utils.HTTPRequest{
		Debug:   u.GetDebug(),
		BaseURL: u.GetBaseURL(),
		Path:    "/fapi/v1/time",
		Method:  http.MethodGet,
	}

	headers, err := u.GenHeaders(umutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
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

func (u *USDMFuturesMarketDataClient) GetExchangeInfo(ctx context.Context) (*types.GetExchangeInfoResp, error) {
	req := utils.HTTPRequest{
		Debug:   u.GetDebug(),
		BaseURL: u.GetBaseURL(),
		Path:    "/fapi/v1/exchangeInfo",
		Method:  http.MethodGet,
	}

	headers, err := u.GenHeaders(umutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
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

func (u *USDMFuturesMarketDataClient) GetOrderbook(ctx context.Context, param types.GetOrderbookParams) (*types.GetOrderbookResp, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   u.GetDebug(),
		BaseURL: u.GetBaseURL(),
		Path:    "/fapi/v1/depth",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := u.GenHeaders(umutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
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

func (u *USDMFuturesMarketDataClient) GetRecentTradeList(ctx context.Context, param types.GetTradeParams) (*types.GetTradeResp, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   u.GetDebug(),
		BaseURL: u.GetBaseURL(),
		Path:    "/fapi/v1/trades",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := u.GenHeaders(umutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
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

func (u *USDMFuturesMarketDataClient) GetAggTrades(ctx context.Context, param types.GetAggTradesParam) (*types.GetAggTradesResp, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   u.GetDebug(),
		BaseURL: u.GetBaseURL(),
		Path:    "/fapi/v1/aggTrades",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := u.GenHeaders(umutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.AggTrade
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetAggTradesResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (u *USDMFuturesMarketDataClient) GetKlines(ctx context.Context, param types.GetKlineParam) (*spottypes.GetKlineResp, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   u.GetDebug(),
		BaseURL: u.GetBaseURL(),
		Path:    "/fapi/v1/klines",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := u.GenHeaders(umutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
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

	var ret []*spottypes.Kline
	for _, v := range arr {
		kline, err := v.Array()
		if err != nil {
			return nil, err
		}

		if len(kline) != 12 {
			return nil, fmt.Errorf("unknown kline value: %s", v.String())
		}

		ret = append(ret, &spottypes.Kline{
			OpenTime:                 kline[0].GetInt64(),
			OpenPrice:                kline[1].String(),
			HighPrice:                kline[2].String(),
			LowPrice:                 kline[3].String(),
			ClosePrice:               kline[4].String(),
			Volume:                   kline[5].String(),
			CloseTime:                kline[6].GetInt64(),
			QuoteAssetVolume:         kline[7].String(),
			NumberOfTrades:           kline[8].GetInt64(),
			TakerBuyBaseAssetVolume:  kline[9].String(),
			TakerBuyQuoteAssetVolume: kline[10].String(),
		})
	}

	data := &spottypes.GetKlineResp{
		Http: resp,
		Body: ret,
	}

	return data, nil
}

func (u *USDMFuturesMarketDataClient) GetMarkPriceForSymbol(ctx context.Context, param types.GetMarkPriceParam) (*types.GetMarkPriceResp, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   u.GetDebug(),
		BaseURL: u.GetBaseURL(),
		Path:    "/fapi/v1/premiumIndex",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := u.GenHeaders(umutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.MarkPrice

	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetMarkPriceResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (u *USDMFuturesMarketDataClient) GetMarkPriceForAllSymbols(ctx context.Context) (*types.GetMarkPriceForAllSymbolsResp, error) {
	req := utils.HTTPRequest{
		Debug:   u.GetDebug(),
		BaseURL: u.GetBaseURL(),
		Path:    "/fapi/v1/premiumIndex",
		Method:  http.MethodGet,
	}

	headers, err := u.GenHeaders(umutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.MarkPrice
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetMarkPriceForAllSymbolsResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (u *USDMFuturesMarketDataClient) GetFundingRateHistory(ctx context.Context, param types.GetFundingRateParam) (*types.GetFundingRateResp, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   u.GetDebug(),
		BaseURL: u.GetBaseURL(),
		Path:    "/fapi/v1/fundingRate",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := u.GenHeaders(umutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
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

func (u *USDMFuturesMarketDataClient) GetTickerPriceForSymbol(ctx context.Context, param types.GetTickerPriceParam) (*types.GetTickerPriceForSymbolResp, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   u.GetDebug(),
		BaseURL: u.GetBaseURL(),
		Path:    "/fapi/v1/ticker/price",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := u.GenHeaders(umutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.TickerPrice
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetTickerPriceForSymbolResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (u *USDMFuturesMarketDataClient) GetTickerPriceForAllSymbols(ctx context.Context) (*types.GetTickerPriceForSymbolsResp, error) {
	req := utils.HTTPRequest{
		Debug:   u.GetDebug(),
		BaseURL: u.GetBaseURL(),
		Path:    "/fapi/v1/ticker/price",
		Method:  http.MethodGet,
	}

	headers, err := u.GenHeaders(umutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.TickerPrice
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetTickerPriceForSymbolsResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (u *USDMFuturesMarketDataClient) GetBookTickerForSymbol(ctx context.Context, param types.GetBookTickerForSymbolParam) (*types.GetBookTickerForSymbolResp, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   u.GetDebug(),
		BaseURL: u.GetBaseURL(),
		Path:    "/fapi/v1/ticker/bookTicker",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := u.GenHeaders(umutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.BookTicker
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetBookTickerForSymbolResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (u *USDMFuturesMarketDataClient) GetBookTickerForAllSymbols(ctx context.Context) (*types.GetBookTickerForSymbolsResp, error) {
	req := utils.HTTPRequest{
		Debug:   u.GetDebug(),
		BaseURL: u.GetBaseURL(),
		Path:    "/fapi/v1/ticker/bookTicker",
		Method:  http.MethodGet,
	}

	headers, err := u.GenHeaders(umutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.BookTicker
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetBookTickerForSymbolsResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (u *USDMFuturesMarketDataClient) GetOpenInterestHistory(ctx context.Context, param types.GetOpenInterestHistParam) (*types.GetOpenInterestHistResp, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   u.GetDebug(),
		BaseURL: u.GetBaseURL(),
		Path:    "/futures/data/openInterestHist",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := u.GenHeaders(umutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
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
