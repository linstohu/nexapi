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
	"strconv"
)

// A PaginationParam represents the pagination parameters `currentPage` `pageSize` in a request .
type PaginationParam struct {
	CurrentPage int64
	PageSize    int64
}

// ReadParam read pagination parameters into params.
func (p *PaginationParam) ReadParam(params map[string]string) {
	params["currentPage"], params["pageSize"] = IntToString(p.CurrentPage), IntToString(p.PageSize)
}

// A PaginationModel represents the pagination in a response.
type PaginationModel struct {
	CurrentPage int64           `json:"currentPage"`
	PageSize    int64           `json:"pageSize"`
	TotalNum    int64           `json:"totalNum"`
	TotalPage   int64           `json:"totalPage"`
	RawItems    json.RawMessage `json:"items"` // delay parsing
}

// ReadItems read the `items` into v.
func (p *PaginationModel) ReadItems(v interface{}) error {
	return json.Unmarshal(p.RawItems, v)
}

// IntToString converts int64 to string.
func IntToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// ToJsonString converts any value to JSON string.
func ToJsonString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}
