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

type KlineInterval string

var (
	Minute1  KlineInterval = "1min"
	Minute5  KlineInterval = "5min"
	Minute15 KlineInterval = "15min"
	Minute30 KlineInterval = "30min"
	Hour1    KlineInterval = "1hour"
	Hour4    KlineInterval = "4hour"
	Day1     KlineInterval = "1day"
	Month1   KlineInterval = "1mon"
)
