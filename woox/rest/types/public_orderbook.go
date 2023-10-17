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

type GetOrderbookParam struct {
	MaxLevel int `url:"max_level,omitempty"`
}

type Orderbook struct {
	Response
	Asks []struct {
		Price    float64 `json:"price"`
		Quantity float64 `json:"quantity"`
	} `json:"asks"`
	Bids []struct {
		Price    float64 `json:"price"`
		Quantity float64 `json:"quantity"`
	} `json:"bids"`
	Timestamp int64 `json:"timestamp"`
}
