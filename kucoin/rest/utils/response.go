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

// The predefined API codes
const (
	ApiSuccess = "200000"
)

// An ApiResponse represents a API response wrapped Response.
type ApiResponse struct {
	Resp    *HTTPResponse
	Code    string          `json:"code"`
	RawData json.RawMessage `json:"data"` // delay parsing
	Message string          `json:"msg"`
}

// HttpSuccessful judges the success of http.
func (ar *ApiResponse) HttpSuccessful() bool {
	return ar.Resp.Resp.StatusCode == http.StatusOK
}

// ApiSuccessful judges the success of API.
func (ar *ApiResponse) ApiSuccessful() bool {
	return ar.Code == ApiSuccess
}

// ReadData read the api response `data` as JSON into v.
func (ar *ApiResponse) ReadData(v interface{}) error {
	reqURI, err := ar.Resp.Req.RequestURI()
	if err != nil {
		return err
	}

	reqBody, err := ar.Resp.Req.RequestBody()
	if err != nil {
		return err
	}

	if !ar.HttpSuccessful() {
		rsb, _ := ar.Resp.ReadBody()
		m := fmt.Sprintf("[HTTP]Failure: status code is NOT 200, %s %s with body=%s, respond code=%d body=%s",
			ar.Resp.Req.Method,
			reqURI,
			reqBody,
			ar.Resp.Resp.StatusCode,
			string(rsb),
		)
		return errors.New(m)
	}

	if !ar.ApiSuccessful() {
		m := fmt.Sprintf("[API]Failure: api code is NOT %s, %s %s with body=%s, respond code=%s message=\"%s\" data=%s",
			ApiSuccess,
			ar.Resp.Req.Method,
			reqURI,
			reqBody,
			ar.Code,
			ar.Message,
			string(ar.RawData),
		)
		return errors.New(m)
	}
	// when input parameter v is nil, read nothing and return nil
	if v == nil {
		return nil
	}

	if len(ar.RawData) == 0 {
		m := fmt.Sprintf("[API]Failure: try to read empty data, %s %s with body=%s, respond code=%s message=\"%s\" data=%s",
			ar.Resp.Req.Method,
			reqURI,
			reqBody,
			ar.Code,
			ar.Message,
			string(ar.RawData),
		)
		return errors.New(m)
	}

	return json.Unmarshal(ar.RawData, v)
}

// ReadPaginationData read the data `items` as JSON into v, and returns *PaginationModel.
func (ar *ApiResponse) ReadPaginationData(v interface{}) (*PaginationModel, error) {
	p := &PaginationModel{}
	if err := ar.ReadData(p); err != nil {
		return nil, err
	}
	if err := p.ReadItems(v); err != nil {
		return p, err
	}
	return p, nil
}
