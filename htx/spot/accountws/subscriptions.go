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

package accountws

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/linstohu/nexapi/htx/spot/accountws/types"
)

func (m *AccountWsClient) Subscribe(topic string) error {
	return m.subscribe(topic)
}

func (m *AccountWsClient) UnSubscribe(topic string) error {
	return m.unsubscribe(topic)
}

func (m *AccountWsClient) handle(msg *Message) error {
	if m.debug {
		m.logger.Info(fmt.Sprintf("subscribed message, channel: %s", msg.Channel))
	}

	switch {
	case strings.HasPrefix(msg.Channel, "accounts.update"):
		var data types.Account
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
