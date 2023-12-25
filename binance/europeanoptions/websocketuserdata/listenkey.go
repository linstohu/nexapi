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

package websocketuserdata

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator"
	eoutils "github.com/linstohu/nexapi/binance/europeanoptions/utils"
	usdmutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	"github.com/linstohu/nexapi/utils"
)

type httpAuthClient struct {
	*eoutils.OptionsClient
}

type httpAuthClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func newHttpAuthClient(cfg *httpAuthClientCfg) (*httpAuthClient, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := eoutils.NewOptionsClient(&eoutils.OptionsClientCfg{
		Debug:      cfg.Debug,
		Logger:     cfg.Logger,
		BaseURL:    cfg.BaseURL,
		Key:        cfg.Key,
		Secret:     cfg.Secret,
		RecvWindow: cfg.RecvWindow,
	})
	if err != nil {
		return nil, err
	}

	return &httpAuthClient{
		OptionsClient: cli,
	}, nil
}

type listenKeyResp struct {
	Http *utils.ApiResponse
	Body *listenKeyAPIResp
}

type listenKeyAPIResp struct {
	ListenKey string `json:"listenKey,omitempty"`
}

func (h *httpAuthClient) genListenKey(ctx context.Context) (*listenKeyResp, error) {
	req := utils.HTTPRequest{
		Debug:   h.GetDebug(),
		BaseURL: h.GetBaseURL(),
		Path:    "/eapi/v1/listenKey",
		Method:  http.MethodPost,
	}

	securityType := usdmutils.USER_STREAM

	{
		headers, err := h.GenHeaders(securityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := h.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body listenKeyAPIResp
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &listenKeyResp{
		Http: resp,
		Body: &body,
	}

	return data, nil
}

func (h *httpAuthClient) updateListenKey(ctx context.Context) error {
	req := utils.HTTPRequest{
		Debug:   h.GetDebug(),
		BaseURL: h.GetBaseURL(),
		Path:    "/eapi/v1/listenKey",
		Method:  http.MethodPut,
	}

	securityType := usdmutils.USER_STREAM

	{
		headers, err := h.GenHeaders(securityType)
		if err != nil {
			return err
		}
		req.Headers = headers
	}

	_, err := h.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
