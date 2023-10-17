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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/go-querystring/query"
	"github.com/linstohu/nexapi/deribit/rest/types"
	"github.com/linstohu/nexapi/deribit/rest/types/auth"
)

type DeribitRestClient struct {
	baseURL     string
	key, secret string
	// debug mode
	debug bool
	// logger
	logger *slog.Logger
	// validate struct fields
	validate *validator.Validate

	auth struct {
		token     string
		expiresAt int64
	}
}

type DeribitRestClientCfg struct {
	BaseURL string `validate:"required"`
	Key     string
	Secret  string
	Debug   bool
	// Logger
	Logger *slog.Logger
}

func NewDeribitRestClient(cfg *DeribitRestClientCfg) (*DeribitRestClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := DeribitRestClient{
		baseURL: cfg.BaseURL,
		key:     cfg.Key,
		secret:  cfg.Secret,
		debug:   cfg.Debug,
		logger:  cfg.Logger,

		validate: validator,
	}

	if cli.logger == nil {
		cli.logger = slog.Default()
	}

	if cfg.Key != "" && cfg.Secret != "" {
		token, err := cli.Auth(context.TODO(), auth.AuthParams{
			GrantType:    "client_credentials",
			ClientID:     cfg.Key,
			ClientSecret: cfg.Secret,
		})
		if err != nil {
			return nil, fmt.Errorf("init private rest client failed, error: %v", err)
		}

		now := time.Now().Unix()
		cli.auth.token = token.AccessToken
		cli.auth.expiresAt = now + token.ExpiresIn - 5
	}

	return &cli, nil
}

func (d *DeribitRestClient) SendHTTPRequest(ctx context.Context, req types.HTTPRequest) ([]byte, error) {
	client := http.Client{}

	var body io.Reader
	if req.Body != nil {
		jsonBody, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(jsonBody)
	}

	url, err := url.Parse(req.URL)
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

	if req.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		d.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if req.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		d.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned a non-200 status code: [%d] - [%s]", resp.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}
