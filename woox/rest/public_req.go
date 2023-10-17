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
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/linstohu/nexapi/woox/rest/types"
)

func (w *WooXRestClient) GetPublicInfo(ctx context.Context) (*types.AvailableSymbols, error) {
	req := types.HTTPRequest{
		URL:     w.baseURL + "/v1/public/info",
		Method:  http.MethodGet,
		Headers: V1DefaultContentType,
		Debug:   w.debug,
	}

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.AvailableSymbols
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXRestClient) GetPublicInfoForSymbol(ctx context.Context, symbol string) (*types.SymbolInfo, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol must be given by api [/v1/public/info/:symbol]")
	}

	req := types.HTTPRequest{
		URL:     fmt.Sprintf("%s%s%s", w.baseURL, "/v1/public/info/", symbol),
		Method:  http.MethodGet,
		Headers: V1DefaultContentType,
		Debug:   w.debug,
	}

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.SymbolInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXRestClient) GetPublicMarketTrades(ctx context.Context, params *types.GetMarketTradesParam) (*types.MarketTrade, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     w.baseURL + "/v1/public/market_trades",
		Method:  http.MethodGet,
		Headers: V1DefaultContentType,
		Query:   params,
		Debug:   w.debug,
	}

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.MarketTrade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXRestClient) GetPublicOrderbook(ctx context.Context, symbol string, params *types.GetOrderbookParam) (*types.Orderbook, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol must be given by api [/v1/public/orderbook/:symbol]")
	}

	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     fmt.Sprintf("%s%s%s", w.baseURL, "/v1/public/orderbook/", symbol),
		Method:  http.MethodGet,
		Headers: V1DefaultContentType,
		Query:   params,
		Debug:   w.debug,
	}

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Orderbook
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXRestClient) GetPublicKline(ctx context.Context, params *types.GetKlineParam) (*types.Kline, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     w.baseURL + "/v1/public/kline",
		Method:  http.MethodGet,
		Headers: V1DefaultContentType,
		Query:   params,
		Debug:   w.debug,
	}

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Kline
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXRestClient) GetPublicTokens(ctx context.Context) (*types.Tokens, error) {
	req := types.HTTPRequest{
		URL:     w.baseURL + "/v1/public/token",
		Method:  http.MethodGet,
		Headers: V1DefaultContentType,
		Debug:   w.debug,
	}

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Tokens
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXRestClient) GetPublicFundingRates(ctx context.Context) (*types.FundingRates, error) {
	req := types.HTTPRequest{
		URL:     w.baseURL + "/v1/public/funding_rates",
		Method:  http.MethodGet,
		Headers: V1DefaultContentType,
		Debug:   w.debug,
	}

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.FundingRates
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXRestClient) GetPublicFundingRateForSymbol(ctx context.Context, symbol string) (*types.FundingRate, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol must be given by api [/v1/public/funding_rate/:symbol]")
	}

	req := types.HTTPRequest{
		URL:     fmt.Sprintf("%s%s%s", w.baseURL, "/v1/public/funding_rate/", symbol),
		Method:  http.MethodGet,
		Headers: V1DefaultContentType,
		Debug:   w.debug,
	}

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.FundingRate
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXRestClient) GetPublicAllFuturesInfo(ctx context.Context) (*types.AllFuturesInfo, error) {
	req := types.HTTPRequest{
		URL:     w.baseURL + "/v1/public/futures",
		Method:  http.MethodGet,
		Headers: V1DefaultContentType,
		Debug:   w.debug,
	}

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.AllFuturesInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXRestClient) GetPublicFuturesInfoForSymbol(ctx context.Context, symbol string) (*types.OneFuturesInfo, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol must be given by api [/v1/public/futures/:symbol]")
	}

	req := types.HTTPRequest{
		URL:     fmt.Sprintf("%s%s%s", w.baseURL, "/v1/public/futures/", symbol),
		Method:  http.MethodGet,
		Headers: V1DefaultContentType,
		Debug:   w.debug,
	}

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.OneFuturesInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
