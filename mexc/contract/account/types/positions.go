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

type GetOpenPositionsParams struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
}

type GetOpenPositions struct {
	Response
	Data []*OpenPosition `json:"data"`
}

type OpenPosition struct {
	PositionID   int64  `json:"positionId"`
	Symbol       string `json:"symbol"`
	PositionType int    `json:"positionType"`
	OpenType     int    `json:"openType"`
	State        int    `json:"state"`

	FrozenVol      float64 `json:"frozenVol"`
	CloseVol       float64 `json:"closeVol"`
	HoldAvgPrice   float64 `json:"holdAvgPrice"`
	CloseAvgPrice  float64 `json:"closeAvgPrice"`
	OpenAvgPrice   float64 `json:"openAvgPrice"`
	LiquidatePrice float64 `json:"liquidatePrice"`
	Oim            float64 `json:"oim"`
	Im             float64 `json:"im"`
	HoldFee        float64 `json:"holdFee"`
	Realised       float64 `json:"realised"`

	HoldVol    float64 `json:"holdVol"`
	Leverage   int     `json:"leverage"`
	CreateTime int64   `json:"createTime"`
	UpdateTime int64   `json:"updateTime"`
	AutoAddIm  bool    `json:"autoAddIm"`
}
