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

type Account struct {
	Currency    string `json:"currency,omitempty"`
	AccountID   int64  `json:"accountId,omitempty"`
	Balance     string `json:"balance,omitempty"`
	Available   string `json:"available,omitempty"`
	ChangeType  string `json:"changeType,omitempty"`
	AccountType string `json:"accountType,omitempty"`
	ChangeTime  int64  `json:"changeTime,omitempty"`
	SeqNum      int64  `json:"seqNum,omitempty"`
}
