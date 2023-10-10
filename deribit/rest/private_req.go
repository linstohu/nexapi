package rest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/linstohu/nexapi/deribit/rest/types"
	"github.com/linstohu/nexapi/deribit/rest/types/account"
	"github.com/linstohu/nexapi/deribit/rest/types/auth"
)

func (d *DeribitRestClient) Auth(ctx context.Context, param auth.AuthParams) (*auth.AuthResponse, error) {
	err := d.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/public/auth",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "public/auth",
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

	var ret auth.AuthResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

var ErrAuth = errors.New("auth error, you should reinitialize client using key and secret")

func (d *DeribitRestClient) GetAccountSummary(ctx context.Context, param account.GetAccountSummaryParams) (*account.AccountSummary, error) {
	if d.auth.token == "" {
		return nil, ErrAuth
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/get_account_summary",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/get_account_summary",
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

	var ret account.AccountSummary
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (d *DeribitRestClient) GetPositions(ctx context.Context, param account.GetPositionsParams) ([]*account.Position, error) {
	if d.auth.token == "" {
		return nil, ErrAuth
	}

	req := types.HTTPRequest{
		URL:     d.baseURL + "/api/v2/private/get_positions",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "private/get_positions",
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

	var ret []*account.Position
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
