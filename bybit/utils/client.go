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
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	goquery "github.com/google/go-querystring/query"
	"github.com/linstohu/nexapi/utils"
)

type BybitClient struct {
	// debug mode
	debug bool
	// logger
	logger *slog.Logger

	baseURL     string
	key, secret string
	recvWindow  int
}

type BybitClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string
	Secret     string
	RecvWindow int
}

func NewBybitClient(cfg *BybitClientCfg) (*BybitClient, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := BybitClient{
		debug:   cfg.Debug,
		logger:  cfg.Logger,
		baseURL: cfg.BaseURL,
		key:     cfg.Key,
		secret:  cfg.Secret,
	}

	if cfg.RecvWindow == 0 {
		cli.recvWindow = 5000
	}

	if cli.logger == nil {
		cli.logger = slog.Default()
	}

	return &cli, nil
}

func (bb *BybitClient) GetDebug() bool {
	return bb.debug
}

func (bb *BybitClient) GetBaseURL() string {
	return bb.baseURL
}

func (bb *BybitClient) GetKey() string {
	return bb.key
}

func (bb *BybitClient) GetSecret() string {
	return bb.secret
}

func (bb *BybitClient) GetRecvWindow() int {
	return bb.recvWindow
}

func (bb *BybitClient) GenAuthHeaders(req utils.HTTPRequest) (map[string]string, error) {
	if bb.GetKey() == "" || bb.GetSecret() == "" {
		return nil, fmt.Errorf("key and secret needed when auth headers")
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	strBody := ""
	if req.Body != nil {
		jsonBody, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		strBody = string(jsonBody)
	}

	strQuery := ""
	if req.Query != nil {
		// attention: do not forget url tag after struct's fields
		q, err := goquery.Values(req.Query)
		if err != nil {
			return nil, err
		}
		strQuery = q.Encode()
	}

	timestamp := time.Now().UnixMilli()
	signString := fmt.Sprintf("%d%s%s%s", timestamp, bb.GetKey(), strBody, strQuery)

	h := hmac.New(sha256.New, []byte(bb.GetSecret()))
	h.Write([]byte(signString))
	signature := hex.EncodeToString(h.Sum(nil))

	headers["X-BAPI-API-KEY"] = bb.GetKey()
	headers["X-BAPI-SIGN"] = signature
	headers["X-BAPI-TIMESTAMP"] = strconv.FormatInt(timestamp, 10)

	return headers, nil
}

func (bb *BybitClient) SendHTTPRequest(ctx context.Context, req utils.HTTPRequest) (*utils.ApiResponse, error) {
	client := http.Client{}

	var body io.Reader
	if req.Body != nil {
		jsonBody, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(jsonBody)
	}

	url, err := url.Parse(req.BaseURL + req.Path)
	if err != nil {
		return nil, err
	}

	if req.Query != nil {
		q, err := goquery.Values(req.Query)
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

	if bb.GetDebug() {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}

		bb.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if bb.GetDebug() {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		bb.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	return utils.NewApiResponse(&req, resp), nil
}
