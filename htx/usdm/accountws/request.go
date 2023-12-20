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

const (
	SUB  = "sub"
	PING = "ping"
	PONG = "pong"
	AUTH = "auth"
)

type AuthRequest struct {
	Operation        string `json:"op,omitempty"`
	Type             string `json:"type,omitempty"`
	AccessKeyId      string `json:"AccessKeyId,omitempty"`
	SignatureMethod  string `json:"SignatureMethod,omitempty"`
	SignatureVersion string `json:"SignatureVersion,omitempty"`
	Timestamp        string `json:"Timestamp,omitempty"`
	Signature        string `json:"Signature,omitempty"`
}
