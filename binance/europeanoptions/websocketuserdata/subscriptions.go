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

package websocketuserdata

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

func (o *OptionsUserDataStreamClient) handle(origind []byte) error {
	var p fastjson.Parser
	pb, err := p.ParseBytes(origind)
	if err != nil {
		return err
	}

	eventType := string(pb.GetStringBytes("e"))

	if o.debug {
		o.logger.Info(fmt.Sprintf("%s: subscribed message, event-type: %s", logPrefix, eventType))
	}

	switch eventType {
	case o.GenAccountDataTopic():
		var data AccountData
		err := json.Unmarshal(origind, &data)
		if err != nil {
			return err
		}
		o.GetListeners(o.GenAccountDataTopic(), &data)
	case o.GenOrderUpdateTopic():
		var data OrderUpdate
		err := json.Unmarshal(origind, &data)
		if err != nil {
			return err
		}
		o.GetListeners(o.GenOrderUpdateTopic(), &data)
	default:
		return fmt.Errorf("unknown message, event-type: %s", eventType)
	}

	return nil
}

func (o *OptionsUserDataStreamClient) GenAccountDataTopic() string {
	return "ACCOUNT_UPDATE"
}

func (o *OptionsUserDataStreamClient) GenOrderUpdateTopic() string {
	return "ORDER_TRADE_UPDATE"
}
