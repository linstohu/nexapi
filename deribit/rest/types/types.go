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

package types

import (
	"encoding/json"
	"fmt"
)

type HTTPRequest struct {
	URL     string
	Method  string
	Headers map[string]string
	Query   any
	Body    *Body
	Debug   bool
}

type Body struct {
	Jsonrpc string `json:"jsonrpc,omitempty"`
	ID      any    `json:"id,omitempty"`
	Method  string `json:"method,omitempty"`
	Params  any    `json:"params,omitempty"`
}

type JsonrpcMessage struct {
	Jsonrpc string          `json:"jsonrpc,omitempty"`
	ID      any             `json:"id,omitempty"`
	Error   *JsonError      `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	UsIn    int64           `json:"usIn,omitempty"`
	UsOut   int64           `json:"usOut,omitempty"`
	UsDiff  int64           `json:"usDiff,omitempty"`
	TestNet bool            `json:"testnet,omitempty"`
}

type JsonError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func (e *JsonError) Error() string {
	return fmt.Sprintf("json-rpc error (%d): %s", e.Code, e.Message)
}
