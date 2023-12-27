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
	"net/http"

	"github.com/linstohu/nexapi/htx/usdm/rest/types"
	"github.com/linstohu/nexapi/htx/utils"
)

func (ucli *UsdmClient) PlaceIsolatedOrder(ctx context.Context, param types.PlaceIsolatedOrderParam) (*types.PlaceOrderResp, error) {
	if err := ucli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := ucli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-api/v1/swap_order",
		Method:  http.MethodPost,
		Body:    param,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := ucli.cli.GenAuthParams()

		signStr, err := ucli.cli.NormalizeRequestContent(req, query)
		if err != nil {
			return nil, err
		}

		h := ucli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		query.Signature = h

		req.Query = query
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.PlaceOrderResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (ucli *UsdmClient) PlaceCrossOrder(ctx context.Context, param types.PlaceCrossOrderParam) (*types.PlaceOrderResp, error) {
	if err := ucli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := ucli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-api/v1/swap_cross_order",
		Method:  http.MethodPost,
		Body:    param,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := ucli.cli.GenAuthParams()

		signStr, err := ucli.cli.NormalizeRequestContent(req, query)
		if err != nil {
			return nil, err
		}

		h := ucli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		query.Signature = h

		req.Query = query
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.PlaceOrderResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (ucli *UsdmClient) CancelIsolatedOrder(ctx context.Context, param types.CancelIsolatedOrderParam) (*types.CancelOrderResp, error) {
	if err := ucli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := ucli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-api/v1/swap_cancel",
		Method:  http.MethodPost,
		Body:    param,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := ucli.cli.GenAuthParams()

		signStr, err := ucli.cli.NormalizeRequestContent(req, query)
		if err != nil {
			return nil, err
		}

		h := ucli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		query.Signature = h

		req.Query = query
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.CancelOrderResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (ucli *UsdmClient) CancelCrossOrder(ctx context.Context, param types.CancelCrossOrderParam) (*types.CancelOrderResp, error) {
	if err := ucli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := ucli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-api/v1/swap_cross_cancel",
		Method:  http.MethodPost,
		Body:    param,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := ucli.cli.GenAuthParams()

		signStr, err := ucli.cli.NormalizeRequestContent(req, query)
		if err != nil {
			return nil, err
		}

		h := ucli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		query.Signature = h

		req.Query = query
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.CancelOrderResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (ucli *UsdmClient) CancelAllIsolatedOrders(ctx context.Context, param types.CancelAllIsolatedOrdersParam) (*types.CancelOrderResp, error) {
	if err := ucli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := ucli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-api/v1/swap_cancelall",
		Method:  http.MethodPost,
		Body:    param,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := ucli.cli.GenAuthParams()

		signStr, err := ucli.cli.NormalizeRequestContent(req, query)
		if err != nil {
			return nil, err
		}

		h := ucli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		query.Signature = h

		req.Query = query
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.CancelOrderResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (ucli *UsdmClient) CancelAllCrossOrders(ctx context.Context, param types.CancelAllCrossOrdersParam) (*types.CancelOrderResp, error) {
	if err := ucli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := ucli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-api/v1/swap_cross_cancelall",
		Method:  http.MethodPost,
		Body:    param,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := ucli.cli.GenAuthParams()

		signStr, err := ucli.cli.NormalizeRequestContent(req, query)
		if err != nil {
			return nil, err
		}

		h := ucli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		query.Signature = h

		req.Query = query
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.CancelOrderResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (ucli *UsdmClient) GetIsolatedOpenOrders(ctx context.Context, param types.GetIsolatedOpenOrdersParam) (*types.GetIsolatedOpenOrdersResp, error) {
	if err := ucli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := ucli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-api/v1/swap_openorders",
		Method:  http.MethodPost,
		Body:    param,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := ucli.cli.GenAuthParams()

		signStr, err := ucli.cli.NormalizeRequestContent(req, query)
		if err != nil {
			return nil, err
		}

		h := ucli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		query.Signature = h

		req.Query = query
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetIsolatedOpenOrdersResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (ucli *UsdmClient) GetCrossOpenOrders(ctx context.Context, param types.GetCrossOpenOrdersParam) (*types.GetCrossOpenOrdersResp, error) {
	if err := ucli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := ucli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-api/v1/swap_cross_openorders",
		Method:  http.MethodPost,
		Body:    param,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := ucli.cli.GenAuthParams()

		signStr, err := ucli.cli.NormalizeRequestContent(req, query)
		if err != nil {
			return nil, err
		}

		h := ucli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		query.Signature = h

		req.Query = query
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetCrossOpenOrdersResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (ucli *UsdmClient) GetIsolatedHistoryMatchResults(ctx context.Context, param types.GetIsolatedHistoryMatchResultsParam) (*types.HistoryMatchResultsResp, error) {
	if err := ucli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := ucli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-api/v3/swap_matchresults",
		Method:  http.MethodPost,
		Body:    param,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := ucli.cli.GenAuthParams()

		signStr, err := ucli.cli.NormalizeRequestContent(req, query)
		if err != nil {
			return nil, err
		}

		h := ucli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		query.Signature = h

		req.Query = query
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.HistoryMatchResultsResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}

func (ucli *UsdmClient) GetCrossHistoryMatchResults(ctx context.Context, param types.GetCrossHistoryMatchResultsParam) (*types.HistoryMatchResultsResp, error) {
	if err := ucli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := ucli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-api/v3/swap_cross_matchresults",
		Method:  http.MethodPost,
		Body:    param,
	}

	{
		headers, err := ucli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := ucli.cli.GenAuthParams()

		signStr, err := ucli.cli.NormalizeRequestContent(req, query)
		if err != nil {
			return nil, err
		}

		h := ucli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		query.Signature = h

		req.Query = query
	}

	resp, err := ucli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.HistoryMatchResultsResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}
