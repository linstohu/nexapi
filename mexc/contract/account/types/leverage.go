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

type GetLeverageParams struct {
	Symbol string `url:"symbol,omitempty" validate:"required"`
}

type GetLeverageResp struct {
	Response
	Data struct {
		PositionType int     `json:"positionType"`
		Level        int     `json:"level"`
		Imr          float64 `json:"imr"`
		Mmr          float64 `json:"mmr"`
		Leverage     int     `json:"leverage"`
	} `json:"data"`
}

type SetLeverageParams struct {
	PositionId   int64  `url:"positionId,omitempty" validate:"omitempty"`
	Leverage     int    `url:"positionId,omitempty" validate:"required"`
	OpenType     int    `url:"openType,omitempty" validate:"omitempty"`
	Symbol       string `url:"symbol,omitempty" validate:"omitempty"`
	PositionType int    `url:"positionType,omitempty" validate:"omitempty"`
}

type SetLeverageResp struct {
	Response
	Data struct {
		PositionId   int64  `json:"positionId"`
		Leverage     int    `json:"leverage"`
		Symbol       string `json:"symbol"`
		PositionType int    `json:"positionType"`
	} `json:"data"`
}
