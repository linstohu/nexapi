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

type ActionType = string

const (
	SUB   ActionType = "sub"
	UNSUB ActionType = "unsub"
	REQ   ActionType = "req"
	PING  ActionType = "ping"
	PONG  ActionType = "pong"
	PUSH  ActionType = "push"
)

type Message struct {
	Action  ActionType      `json:"action,omitempty"`
	Channel string          `json:"ch,omitempty"`
	Code    int             `json:"code,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func (m *Message) UnmarshalJSON(data []byte) error {
	var p fastjson.Parser
	v, err := p.ParseBytes(data)
	if err != nil {
		return err
	}

	m.Action = string(v.GetStringBytes("action"))
	m.Channel = string(v.GetStringBytes("ch"))
	m.Code = v.GetInt("code")
	if v.Get("data") != nil {
		m.Data = v.Get("data").MarshalTo(nil)
	}

	return nil
}
