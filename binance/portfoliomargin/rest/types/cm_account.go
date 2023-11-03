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

type CMAccount struct {
	Assets []struct {
		Asset                  string `json:"asset"`
		CrossWalletBalance     string `json:"crossWalletBalance"`
		CrossUnPnl             string `json:"crossUnPnl"`
		MaintMargin            string `json:"maintMargin"`
		InitialMargin          string `json:"initialMargin"`
		PositionInitialMargin  string `json:"positionInitialMargin"`
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
		UpdateTime             int64  `json:"updateTime"`
	} `json:"assets"`
	Positions []struct {
		Symbol                 string `json:"symbol"`
		PositionAmt            string `json:"positionAmt"`
		InitialMargin          string `json:"initialMargin"`
		MaintMargin            string `json:"maintMargin"`
		UnrealizedProfit       string `json:"unrealizedProfit"`
		PositionInitialMargin  string `json:"positionInitialMargin"`
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
		Leverage               string `json:"leverage"`
		PositionSide           string `json:"positionSide"`
		EntryPrice             string `json:"entryPrice"`
		MaxQty                 string `json:"maxQty"`
		UpdateTime             int64  `json:"updateTime"`
	} `json:"positions"`
}
