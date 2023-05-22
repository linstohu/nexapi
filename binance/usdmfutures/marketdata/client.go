package marketdata

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	spottypes "github.com/linstohu/nexapi/binance/spot/marketdata/types"
	"github.com/linstohu/nexapi/binance/usdmfutures/marketdata/types"
	"github.com/linstohu/nexapi/binance/usdmfutures/utils"
	"github.com/valyala/fastjson"
)

type USDMFuturesMarketDataClient struct {
	*utils.USDMarginedClient

	// validate struct fields
	validate *validator.Validate
}

func NewUSDMFuturesMarketDataClient(cfg *utils.USDMarginedClientCfg) (*USDMFuturesMarketDataClient, error) {
	cli, err := utils.NewUSDMarginedClient(cfg)
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
		SecurityType: utils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/ping",
		Method:       http.MethodGet,
	}

	headers, err := u.GenHeaders(req.SecurityType)
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

func (u *USDMFuturesMarketDataClient) GetServerTime(ctx context.Context) (*spottypes.ServerTime, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/time",
		Method:       http.MethodGet,
	}

	headers, err := u.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret spottypes.ServerTime
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (u *USDMFuturesMarketDataClient) GetExchangeInfo(ctx context.Context) (*types.ExchangeInfo, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/exchangeInfo",
		Method:       http.MethodGet,
	}

	headers, err := u.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.ExchangeInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (u *USDMFuturesMarketDataClient) GetOrderbook(ctx context.Context, param types.GetOrderbookParams) (*types.Orderbook, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: utils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/depth",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := u.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Orderbook
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (u *USDMFuturesMarketDataClient) GetRecentTradeList(ctx context.Context, param types.GetTradeParams) ([]*types.Trade, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: utils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/trades",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := u.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.Trade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *USDMFuturesMarketDataClient) GetAggTrades(ctx context.Context, param types.GetAggTradesParam) ([]*types.AggTrade, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: utils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/aggTrades",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := u.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.AggTrade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *USDMFuturesMarketDataClient) GetKlines(ctx context.Context, param types.GetKlineParam) ([]*spottypes.Kline, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: utils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/klines",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := u.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var p fastjson.Parser
	js, err := p.ParseBytes(resp)
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

	return ret, nil
}

func (u *USDMFuturesMarketDataClient) GetMarkPriceForSymbol(ctx context.Context, param types.GetMarkPriceParam) (*types.MarkPrice, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: utils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/premiumIndex",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := u.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret *types.MarkPrice
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *USDMFuturesMarketDataClient) GetMarkPriceForAllSymbols(ctx context.Context) ([]*types.MarkPrice, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/premiumIndex",
		Method:       http.MethodGet,
	}

	headers, err := u.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.MarkPrice
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *USDMFuturesMarketDataClient) GetFundingRateHistory(ctx context.Context, param types.GetFundingRateParam) ([]*types.FundingRate, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: utils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/fundingRate",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := u.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.FundingRate
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *USDMFuturesMarketDataClient) GetTickerPriceForSymbol(ctx context.Context, param types.GetTickerPriceParam) (*types.TickerPrice, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: utils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/ticker/price",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := u.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret *types.TickerPrice
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *USDMFuturesMarketDataClient) GetTickerPriceForAllSymbols(ctx context.Context) ([]*types.TickerPrice, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/ticker/price",
		Method:       http.MethodGet,
	}

	headers, err := u.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.TickerPrice
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *USDMFuturesMarketDataClient) GetBookTickerForSymbol(ctx context.Context, param types.GetBookTickerForSymbolParam) (*types.BookTicker, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: utils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/ticker/bookTicker",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := u.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret *types.BookTicker
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *USDMFuturesMarketDataClient) GetBookTickerForAllSymbols(ctx context.Context) ([]*types.BookTicker, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/fapi/v1/ticker/bookTicker",
		Method:       http.MethodGet,
	}

	headers, err := u.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.BookTicker
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *USDMFuturesMarketDataClient) GetOpenInterestHistory(ctx context.Context, param types.GetOpenInterestHistParam) ([]*types.OpenInterestHist, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: utils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/futures/data/openInterestHist",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := u.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := u.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.OpenInterestHist
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
