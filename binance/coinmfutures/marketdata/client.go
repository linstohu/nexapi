package marketdata

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/coinmfutures/marketdata/types"
	"github.com/linstohu/nexapi/binance/coinmfutures/utils"
	spottypes "github.com/linstohu/nexapi/binance/spot/marketdata/types"
	usdmtypes "github.com/linstohu/nexapi/binance/usdmfutures/marketdata/types"
	usdmutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	"github.com/valyala/fastjson"
)

type CoinMFuturesMarketDataClient struct {
	*utils.CoinMarginedClient

	// validate struct fields
	validate *validator.Validate
}

func NewCoinMFuturesMarketDataClient(cfg *utils.CoinMarginedClientCfg) (*CoinMFuturesMarketDataClient, error) {
	cli, err := utils.NewCoinMarginedClient(cfg)
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
		SecurityType: usdmutils.NONE,
		BaseURL:      c.GetBaseURL(),
		Path:         "/dapi/v1/ping",
		Method:       http.MethodGet,
	}

	headers, err := c.GenHeaders(req.SecurityType)
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

func (c *CoinMFuturesMarketDataClient) GetServerTime(ctx context.Context) (*spottypes.ServerTime, error) {
	req := utils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      c.GetBaseURL(),
		Path:         "/dapi/v1/time",
		Method:       http.MethodGet,
	}

	headers, err := c.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret spottypes.ServerTime
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *CoinMFuturesMarketDataClient) GetExchangeInfo(ctx context.Context) (*types.ExchangeInfo, error) {
	req := utils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      c.GetBaseURL(),
		Path:         "/dapi/v1/exchangeInfo",
		Method:       http.MethodGet,
	}

	headers, err := c.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.ExchangeInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *CoinMFuturesMarketDataClient) GetOrderbook(ctx context.Context, param types.GetOrderbookParams) (*types.Orderbook, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      c.GetBaseURL(),
		Path:         "/dapi/v1/depth",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := c.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Orderbook
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *CoinMFuturesMarketDataClient) GetRecentTradeList(ctx context.Context, param types.GetTradeParams) ([]*types.Trade, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      c.GetBaseURL(),
		Path:         "/dapi/v1/trades",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := c.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.Trade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *CoinMFuturesMarketDataClient) GetAggTrades(ctx context.Context, param usdmtypes.GetAggTradesParam) ([]*usdmtypes.AggTrade, error) {
	err := u.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      u.GetBaseURL(),
		Path:         "/dapi/v1/aggTrades",
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

	var ret []*usdmtypes.AggTrade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *CoinMFuturesMarketDataClient) GetMarkPrice(ctx context.Context, param types.GetMarkPriceParam) ([]*types.MarkPrice, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      c.GetBaseURL(),
		Path:         "/dapi/v1/premiumIndex",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := c.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.MarkPrice
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *CoinMFuturesMarketDataClient) GetFundingRateHistory(ctx context.Context, param types.GetFundingRateParam) ([]*types.FundingRate, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      c.GetBaseURL(),
		Path:         "/dapi/v1/fundingRate",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := c.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.FundingRate
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *CoinMFuturesMarketDataClient) GetKlines(ctx context.Context, param usdmtypes.GetKlineParam) ([]*types.Kline, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      c.GetBaseURL(),
		Path:         "/dapi/v1/klines",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := c.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
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

	return ret, nil
}

func (c *CoinMFuturesMarketDataClient) GetTickerPrice(ctx context.Context, param types.GetPriceTickerParam) ([]*types.PriceTicker, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      c.GetBaseURL(),
		Path:         "/dapi/v1/ticker/price",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := c.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.PriceTicker
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *CoinMFuturesMarketDataClient) GetBookTicker(ctx context.Context, param types.GetBookTickerParam) ([]*types.BookTicker, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      c.GetBaseURL(),
		Path:         "/dapi/v1/ticker/bookTicker",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := c.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.BookTicker
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *CoinMFuturesMarketDataClient) GetOpenInterest(ctx context.Context, param types.GetOpenInterestParam) (*types.OpenInterest, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      c.GetBaseURL(),
		Path:         "/dapi/v1/openInterest",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := c.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.OpenInterest
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *CoinMFuturesMarketDataClient) GetOpenInterestHistory(ctx context.Context, param types.GetOpenInterestHistParam) ([]*types.OpenInterestHist, error) {
	err := c.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		SecurityType: usdmutils.NONE,
		BaseURL:      c.GetBaseURL(),
		Path:         "/futures/data/openInterestHist",
		Method:       http.MethodGet,
		Query:        param,
	}

	headers, err := c.GenHeaders(req.SecurityType)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.OpenInterestHist
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
