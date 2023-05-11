package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/linstohu/nexapi/woox/api/types"
)

func (w *WooXClient) SendOrder(ctx context.Context, params types.SendOrderReq) (*types.SendOrderResp, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	path := "/v1/order"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodPost,
		Body:   params,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.SendOrderResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) GetPrivateBalances(ctx context.Context) (*types.Balance, error) {
	path := "/v3/balances"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodGet,
		Debug:  w.debug,
	}

	headers, err := w.GenV3APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Balance
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) GetAssetHisotry(ctx context.Context, params types.GetAssetHisotryParam) (*types.AssetHisotryResp, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	path := "/v1/asset/history"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodGet,
		Query:  params,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.AssetHisotryResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (w *WooXClient) GetIPRestriction(ctx context.Context) (*types.IPRestriction, error) {
	path := "/v1/sub_account/ip_restriction"

	req := types.HTTPRequest{
		URL:    w.basePath + path,
		Path:   path,
		Method: http.MethodGet,
		Debug:  w.debug,
	}

	headers, err := w.GenV1APIAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.IPRestriction
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
