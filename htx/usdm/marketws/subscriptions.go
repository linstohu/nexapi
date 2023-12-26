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

package marketws

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/linstohu/nexapi/htx/usdm/marketws/types"
)

func (m *MarketWsClient) Subscribe(topic string) error {
	return m.subscribe(topic)
}

func (m *MarketWsClient) UnSubscribe(topic string) error {
	return m.unsubscribe(topic)
}

func (m *MarketWsClient) handle(msg *SubscribedMessage) error {
	if m.debug {
		m.logger.Info(fmt.Sprintf("%s: subscribed message, channel: %s", logPrefix, msg.Channel))
	}

	switch {
	case strings.Contains(msg.Channel, "kline"):
		var data types.Kline
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		m.GetListeners(msg.Channel, &data)
	case strings.Contains(msg.Channel, "depth"):
		var data types.Depth
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		m.GetListeners(msg.Channel, &data)
	case strings.Contains(msg.Channel, "bbo"):
		var data types.BBO
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		m.GetListeners(msg.Channel, &data)
	case strings.Contains(msg.Channel, "trade"):
		var data types.MarketTradeMsg
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		m.GetListeners(msg.Channel, &data)
	default:
		return fmt.Errorf("unknown message, topic: %s", msg.Channel)
	}

	return nil
}
