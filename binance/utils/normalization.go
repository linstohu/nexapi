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
	goquery "github.com/google/go-querystring/query"
)

// NormalizeRequestContent sign endpoint security
// doc: https://binance-docs.github.io/apidocs/spot/en/#signed-trade-user_data-and-margin-endpoint-security
func NormalizeRequestContent(query any, body any) (string, error) {
	var ret string

	if query != nil {
		// attention: do not forget url tag after struct's fields
		q, err := goquery.Values(query)
		if err != nil {
			return "", err
		}
		ret += q.Encode()
	}

	if body != nil {
		// attention: do not forget url tag after struct's fields
		q, err := goquery.Values(body)
		if err != nil {
			return "", err
		}
		ret += q.Encode()
	}

	return ret, nil
}
