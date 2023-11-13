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
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
)

type HTTPRequest struct {
	BaseURL string
	Path    string
	Method  string
	Headers map[string]string
	Query   any
	Body    any
}

// RequestURI returns the request uri.
func (h *HTTPRequest) RequestURI() (string, error) {
	url, err := url.Parse(h.BaseURL + h.Path)
	if err != nil {
		return "", err
	}

	if h.Query != nil {
		q, err := query.Values(h.Query)
		if err != nil {
			return "", err
		}
		url.RawQuery = q.Encode()
	}

	return url.RequestURI(), nil
}

func (h *HTTPRequest) RequestBody() (string, error) {
	if h.Body == nil {
		return NIL, nil
	}

	var body string
	if h.Body != nil {
		jsonBody, err := json.Marshal(h.Body)
		if err != nil {
			return "", err
		}
		body = string(jsonBody)
	}

	return body, nil
}
