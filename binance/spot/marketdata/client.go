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
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/spot/marketdata/types"
	spotutils "github.com/linstohu/nexapi/binance/spot/utils"
	"github.com/linstohu/nexapi/utils"
	"github.com/valyala/fastjson"
)

type SpotMarketDataClient struct {
	*spotutils.SpotClient

	// validate struct fields
	validate *validator.Validate
}

func NewSpotMarketDataClient(cfg *spotutils.SpotClientCfg) (*SpotMarketDataClient, error) {
	cli, err := spotutils.NewSpotClient(cfg)
	if err != nil {
		return nil, err
	}

	validator := validator.New()

	return &SpotMarketDataClient{
		SpotClient: cli,
		validate:   validator,
	}, nil
}

func (s *SpotMarketDataClient) Ping(ctx context.Context) error {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/ping",
		Method:  http.MethodGet,
	}

	headers, err := s.GenHeaders(spotutils.NONE)
	if err != nil {
		return err
	}
	req.Headers = headers

	_, err = s.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *SpotMarketDataClient) GetServerTime(ctx context.Context) (*types.ServerTimeResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/time",
		Method:  http.MethodGet,
	}

	headers, err := s.GenHeaders(spotutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.ServerTime

	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.ServerTimeResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (s *SpotMarketDataClient) GetExchangeInfo(ctx context.Context, param types.GetExchangeInfoParam) (*types.GetExchangeInfoResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/exchangeInfo",
		Method:  http.MethodGet,
	}

	query := types.GetExchangeInfoParams{}
	if len(param.Symbols) != 0 {
		stringsJson, err := json.Marshal(param.Symbols)
		if err != nil {
			return nil, err
		}
		query.Symbols = string(stringsJson)
	}
	if len(param.Permissions) != 0 {
		stringsJson, err := json.Marshal(param.Permissions)
		if err != nil {
			return nil, err
		}
		query.Permissions = string(stringsJson)
	}
	req.Query = query

	headers, err := s.GenHeaders(spotutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
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

func (s *SpotMarketDataClient) GetOrderbook(ctx context.Context, param types.GetOrderbookParams) (*types.GetOrderbookResp, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/depth",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := s.GenHeaders(spotutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
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

func (s *SpotMarketDataClient) GetRecentTradeList(ctx context.Context, param types.GetTradeParams) (*types.GetTradeResp, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/trades",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := s.GenHeaders(spotutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
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

func (s *SpotMarketDataClient) GetAggTrades(ctx context.Context, param types.GetAggTradesParam) (*types.GetAggTradesResp, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/aggTrades",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := s.GenHeaders(spotutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
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

func (s *SpotMarketDataClient) GetKlines(ctx context.Context, param types.GetKlineParam) (*types.GetKlineResp, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/klines",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := s.GenHeaders(spotutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
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

	data := &types.GetKlineResp{
		Http: resp,
		Body: ret,
	}

	return data, nil
}

func (s *SpotMarketDataClient) GetAvgPrice(ctx context.Context, param types.GetAvgPriceParam) (*types.GetAvgPriceResp, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/avgPrice",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := s.GenHeaders(spotutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.AvgPrice
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetAvgPriceResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (s *SpotMarketDataClient) GetTickerForSymbol(ctx context.Context, param types.GetTickerForSymbolParam) (*types.GetTickerForSymbolResp, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/ticker/24hr",
		Method:  http.MethodGet,
	}

	query := types.TickerParams{
		Symbol: param.Symbol,
	}
	req.Query = query

	headers, err := s.GenHeaders(spotutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body types.Ticker
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetTickerForSymbolResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (s *SpotMarketDataClient) GetTickerForSymbols(ctx context.Context, param types.GetTickerForSymbolsParam) (*types.GetTickerForSymbolsResp, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/ticker/24hr",
		Method:  http.MethodGet,
	}

	query := types.TickerParams{}
	stringsJson, err := json.Marshal(param.Symbols)
	if err != nil {
		return nil, err
	}
	query.Symbols = string(stringsJson)
	req.Query = query

	headers, err := s.GenHeaders(spotutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.Ticker
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetTickerForSymbolsResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}

func (s *SpotMarketDataClient) GetTickerPriceForSymbol(ctx context.Context, param types.GetTickerPriceForSymbolParam) (*types.GetTickerPriceForSymbolResp, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/ticker/price",
		Method:  http.MethodGet,
	}

	query := types.TickerPriceParams{
		Symbol: param.Symbol,
	}
	req.Query = query

	headers, err := s.GenHeaders(spotutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
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

func (s *SpotMarketDataClient) GetTickerPriceForSymbols(ctx context.Context, param types.GetTickerPriceForSymbolsParam) (*types.GetTickerPriceForSymbolsResp, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/ticker/price",
		Method:  http.MethodGet,
	}

	query := types.TickerPriceParams{}
	stringsJson, err := json.Marshal(param.Symbols)
	if err != nil {
		return nil, err
	}
	query.Symbols = string(stringsJson)
	req.Query = query

	headers, err := s.GenHeaders(spotutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
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

func (s *SpotMarketDataClient) GetBookTickerForSymbol(ctx context.Context, param types.GetBookTickerForSymbolParam) (*types.GetBookTickerForSymbolResp, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/ticker/bookTicker",
		Method:  http.MethodGet,
	}

	query := types.BookTickerParams{
		Symbol: param.Symbol,
	}
	req.Query = query

	headers, err := s.GenHeaders(spotutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
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

func (s *SpotMarketDataClient) GetBookTickerForSymbols(ctx context.Context, param types.GetBookTickerForSymbolsParam) (*types.GetBookTickerForSymbolsResp, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/ticker/bookTicker",
		Method:  http.MethodGet,
	}

	query := types.BookTickerParams{}
	stringsJson, err := json.Marshal(param.Symbols)
	if err != nil {
		return nil, err
	}
	query.Symbols = string(stringsJson)
	req.Query = query

	headers, err := s.GenHeaders(spotutils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
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
