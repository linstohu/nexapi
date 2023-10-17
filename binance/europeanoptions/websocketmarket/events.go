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

package websocketmarket

import "github.com/chuckpreslar/emission"

type Listener func(any)

func (o *OptionsMarketStreamClient) AddListener(event string, listener Listener) *emission.Emitter {
	return o.emitter.On(event, listener)
}

func (o *OptionsMarketStreamClient) RemoveListener(event string, listener Listener) *emission.Emitter {
	return o.emitter.Off(o, listener)
}

func (o *OptionsMarketStreamClient) GetListeners(event string, argument any) *emission.Emitter {
	return o.emitter.Emit(event, argument)
}
