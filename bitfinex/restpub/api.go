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

package bitfinexrestpub

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
)

var (
	BaseURL = "https://api-pub.bitfinex.com"
)

type RestPubClient struct {
	*BitfinexClient

	// validate struct fields
	validate *validator.Validate
}

func NewRestPubClient(cfg *BitfinexClientCfg) (*RestPubClient, error) {
	cli, err := NewBitfinexClient(cfg)
	if err != nil {
		return nil, err
	}

	validator := validator.New()

	return &RestPubClient{
		BitfinexClient: cli,
		validate:       validator,
	}, nil
}

func (r *RestPubClient) PlatformStatus(ctx context.Context) error {
	req := HTTPRequest{
		BaseURL: r.GetBaseURL(),
		Path:    "/v2/platform/status",
		Method:  http.MethodGet,
		Headers: make(map[string]string),
	}

	_, err := r.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
