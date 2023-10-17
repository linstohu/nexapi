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

package ctutils

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
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/go-querystring/query"
)

type ContractClient struct {
	// debug mode
	debug bool
	// logger
	logger *slog.Logger

	baseURL     string
	key, secret string
	recvWindow  int
}

type ContractClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string
	Secret     string
	RecvWindow int
}

func NewContractClient(cfg *ContractClientCfg) (*ContractClient, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := ContractClient{
		debug:      cfg.Debug,
		logger:     cfg.Logger,
		baseURL:    cfg.BaseURL,
		key:        cfg.Key,
		secret:     cfg.Secret,
		recvWindow: cfg.RecvWindow,
	}

	if cfg.RecvWindow == 0 {
		cli.recvWindow = 10
	}

	if cli.logger == nil {
		cli.logger = slog.Default()
	}

	return &cli, nil
}

func (c *ContractClient) GetDebug() bool {
	return c.debug
}

func (c *ContractClient) GetBaseURL() string {
	return c.baseURL
}

func (c *ContractClient) GetKey() string {
	return c.key
}

func (c *ContractClient) GetSecret() string {
	return c.secret
}

func (c *ContractClient) GetRecvWindow() int {
	return c.recvWindow
}

func (c *ContractClient) GenPubHeaders() (map[string]string, error) {
	return map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}, nil
}

func (c *ContractClient) GenAuthHeaders(req HTTPRequest) (map[string]string, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	var signString string

	switch req.Method {
	case http.MethodGet, http.MethodDelete:
		// attention: do not forget url tag after struct's fields
		q, err := query.Values(req.Query)
		if err != nil {
			return nil, err
		}

		signString = q.Encode()
	case http.MethodPost:
		if req.Body != nil {
			jsonBody, err := json.Marshal(req.Body)
			if err != nil {
				return nil, err
			}
			signString = string(jsonBody)
		}
	default:
		return nil, fmt.Errorf("unknown request method")
	}

	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())

	if signString != "" {
		sign := fmt.Sprintf("%s%s%s", c.key, timestamp, signString)
		h := hmac.New(sha256.New, []byte(c.secret))
		h.Write([]byte(sign))
		signature := hex.EncodeToString(h.Sum(nil))
		headers["Signature"] = signature
	}

	headers["ApiKey"] = c.key
	headers["Request-Time"] = timestamp

	return headers, nil
}

func (c *ContractClient) SendHTTPRequest(ctx context.Context, req HTTPRequest) ([]byte, error) {
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

	if c.GetDebug() {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		c.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if c.GetDebug() {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		c.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned a non-200 status code: [%d] - [%s]", resp.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}
