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
	BaseURL = "https://contract.mexc.com"
)

type KlineInterval string

var (
	Minute1  KlineInterval = "Min1"
	Minute5  KlineInterval = "Min5"
	Minute15 KlineInterval = "Min15"
	Minute30 KlineInterval = "Min30"
	Minute60 KlineInterval = "Min60"
	Hour4    KlineInterval = "Hour4"
	Hour8    KlineInterval = "Hour8"
	Day1     KlineInterval = "Day1"
	Week1    KlineInterval = "Week1"
	Month1   KlineInterval = "Month1"
)
