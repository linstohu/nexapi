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

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/linstohu/nexapi/binance/utils"
)

func (o *OptionsMarketStreamClient) Subscribe(topics []string) error {
	return o.subscribe(topics)
}

func (o *OptionsMarketStreamClient) UnSubscribe(topics []string) error {
	return o.unsubscribe(topics)
}

func (o *OptionsMarketStreamClient) handle(msg *utils.SubscribedMessage) error {
	if o.debug {
		o.logger.Info(fmt.Sprintf("subscribed message, stream: %s", msg.Stream))
	}

	switch {
	case strings.HasSuffix(msg.Stream, "@trade"):
		var data Trade
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		o.GetListeners(msg.Stream, &data)
	case strings.HasSuffix(msg.Stream, "@index"):
		var data IndexPrice
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		o.GetListeners(msg.Stream, &data)
	case strings.HasSuffix(msg.Stream, "@markPrice"):
		var data []*MarkPrice
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		o.GetListeners(msg.Stream, data)
	case strings.Contains(msg.Stream, "@kline_"):
		var data Kline
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		o.GetListeners(msg.Stream, &data)
	case strings.Contains(msg.Stream, "@ticker"):
		if strings.HasSuffix(msg.Stream, "@ticker") {
			// 24hr ticker info for all symbols. Only symbols whose ticker info changed will be sent.
			var data Ticker
			err := json.Unmarshal(msg.Data, &data)
			if err != nil {
				return err
			}
			o.GetListeners(msg.Stream, &data)
		} else {
			// 24hr ticker info by underlying asset and expiration date
			var data []*Ticker
			err := json.Unmarshal(msg.Data, &data)
			if err != nil {
				return err
			}
			o.GetListeners(msg.Stream, data)
		}
	case strings.Contains(msg.Stream, "@openInterest"):
		var data []*OpenInterest
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		o.GetListeners(msg.Stream, data)
	case strings.Contains(msg.Stream, "@depth"):
		var data OrderbookDepth
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		o.GetListeners(msg.Stream, &data)
	default:
		return fmt.Errorf("unknown message, topic: %s", msg.Stream)
	}

	return nil
}
