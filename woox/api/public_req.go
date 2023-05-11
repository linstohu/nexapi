package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/linstohu/nexapi/woox/api/types"
)

func (w *WooXClient) GetPublicInfo(ctx context.Context) (*types.AvailableSymbols, error) {
	req := types.HTTPRequest{
		URL:     w.basePath + "/v1/public/info",
		Method:  http.MethodGet,
		Headers: DefaultContentType,
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

func (w *WooXClient) GetPublicInfoForSymbol(ctx context.Context, symbol string) (*types.SymbolInfo, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol must be given by api [/v1/public/info/:symbol]")
	}

	req := types.HTTPRequest{
		URL:     fmt.Sprintf("%s%s%s", w.basePath, "/v1/public/info/", symbol),
		Method:  http.MethodGet,
		Headers: DefaultContentType,
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

func (w *WooXClient) GetPublicMarketTrades(ctx context.Context, params *types.GetMarketTradesParam) (*types.MarketTrade, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     w.basePath + "/v1/public/market_trades",
		Method:  http.MethodGet,
		Headers: DefaultContentType,
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

func (w *WooXClient) GetPublicOrderbook(ctx context.Context, symbol string, params *types.GetOrderbookParam) (*types.Orderbook, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol must be given by api [/v1/public/orderbook/:symbol]")
	}

	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     fmt.Sprintf("%s%s%s", w.basePath, "/v1/public/orderbook/", symbol),
		Method:  http.MethodGet,
		Headers: DefaultContentType,
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

func (w *WooXClient) GetPublicKline(ctx context.Context, params *types.GetKlineParam) (*types.Kline, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     w.basePath + "/v1/public/kline",
		Method:  http.MethodGet,
		Headers: DefaultContentType,
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

func (w *WooXClient) GetPublicTokens(ctx context.Context) (*types.Tokens, error) {
	req := types.HTTPRequest{
		URL:     w.basePath + "/v1/public/token",
		Method:  http.MethodGet,
		Headers: DefaultContentType,
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

func (w *WooXClient) GetPublicFundingRates(ctx context.Context) (*types.FundingRates, error) {
	req := types.HTTPRequest{
		URL:     w.basePath + "/v1/public/funding_rates",
		Method:  http.MethodGet,
		Headers: DefaultContentType,
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

func (w *WooXClient) GetPublicFundingRateForSymbol(ctx context.Context, symbol string) (*types.FundingRate, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol must be given by api [/v1/public/funding_rate/:symbol]")
	}

	req := types.HTTPRequest{
		URL:     fmt.Sprintf("%s%s%s", w.basePath, "/v1/public/funding_rate/", symbol),
		Method:  http.MethodGet,
		Headers: DefaultContentType,
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

func (w *WooXClient) GetPublicAllFuturesInfo(ctx context.Context) (*types.AllFuturesInfo, error) {
	req := types.HTTPRequest{
		URL:     w.basePath + "/v1/public/futures",
		Method:  http.MethodGet,
		Headers: DefaultContentType,
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

func (w *WooXClient) GetPublicFuturesInfoForSymbol(ctx context.Context, symbol string) (*types.OneFuturesInfo, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol must be given by api [/v1/public/futures/:symbol]")
	}

	req := types.HTTPRequest{
		URL:     fmt.Sprintf("%s%s%s", w.basePath, "/v1/public/futures/", symbol),
		Method:  http.MethodGet,
		Headers: DefaultContentType,
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
