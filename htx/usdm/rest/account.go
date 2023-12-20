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

func (ucli *UsdmClient) GetAssetValuation(ctx context.Context, param types.GetAssetValuationParam) (*types.GetAssetValuationResp, error) {
	if err := ucli.cli.CheckAuth(); err != nil {
		return nil, err
	}

	err := ucli.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		BaseURL: ucli.cli.GetBaseURL(),
		Path:    "/linear-swap-api/v1/swap_balance_valuation",
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

	var ret types.GetAssetValuationResp
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}
