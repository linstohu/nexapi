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

package accountws

import (
	"encoding/json"

	"github.com/valyala/fastjson"
)

type Message struct {
	Operation string `json:"op,omitempty"`
	Topic     string `json:"topic,omitempty"`
	Ts        int64  `json:"ts,omitempty"`

	ErrCode int    `json:"err-code,omitempty"`
	ErrMsg  string `json:"err-msg,omitempty"`

	Raw json.RawMessage `json:"-"`
}

func (m *Message) UnmarshalJSON(data []byte) error {
	var p fastjson.Parser
	v, err := p.ParseBytes(data)
	if err != nil {
		return err
	}

	m.Operation = string(v.GetStringBytes("op"))
	m.Topic = string(v.GetStringBytes("topic"))
	m.Ts = v.GetInt64("ts")

	m.ErrCode = v.GetInt("err-code")
	m.ErrMsg = string(v.GetStringBytes("err-msg"))

	m.Raw = data

	return nil
}
