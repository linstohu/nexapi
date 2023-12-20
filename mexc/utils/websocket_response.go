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
	"errors"

	"github.com/valyala/fastjson"
)

type Request struct {
	ID     uint32   `json:"id,omitempty"`
	Method string   `json:"method,omitempty"`
	Params []string `json:"params,omitempty"`
}

// AnyMessage represents either a JSON Response or SubscribedMessage.
type AnyMessage struct {
	Response          *Response
	SubscribedMessage *SubscribedMessage
}

type Response struct {
	ID     uint `json:"id"`
	Result any  `json:"result"`
}

type SubscribedMessage struct {
	Stream string          `json:"stream"`
	Data   json.RawMessage `json:"data"`
}

func (m AnyMessage) MarshalJSON() ([]byte, error) {
	var v any

	switch {
	case m.Response != nil && m.SubscribedMessage == nil:
		v = m.Response
	case m.Response == nil && m.SubscribedMessage != nil:
		v = m.SubscribedMessage
	}

	if v != nil {
		return json.Marshal(v)
	}

	return nil, errors.New("message must have exactly one of the Response or SubscribedMessage fields set")
}

func (m *AnyMessage) UnmarshalJSON(data []byte) error {
	var p fastjson.Parser
	v, err := p.ParseBytes(data)
	if err != nil {
		return err
	}

	if v.Exists("id") {
		var resp Response

		if err := json.Unmarshal(data, &resp); err != nil {
			return err
		}

		m.Response = &resp

		return nil
	}

	if v.Exists("stream") {
		msg := &SubscribedMessage{
			Stream: string(v.GetStringBytes("stream")),
		}

		if v.Get("data") != nil {
			msg.Data = v.Get("data").MarshalTo(nil)
		}

		m.SubscribedMessage = msg

		return nil
	}

	return nil
}
