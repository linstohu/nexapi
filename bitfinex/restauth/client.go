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

package bitfinexrestauth

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-playground/validator"
	"github.com/google/go-querystring/query"
	"github.com/linstohu/nexapi/bitfinex/utils"
)

type BitfinexAuthClient struct {
	// debug mode
	debug bool
	// logger
	logger *slog.Logger
	nonce  utils.NonceGenerator

	baseURL     string
	key, secret string
}

type BitfinexClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewBitfinexClient(cfg *BitfinexClientCfg) (*BitfinexAuthClient, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := BitfinexAuthClient{
		debug:   cfg.Debug,
		logger:  cfg.Logger,
		nonce:   utils.NewEpochNonceGenerator(),
		baseURL: cfg.BaseURL,
		key:     cfg.Key,
		secret:  cfg.Secret,
	}

	if cli.logger == nil {
		cli.logger = slog.Default()
	}

	return &cli, nil
}

type HTTPRequest struct {
	BaseURL string
	Path    string
	Method  string
	Headers map[string]string
	Query   any
	Body    any
}

func (b *BitfinexAuthClient) GetDebug() bool {
	return b.debug
}

func (b *BitfinexAuthClient) GetBaseURL() string {
	return b.baseURL
}

func (b *BitfinexAuthClient) GenHeaders(req HTTPRequest) (map[string]string, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	nonce := b.nonce.GetNonce()
	signaturePayload := fmt.Sprintf("/api%s%s", req.Path, nonce)

	if req.Body != nil {
		jsonBody, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		signaturePayload += string(jsonBody)
	}

	sig := hmac.New(sha512.New384, []byte(b.secret))
	_, err := sig.Write([]byte(signaturePayload))
	if err != nil {
		return nil, err
	}

	headers["bfx-nonce"] = nonce
	headers["bfx-signature"] = hex.EncodeToString(sig.Sum(nil))
	headers["bfx-apikey"] = b.key

	return headers, nil
}

func (b *BitfinexAuthClient) SendHTTPRequest(ctx context.Context, req HTTPRequest) ([]byte, error) {
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

	if b.GetDebug() {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		b.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if b.GetDebug() {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		b.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned a non-200 status code: [%d] - [%s]", resp.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}
