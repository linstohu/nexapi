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
	"errors"
	"fmt"
	"net/http"

	"github.com/linstohu/nexapi/htx/spot/rest/types"
	"github.com/linstohu/nexapi/htx/utils"
)

func (scli *SpotClient) NewOrder(ctx context.Context, param types.NewOrderParam) (*types.NewOrderResp, error) {
	if err := scli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := scli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: scli.cli.GetBaseURL(),
		Path:    "/v1/order/orders/place",
		Method:  http.MethodPost,
		Body:    param,
	}

	{
		headers, err := scli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := scli.cli.GenAuthParams()

		signStr, err := scli.cli.NormalizeRequestContent(req, query)
		if err != nil {
			return nil, err
		}

		h := scli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		query.Signature = h

		req.Query = query
	}

	resp, err := scli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.NewOrderResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (scli *SpotClient) CancelOrder(ctx context.Context, param types.CancelOrderParams) (*types.CancelOrderResp, error) {
	if err := scli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := scli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: scli.cli.GetBaseURL(),
		Path:    fmt.Sprintf("/v1/order/orders/%s/submitcancel", param.OrderID),
		Method:  http.MethodPost,
		Body: types.CancelOrderParam{
			Symbol: param.Symbol,
		},
	}

	{
		headers, err := scli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := scli.cli.GenAuthParams()

		signStr, err := scli.cli.NormalizeRequestContent(req, query)
		if err != nil {
			return nil, err
		}

		h := scli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		query.Signature = h

		req.Query = query
	}

	resp, err := scli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.CancelOrderResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (scli *SpotClient) GetOpenOrders(ctx context.Context, param types.GetOpenOrdersParam) (*types.GetOpenOrdersResp, error) {
	if err := scli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := scli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: scli.cli.GetBaseURL(),
		Path:    "/v1/order/openOrders",
		Method:  http.MethodGet,
	}

	{
		headers, err := scli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetOpenOrdersParams{
			GetOpenOrdersParam: param,
			DefaultAuthParam:   scli.cli.GenAuthParams(),
		}

		signStr, err := scli.cli.NormalizeRequestContent(req, query)
		if err != nil {
			return nil, err
		}

		h := scli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		query.Signature = h

		req.Query = query
	}

	resp, err := scli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetOpenOrdersResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (scli *SpotClient) CancelOrders(ctx context.Context, param types.CancelOrdersParam) (*types.CancelOrdersResp, error) {
	if err := scli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := scli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: scli.cli.GetBaseURL(),
		Path:    "/v1/order/orders/batchCancelOpenOrders",
		Method:  http.MethodPost,
		Body:    param,
	}

	{
		headers, err := scli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := scli.cli.GenAuthParams()

		signStr, err := scli.cli.NormalizeRequestContent(req, query)
		if err != nil {
			return nil, err
		}

		h := scli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		query.Signature = h

		req.Query = query
	}

	resp, err := scli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.CancelOrdersResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (scli *SpotClient) SearchMatchResults(ctx context.Context, param types.SearchMatchResultsParam) (*types.SearchMatchResultsResp, error) {
	if err := scli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := scli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: scli.cli.GetBaseURL(),
		Path:    "/v1/order/matchresults",
		Method:  http.MethodGet,
	}

	{
		headers, err := scli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.SearchMatchResultsParams{
			SearchMatchResultsParam: param,
			DefaultAuthParam:        scli.cli.GenAuthParams(),
		}

		signStr, err := scli.cli.NormalizeRequestContent(req, query)
		if err != nil {
			return nil, err
		}

		h := scli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		query.Signature = h

		req.Query = query
	}

	resp, err := scli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.SearchMatchResultsResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}
