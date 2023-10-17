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

type Balance struct {
	Response
	Data struct {
		Holding []struct {
			Token            string  `json:"token"`
			Holding          float64 `json:"holding"`
			Frozen           float64 `json:"frozen"`
			Staked           float64 `json:"staked"`
			Unbonding        float64 `json:"unbonding"`
			Vault            float64 `json:"vault"`
			Interest         float64 `json:"interest"`
			PendingShortQty  float64 `json:"pendingShortQty"`
			PendingLongQty   float64 `json:"pendingLongQty"`
			AvailableBalance float64 `json:"availableBalance"`
			UpdatedTime      float64 `json:"updatedTime"`
		} `json:"holding"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}
