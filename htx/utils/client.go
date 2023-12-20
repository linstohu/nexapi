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
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/go-playground/validator"
	goquery "github.com/google/go-querystring/query"
)

type HTXClient struct {
	// debug mode
	debug bool
	// logger
	logger *slog.Logger

	baseURL     string
	key, secret string
	signVersion string
}

type HTXClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL     string `validate:"required"`
	Key         string
	Secret      string
	SignVersion string
}

func NewHTXRestClient(cfg *HTXClientCfg) (*HTXClient, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := HTXClient{
		debug:       cfg.Debug,
		logger:      cfg.Logger,
		baseURL:     cfg.BaseURL,
		key:         cfg.Key,
		secret:      cfg.Secret,
		signVersion: cfg.SignVersion,
	}

	if cli.logger == nil {
		cli.logger = slog.Default()
	}

	return &cli, nil
}

func (htx *HTXClient) GetDebug() bool {
	return htx.debug
}

func (htx *HTXClient) GetBaseURL() string {
	return htx.baseURL
}

func (htx *HTXClient) GetKey() string {
	return htx.key
}

func (htx *HTXClient) GetSecret() string {
	return htx.secret
}

func (htx *HTXClient) GetSignVersion() string {
	return htx.signVersion
}

func (htx *HTXClient) GetHeaders() (map[string]string, error) {
	return map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}, nil
}

var ErrAuth = errors.New("auth error, you should reinitialize client using key and secret")

func (htx *HTXClient) CheckAuth() error {
	if htx.GetKey() == "" || htx.GetSecret() == "" || htx.GetSignVersion() == "" {
		return ErrAuth
	}

	return nil
}

func (htx *HTXClient) GenAuthParams() DefaultAuthParam {
	return DefaultAuthParam{
		AccessKeyId:      htx.key,
		SignatureMethod:  "HmacSHA256",
		SignatureVersion: htx.signVersion,
		Timestamp:        time.Now().UTC().Format("2006-01-02T15:04:05"),
	}
}

func (htx *HTXClient) NormalizeRequestContent(req HTTPRequest, tempQuery any) (string, error) {
	if req.Method == "" || req.BaseURL == "" || req.Path == "" {
		return "", fmt.Errorf("gen signature error: method(%s), baseurl(%s) and path(%s) should not be empty",
			req.Method, req.BaseURL, req.Path)
	}

	parameters := url.Values{}

	if tempQuery != nil {
		q, err := goquery.Values(tempQuery)
		if err != nil {
			return "", err
		}
		parameters = q
	}

	urls, err := url.Parse(req.BaseURL + req.Path)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString(req.Method)
	sb.WriteString("\n")
	sb.WriteString(urls.Host)
	sb.WriteString("\n")
	sb.WriteString(req.Path)
	sb.WriteString("\n")
	sb.WriteString(parameters.Encode())

	return sb.String(), nil
}

// sign makes a signature by sha256.
func (htx *HTXClient) Sign(plain []byte) string {
	hm := hmac.New(sha256.New, []byte(htx.secret))
	hm.Write(plain)
	return base64.StdEncoding.EncodeToString(hm.Sum(nil))
}

func (htx *HTXClient) SendHTTPRequest(ctx context.Context, req HTTPRequest) (*HTTPResponse, error) {
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

	if htx.GetDebug() {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}

		htx.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if htx.GetDebug() {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		htx.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	return NewResponse(&req, resp, nil), nil
}
