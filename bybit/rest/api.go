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
	"net/http"

	"github.com/linstohu/nexapi/bybit/rest/types"
	"github.com/linstohu/nexapi/utils"
)

// GetUnifiedAccountBalance(Unified Account API): UNIFIED (trade spot/linear/options)
// Note: the Unified account supports inverse trading. However, the margin used is from the inverse derivatives wallet instead of the unified wallet.
// doc: https://bybit-exchange.github.io/docs/v5/intro#current-api-coverage
func (bb *BybitClient) GetUnifiedAccountBalance() (*types.GetWalletBalanceResp, error) {
	req := utils.HTTPRequest{
		BaseURL: bb.cli.GetBaseURL(),
		Path:    "/v5/account/wallet-balance",
		Method:  http.MethodGet,
		Query: types.GetWalletBalanceParam{
			AccountType: types.UNIFIED,
		},
	}

	headers, err := bb.cli.GenAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := bb.cli.SendHTTPRequest(context.TODO(), req)
	if err != nil {
		return nil, err
	}

	var body types.GetWalletBalanceAPIResp

	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetWalletBalanceResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

// GetContractAccountBalance(Unified Account API): CONTRACT(trade inverse)
func (bb *BybitClient) GetUnifiedAccountContractBalance() (*types.GetWalletBalanceResp, error) {
	req := utils.HTTPRequest{
		BaseURL: bb.cli.GetBaseURL(),
		Path:    "/v5/account/wallet-balance",
		Method:  http.MethodGet,
		Query: types.GetWalletBalanceParam{
			AccountType: types.CONTRACT,
		},
	}

	headers, err := bb.cli.GenAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := bb.cli.SendHTTPRequest(context.TODO(), req)
	if err != nil {
		return nil, err
	}

	var body types.GetWalletBalanceAPIResp

	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetWalletBalanceResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

// GetFundAccountBalance get Funding wallet balance
func (bb *BybitClient) GetFundAccountBalance() (*types.GetAccountBalanceResp, error) {
	req := utils.HTTPRequest{
		BaseURL: bb.cli.GetBaseURL(),
		Path:    "/v5/asset/transfer/query-account-coins-balance",
		Method:  http.MethodGet,
		Query: types.GetAccountBalanceParam{
			AccountType: types.FUND,
		},
	}

	headers, err := bb.cli.GenAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := bb.cli.SendHTTPRequest(context.TODO(), req)
	if err != nil {
		return nil, err
	}

	var body types.GetAccountBalanceAPIResp

	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetAccountBalanceResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}
