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

package utils

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/go-playground/validator"
	"github.com/google/go-querystring/query"
	usdmutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	"github.com/linstohu/nexapi/utils"
)

type OptionsClient struct {
	// debug mode
	debug bool
	// logger
	logger *slog.Logger

	baseURL     string
	key, secret string
	recvWindow  int
}

type OptionsClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string
	Secret     string
	RecvWindow int
}

func NewOptionsClient(cfg *OptionsClientCfg) (*OptionsClient, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := OptionsClient{
		debug:      cfg.Debug,
		logger:     cfg.Logger,
		baseURL:    cfg.BaseURL,
		key:        cfg.Key,
		secret:     cfg.Secret,
		recvWindow: cfg.RecvWindow,
	}

	if cfg.RecvWindow == 0 {
		cli.recvWindow = 5000
	}

	if cli.logger == nil {
		cli.logger = slog.Default()
	}

	return &cli, nil
}

func (o *OptionsClient) GetDebug() bool {
	return o.debug
}

func (o *OptionsClient) GetBaseURL() string {
	return o.baseURL
}

func (o *OptionsClient) GetKey() string {
	return o.key
}

func (o *OptionsClient) GetSecret() string {
	return o.secret
}

func (o *OptionsClient) GetRecvWindow() int {
	return o.recvWindow
}

func (o *OptionsClient) GenHeaders(t usdmutils.SecurityType) (map[string]string, error) {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Accept":       "application/json",
	}

	// SecurityType each endpoint has a security type that determines how you will interact with it
	// docs: https://binance-docs.github.io/apidocs/voptions/en/#endpoint-security-type
	switch t {
	case usdmutils.TRADE, usdmutils.USER_DATA, usdmutils.USER_STREAM, usdmutils.MARKET_DATA:
		key := o.GetKey()
		if key == "" {
			return nil, fmt.Errorf("a valid API-Key required")
		}

		headers["X-MBX-APIKEY"] = o.GetKey()
	}

	return headers, nil
}

// For GET endpoints, parameters must be sent as a query string without setting content type in the http headers.
// doc: https://binance-docs.github.io/apidocs/voptions/en/#general-api-information
func (o *OptionsClient) GenGetHeaders(t usdmutils.SecurityType) (map[string]string, error) {
	headers := map[string]string{
		"Accept": "application/json",
	}

	// SecurityType each endpoint has a security type that determines how you will interact with it
	// docs: https://binance-docs.github.io/apidocs/voptions/en/#endpoint-security-type
	switch t {
	case usdmutils.TRADE, usdmutils.USER_DATA, usdmutils.USER_STREAM, usdmutils.MARKET_DATA:
		key := o.GetKey()
		if key == "" {
			return nil, fmt.Errorf("a valid API-Key required")
		}

		headers["X-MBX-APIKEY"] = o.GetKey()
	}

	return headers, nil
}

func (o *OptionsClient) NeedSignature(t usdmutils.SecurityType) bool {
	switch t {
	case usdmutils.TRADE, usdmutils.USER_DATA:
		// TRADE and USER_DATA endpoints are SIGNED endpoints.
		return true
	default:
		return false
	}
}

func (o *OptionsClient) SendHTTPRequest(ctx context.Context, req utils.HTTPRequest) (*utils.ApiResponse, error) {
	client := http.Client{}

	var body io.Reader
	if req.Body != nil {
		formData, err := query.Values(req.Body)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(formData.Encode())
	}

	url, err := url.Parse(req.BaseURL + req.Path)
	if err != nil {
		return nil, err
	}

	if req.Query != nil {
		q, err := query.Values(req.Query)
		if err != nil {
			return nil, err
		}
		url.RawQuery = q.Encode()
	}

	request, err := http.NewRequestWithContext(ctx, req.Method, url.String(), body)
	if err != nil {
		return nil, err
	}

	for k, v := range req.Headers {
		request.Header.Set(k, v)
	}

	if o.GetDebug() {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		o.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if o.GetDebug() {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		o.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	return utils.NewApiResponse(&req, resp), nil
}
