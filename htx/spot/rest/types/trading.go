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
	htxutils "github.com/linstohu/nexapi/htx/utils"
)

type NewOrderParam struct {
	AccountID        string `json:"account-id" validate:"required"`
	Symbol           string `json:"symbol" validate:"required"`
	Type             string `json:"type" validate:"required"`
	Amount           string `json:"amount" validate:"required"`
	Price            string `json:"price,omitempty" validate:"omitempty"`
	Source           string `json:"source,omitempty" validate:"omitempty"`
	ClientOrderID    string `json:"client-order-id,omitempty" validate:"omitempty"`
	SelfMatchPrevent int    `json:"self-match-prevent,omitempty" validate:"omitempty"`
	StopPrice        string `json:"stop-price,omitempty" validate:"omitempty"`
	Operator         string `json:"operator,omitempty" validate:"omitempty"`
}

type NewOrderResp struct {
	htxutils.V1Response
	Data string `json:"data,omitempty"`
}
