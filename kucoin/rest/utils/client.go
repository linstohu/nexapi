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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/go-querystring/query"
)

type KucoinClient struct {
	// debug mode
	debug bool
	// logger
	logger *slog.Logger

	baseURL                             string
	key, secret, passphrase, keyVersion string
}

type KucoinClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string
	KeyVersion string
	Secret     string
	Passphrase string
}

func NewKucoinRestClient(cfg *KucoinClientCfg) (*KucoinClient, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := KucoinClient{
		debug:      cfg.Debug,
		logger:     cfg.Logger,
		baseURL:    cfg.BaseURL,
		key:        cfg.Key,
		keyVersion: cfg.KeyVersion,
		secret:     cfg.Secret,
		passphrase: sign([]byte(cfg.Secret), []byte(cfg.Passphrase)),
	}

	if cli.logger == nil {
		cli.logger = slog.Default()
	}

	return &cli, nil
}

func (k *KucoinClient) GetDebug() bool {
	return k.debug
}

func (k *KucoinClient) GetBaseURL() string {
	return k.baseURL
}

func (k *KucoinClient) GetKey() string {
	return k.key
}

func (k *KucoinClient) GetSecret() string {
	return k.secret
}

func (k *KucoinClient) GetPassphrase() string {
	return k.passphrase
}

func (k *KucoinClient) GetHeaders() (map[string]string, error) {
	return map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}, nil
}

func (k *KucoinClient) GenSignature(req HTTPRequest) (map[string]string, error) {
	uri, err := req.RequestURI()
	if err != nil {
		return nil, err
	}

	reqBody, err := req.RequestBody()
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer

	b.WriteString(req.Method)
	b.WriteString(uri)

	if reqBody != NIL {
		b.WriteString(reqBody)
	}

	t := time.Now().UnixMilli()

	signStr := fmt.Sprintf("%v%s", t, b.String())
	s := sign([]byte(k.secret), []byte(signStr))

	ksHeaders := map[string]string{
		"KC-API-KEY":         k.key,
		"KC-API-PASSPHRASE":  k.passphrase,
		"KC-API-TIMESTAMP":   fmt.Sprintf("%v", t),
		"KC-API-SIGN":        s,
		"KC-API-KEY-VERSION": k.keyVersion,
	}

	return ksHeaders, nil
}

func (s *KucoinClient) SendHTTPRequest(ctx context.Context, req HTTPRequest) (*HTTPResponse, error) {
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

	return NewResponse(&req, resp, nil), nil
}

// sign makes a signature by sha256.
func sign(key, plain []byte) string {
	hm := hmac.New(sha256.New, key)
	hm.Write(plain)
	return base64.StdEncoding.EncodeToString(hm.Sum(nil))
}
