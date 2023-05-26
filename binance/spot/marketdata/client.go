package marketdata

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/spot/marketdata/types"
	spotutils "github.com/linstohu/nexapi/binance/spot/utils"
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
	req := spotutils.HTTPRequest{
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

func (s *SpotMarketDataClient) GetServerTime(ctx context.Context) (*types.ServerTime, error) {
	req := spotutils.HTTPRequest{
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

	var ret types.ServerTime
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SpotMarketDataClient) GetExchangeInfo(ctx context.Context, param types.GetExchangeInfoParam) (*types.ExchangeInfo, error) {
	req := spotutils.HTTPRequest{
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

	var ret types.ExchangeInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SpotMarketDataClient) GetOrderbook(ctx context.Context, param types.GetOrderbookParams) (*types.Orderbook, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := spotutils.HTTPRequest{
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

	var ret types.Orderbook
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SpotMarketDataClient) GetRecentTradeList(ctx context.Context, param types.GetTradeParams) ([]*types.Trade, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := spotutils.HTTPRequest{
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

	var ret []*types.Trade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *SpotMarketDataClient) GetAggTrades(ctx context.Context, param types.GetAggTradesParam) ([]*types.AggTrade, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := spotutils.HTTPRequest{
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

	var ret []*types.AggTrade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *SpotMarketDataClient) GetKlines(ctx context.Context, param types.GetKlineParam) ([]*types.Kline, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := spotutils.HTTPRequest{
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

func (s *SpotMarketDataClient) GetAvgPrice(ctx context.Context, param types.GetAvgPriceParam) (*types.AvgPrice, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := spotutils.HTTPRequest{
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

	var ret types.AvgPrice
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SpotMarketDataClient) GetTickerForSymbol(ctx context.Context, param types.GetTickerForSymbolParam) (*types.Ticker, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := spotutils.HTTPRequest{
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

	var ret types.Ticker
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SpotMarketDataClient) GetTickerForSymbols(ctx context.Context, param types.GetTickerForSymbolsParam) ([]*types.Ticker, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := spotutils.HTTPRequest{
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

	var ret []*types.Ticker
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *SpotMarketDataClient) GetTickerPriceForSymbol(ctx context.Context, param types.GetTickerPriceForSymbolParam) (*types.TickerPrice, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := spotutils.HTTPRequest{
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

	var ret types.TickerPrice
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SpotMarketDataClient) GetTickerPriceForSymbols(ctx context.Context, param types.GetTickerPriceForSymbolsParam) ([]*types.TickerPrice, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := spotutils.HTTPRequest{
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

	var ret []*types.TickerPrice
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *SpotMarketDataClient) GetBookTickerForSymbol(ctx context.Context, param types.GetBookTickerForSymbolParam) (*types.BookTicker, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := spotutils.HTTPRequest{
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

	var ret types.BookTicker
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SpotMarketDataClient) GetBookTickerForSymbols(ctx context.Context, param types.GetBookTickerForSymbolsParam) ([]*types.BookTicker, error) {
	err := s.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := spotutils.HTTPRequest{
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

	var ret []*types.BookTicker
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
