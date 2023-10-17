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

type AccountInfo struct {
	Asset []struct {
		Asset         string `json:"asset"`
		MarginBalance string `json:"marginBalance"`
		Equity        string `json:"equity"`
		Available     string `json:"available"`
		Locked        string `json:"locked"`
		UnrealizedPNL string `json:"unrealizedPNL"`
	} `json:"asset"`
	Greek []struct {
		Underlying string `json:"underlying"`
		Delta      string `json:"delta"`
		Gamma      string `json:"gamma"`
		Theta      string `json:"theta"`
		Vega       string `json:"vega"`
	} `json:"greek"`
	Time int64 `json:"time"`
}
