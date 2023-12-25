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

var (
	BaseURL = "https://api.binance.com"
)

type SecurityType string

var (
	NONE        SecurityType = "NONE"
	TRADE       SecurityType = "TRADE"
	MARGIN      SecurityType = "MARGIN"
	USER_DATA   SecurityType = "USER_DATA"
	USER_STREAM SecurityType = "USER_STREAM"
	MARKET_DATA SecurityType = "MARKET_DATA"
)

type KlineInterval string

var (
	Second1  KlineInterval = "1s"
	Minute1  KlineInterval = "1m"
	Minute3  KlineInterval = "3m"
	Minute5  KlineInterval = "5m"
	Minute15 KlineInterval = "15m"
	Minute30 KlineInterval = "30m"
	Hour1    KlineInterval = "1h"
	Hour2    KlineInterval = "2h"
	Hour4    KlineInterval = "4h"
	Hour6    KlineInterval = "6h"
	Hour8    KlineInterval = "8h"
	Hour12   KlineInterval = "12h"
	Day1     KlineInterval = "1d"
	Day3     KlineInterval = "3d"
	Week1    KlineInterval = "1w"
	Month1   KlineInterval = "1M"
)
