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

package restauth

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/utils"
)

var (
	BaseURL = "https://api.bitfinex.com"
)

type RestAuthClient struct {
	*BitfinexAuthClient

	// validate struct fields
	validate *validator.Validate
}

func NewRestAuthClient(cfg *BitfinexClientCfg) (*RestAuthClient, error) {
	cli, err := NewBitfinexClient(cfg)
	if err != nil {
		return nil, err
	}

	validator := validator.New()

	return &RestAuthClient{
		BitfinexAuthClient: cli,
		validate:           validator,
	}, nil
}

func (r *RestAuthClient) GetWallets(ctx context.Context) error {
	req := utils.HTTPRequest{
		Debug:   r.GetDebug(),
		BaseURL: r.GetBaseURL(),
		Path:    "/v2/auth/r/wallets",
		Method:  http.MethodPost,
	}

	{
		headers, err := r.GenHeaders(req)
		if err != nil {
			return err
		}
		req.Headers = headers
	}

	_, err := r.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
