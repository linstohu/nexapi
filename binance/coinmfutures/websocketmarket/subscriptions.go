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

	spottypes "github.com/linstohu/nexapi/binance/spot/websocketmarket/types"
	usdmtypes "github.com/linstohu/nexapi/binance/usdmfutures/websocketmarket/types"
	"github.com/linstohu/nexapi/binance/utils"
)

func (u *CoinMarginedMarketStreamClient) Subscribe(topics []string) error {
	return u.subscribe(topics)
}

func (u *CoinMarginedMarketStreamClient) UnSubscribe(topics []string) error {
	return u.unsubscribe(topics)
}

func (u *CoinMarginedMarketStreamClient) handle(msg *utils.SubscribedMessage) error {
	if u.debug {
		u.logger.Info(fmt.Sprintf("subscribed message, stream: %s", msg.Stream))
	}

	switch {
	case strings.HasSuffix(msg.Stream, "@aggTrade"):
		var data usdmtypes.AggregateTrade
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		u.GetListeners(msg.Stream, &data)
	case strings.Contains(msg.Stream, "@indexPrice"):
		var data IndexPrice
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		u.GetListeners(msg.Stream, &data)
	case strings.Contains(msg.Stream, "@markPrice"):
		before, _, _ := strings.Cut(msg.Stream, "@markPrice")
		if strings.Contains(before, "_") {
			// Mark Price Stream
			var data MarkPrice
			err := json.Unmarshal(msg.Data, &data)
			if err != nil {
				return err
			}
			u.GetListeners(msg.Stream, &data)
		} else {
			// Mark Price of All Symbols of a Pair
			var data []*MarkPrice
			err := json.Unmarshal(msg.Data, &data)
			if err != nil {
				return err
			}
			u.GetListeners(msg.Stream, data)
		}
	case strings.Contains(msg.Stream, "@kline_"):
		var data spottypes.Kline
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		u.GetListeners(msg.Stream, &data)
	case strings.HasSuffix(msg.Stream, "@miniTicker"):
		var data MiniTicker
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		u.GetListeners(msg.Stream, &data)
	case msg.Stream == "!miniTicker@arr":
		var data []*MiniTicker
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		u.GetListeners(msg.Stream, data)
	case strings.HasSuffix(msg.Stream, "@ticker"):
		var data Ticker
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		u.GetListeners(msg.Stream, &data)
	case msg.Stream == "!ticker@arr":
		var data []*Ticker
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		u.GetListeners(msg.Stream, data)
	case strings.HasSuffix(msg.Stream, "@bookTicker") || msg.Stream == "!bookTicker":
		var data BookTicker
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		u.GetListeners(msg.Stream, &data)
	case strings.HasSuffix(msg.Stream, "@forceOrder") || msg.Stream == "!forceOrder@arr":
		var data LiquidationOrder
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		u.GetListeners(msg.Stream, &data)
	case strings.Contains(msg.Stream, "@depth"):
		var data OrderbookDepth
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		u.GetListeners(msg.Stream, &data)
	default:
		return fmt.Errorf("unknown message, topic: %s", msg.Stream)
	}

	return nil
}
