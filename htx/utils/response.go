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
	"fmt"
	"net/http"
)

// A HTTPResponse represents a HTTP response.
type HTTPResponse struct {
	Req  *HTTPRequest
	Resp *http.Response
	Body []byte
}

// NewResponse Creates a new Response
func NewResponse(
	request *HTTPRequest,
	response *http.Response,
	body []byte,
) *HTTPResponse {
	return &HTTPResponse{
		Req:  request,
		Resp: response,
		Body: body,
	}
}

// ReadBody read the response data, then return it.
func (r *HTTPResponse) ReadBody() ([]byte, error) {
	if r.Body != nil {
		return r.Body, nil
	}

	r.Body = make([]byte, 0)
	defer r.Resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Resp.Body)

	r.Body = buf.Bytes()

	return r.Body, nil
}

// ReadJsonBody read the response data as JSON into v.
func (r *HTTPResponse) ReadJsonBody(v interface{}) error {
	b, err := r.ReadBody()
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

func (r *HTTPResponse) Error() string {
	uri, err := r.Req.RequestURI()
	if err != nil {
		return fmt.Sprintf("get request uri error: %s", err.Error())
	}

	reqBody, err := r.Req.RequestBody()
	if err != nil {
		return fmt.Sprintf("get request body error: %s", err.Error())
	}

	var body []byte
	if r.Body != nil {
		body = r.Body
	} else {
		body = []byte(NIL)
	}

	m := fmt.Sprintf("[Parse]Failure: parse JSON body failed, %s %s with body=%s, respond code=%d body=%s",
		r.Req.Method,
		uri,
		reqBody,
		r.Resp.StatusCode,
		string(body),
	)

	return m
}
