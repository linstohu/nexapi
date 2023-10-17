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
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type GetExerciseRecordParam struct {
	Symbol    string `url:"symbol,omitempty" validate:"omitempty"`
	FromID    int64  `url:"fromId,omitempty" validate:"omitempty"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type GetExerciseRecordParams struct {
	GetExerciseRecordParam
	bnutils.DefaultParam
}

type ExerciseRecord struct {
	ID            string `json:"id"`
	Currency      string `json:"currency"`
	Symbol        string `json:"symbol"`
	ExercisePrice string `json:"exercisePrice"`
	MarkPrice     string `json:"markPrice"`
	Quantity      string `json:"quantity"`
	Amount        string `json:"amount"`
	Fee           string `json:"fee"`
	CreateDate    int64  `json:"createDate"`
	PriceScale    int    `json:"priceScale"`
	QuantityScale int    `json:"quantityScale"`
	OptionSide    string `json:"optionSide"`
	PositionSide  string `json:"positionSide"`
	QuoteAsset    string `json:"quoteAsset"`
}
