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

package spotutils

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
	"github.com/linstohu/nexapi/utils"
)

type SpotClient struct {
	// debug mode
	debug bool
	// logger
	logger *slog.Logger

	baseURL     string
	key, secret string
	recvWindow  int
}

type SpotClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string
	Secret     string
	RecvWindow int
}

func NewSpotClient(cfg *SpotClientCfg) (*SpotClient, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := SpotClient{
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

func (s *SpotClient) GetDebug() bool {
	return s.debug
}

func (s *SpotClient) GetBaseURL() string {
	return s.baseURL
}

func (s *SpotClient) GetKey() string {
	return s.key
}

func (s *SpotClient) GetSecret() string {
	return s.secret
}

func (s *SpotClient) GetRecvWindow() int {
	return s.recvWindow
}

func (s *SpotClient) GenHeaders(t SecurityType) (map[string]string, error) {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Accept":       "application/json",
	}

	// SecurityType each endpoint has a security type that determines how you will interact with it
	// docs: https://binance-docs.github.io/apidocs/spot/en/#endpoint-security-type
	switch t {
	case TRADE, MARGIN, USER_DATA, USER_STREAM, MARKET_DATA:
		key := s.GetKey()
		if key == "" {
			return nil, fmt.Errorf("valid API-Key required")
		}

		headers["X-MBX-APIKEY"] = s.GetKey()
	}

	return headers, nil
}

func (s *SpotClient) NeedSignature(t SecurityType) bool {
	switch t {
	case TRADE, MARGIN, USER_DATA:
		return true
	default:
		return false
	}
}

func (s *SpotClient) SendHTTPRequest(ctx context.Context, req utils.HTTPRequest) (*utils.ApiResponse, error) {
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

	if s.GetDebug() {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}

		s.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if s.GetDebug() {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		s.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	return utils.NewApiResponse(&req, resp), nil
}
