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
	"errors"

	"github.com/valyala/fastjson"
)

type Request struct {
	ID    string `json:"id,omitempty"`
	Topic string `json:"topic,omitempty"`
	Event string `json:"event,omitempty"`
}

// AnyMessage represents either a JSON Response or SubscribedMessage.
type AnyMessage struct {
	Response          *Response
	SubscribedMessage *SubscribedMessage
}

type Response struct {
	ID        string `json:"id"`
	Event     string `json:"event"`
	Success   bool   `json:"success"`
	Timestamp int64  `json:"ts"`
}

type SubscribedMessage struct {
	OriginData []byte
	Topic      string          `json:"topic"`
	Data       json.RawMessage `json:"data"`
	Timestamp  int64           `json:"ts"`
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

	return nil, errors.New("woox websocket: message must have exactly one of the Response or SubscribedMessage fields set")
}

func (m *AnyMessage) UnmarshalJSON(data []byte) error {
	var p fastjson.Parser
	v, err := p.ParseBytes(data)
	if err != nil {
		return err
	}

	if v.Exists("event") {
		var resp Response

		if err := json.Unmarshal(data, &resp); err != nil {
			return err
		}

		m.Response = &resp

		return nil
	}

	if v.Exists("topic") {
		des := make([]byte, len(data))
		copy(des, data)

		var msg = &SubscribedMessage{
			OriginData: des,
			Topic:      string(v.GetStringBytes("topic")),
			Timestamp:  v.GetInt64("ts"),
			Data:       v.GetStringBytes("data"),
		}

		m.SubscribedMessage = msg

		return nil
	}

	return nil
}
