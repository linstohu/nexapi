package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/linstohu/nexapi/deribit/rest/types"
	"github.com/linstohu/nexapi/deribit/rest/types/marketdata"
	"github.com/linstohu/nexapi/deribit/rest/types/supporting"
)

func (d *DeribitRestClient) Test(ctx context.Context) (*supporting.TestResponse, error) {
	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/public/test",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "public/test",
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret supporting.TestResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetBookSummaryByCurrency(ctx context.Context, param marketdata.GetBookSummaryByCurrencyParams) ([]*marketdata.BookSummary, error) {
	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/public/get_book_summary_by_currency",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "public/get_book_summary_by_currency",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret []*marketdata.BookSummary
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (d *DeribitRestClient) GetBookSummaryByInstrument(ctx context.Context, param marketdata.GetBookSummaryByInstrumentParams) ([]*marketdata.BookSummary, error) {
	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/public/get_book_summary_by_instrument",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "public/get_book_summary_by_instrument",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret []*marketdata.BookSummary
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (d *DeribitRestClient) GetContractSize(ctx context.Context, param marketdata.GetContractSizeParams) (*marketdata.GetContractSizeResponse, error) {
	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/public/get_contract_size",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "public/get_contract_size",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret marketdata.GetContractSizeResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetCurrencies(ctx context.Context) ([]*marketdata.Currency, error) {
	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/public/get_currencies",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "public/get_currencies",
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret []*marketdata.Currency
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (d *DeribitRestClient) GetFundingRate(ctx context.Context, param marketdata.GetFundingRateParams) (float64, error) {
	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/public/get_funding_rate_value",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "public/get_funding_rate_value",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return 0, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return 0, err
	}

	var ret float64
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return 0, err
	}

	return ret, nil
}

func (d *DeribitRestClient) GetIndexPrice(ctx context.Context, param marketdata.GetIndexPriceParams) (*marketdata.GetIndexPriceResponse, error) {
	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/public/get_index_price",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "public/get_index_price",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret marketdata.GetIndexPriceResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetInstrument(ctx context.Context, param marketdata.GetInstrumentParams) (*marketdata.Instrument, error) {
	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/public/get_instrument",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "public/get_instrument",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret marketdata.Instrument
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetInstruments(ctx context.Context, param marketdata.GetInstrumentsParams) ([]*marketdata.Instrument, error) {
	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/public/get_instruments",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "public/get_instruments",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret []*marketdata.Instrument
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (d *DeribitRestClient) GetLastTradesByInstrumentAndTime(ctx context.Context, param marketdata.GetLastTradesByInstrumentAndTimeParams) (*marketdata.GetLastTradesResponse, error) {
	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/public/get_last_trades_by_instrument_and_time",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "public/get_last_trades_by_instrument_and_time",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret marketdata.GetLastTradesResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetOrderBook(ctx context.Context, param marketdata.GetOrderBookParams) (*marketdata.GetOrderBookResponse, error) {
	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/public/get_order_book",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "public/get_order_book",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret marketdata.GetOrderBookResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetTradingviewChartData(ctx context.Context, param marketdata.GetTradingviewChartDataParams) (*marketdata.GetTradingviewChartDataResponse, error) {
	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/public/get_tradingview_chart_data",
		Method:  http.MethodGet,
		Headers: DefaultContentType,
		Query:   param,
		Debug:   d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret marketdata.GetTradingviewChartDataResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetTicker(ctx context.Context, param marketdata.TickerParams) (*marketdata.TickerResponse, error) {
	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/public/ticker",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "public/ticker",
			Params:  param,
		},
		Debug: d.debug,
	}

	resp, err := d.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret marketdata.TickerResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
