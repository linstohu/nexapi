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
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type HTTPRequest struct {
	Debug   bool
	BaseURL string
	Path    string
	Method  string
	Headers map[string]string
	Query   any
	Body    any
}

type ApiResponse struct {
	ApiReq *HTTPRequest
	ApiRes *http.Response
	Body   []byte
}

// NewResponse Creates a new Response
func NewApiResponse(request *HTTPRequest, response *http.Response) *ApiResponse {
	return &ApiResponse{
		ApiReq: request,
		ApiRes: response,
	}
}

// ReadBody read the response data, then return it.
func (r *ApiResponse) ReadBody() ([]byte, error) {
	if r.Body != nil {
		return r.Body, nil
	}

	r.Body = make([]byte, 0)
	defer r.ApiRes.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.ApiRes.Body)

	if r.ApiRes.StatusCode != http.StatusOK {
		m := fmt.Sprintf("[HTTP]Failure: status code is NOT 200, %s %s, respond code=%d body=%s",
			r.ApiReq.Method,
			r.ApiReq.BaseURL+r.ApiReq.Path,
			r.ApiRes.StatusCode,
			buf.String(),
		)
		return nil, errors.New(m)
	}

	r.Body = buf.Bytes()

	return r.Body, nil
}

// ReadJsonBody read the response data as JSON into v.
func (r *ApiResponse) ReadJsonBody(v interface{}) error {
	b, err := r.ReadBody()
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}
