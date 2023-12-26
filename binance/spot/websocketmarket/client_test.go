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

package websocketmarket_test

import (
	"fmt"
	"testing"
	"time"

	spotws "github.com/linstohu/nexapi/binance/spot/websocketmarket"
	"github.com/linstohu/nexapi/binance/spot/websocketmarket/types"
	"github.com/stretchr/testify/assert"
)

func testNewSpotMarketStreamClient(t *testing.T) *spotws.SpotMarketStreamClient {
	cli, err := spotws.NewSpotMarketStreamClient(&spotws.SpotMarketStreamCfg{
		Debug:         false,
		BaseURL:       spotws.SpotMarketStreamBaseURL,
		AutoReconnect: true,
	})

	if err != nil {
		t.Fatalf("Could not create websocket client, %s", err)
	}

	return cli
}

func TestSubscribeAggTrade(t *testing.T) {
	cli := testNewSpotMarketStreamClient(t)
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetAggTradeTopic("btcusdt")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		trade, ok := e.(*types.AggregateTrade)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, Price: %v, Quantity: %v, Time: %v\n",
			topic, trade.Symbol, trade.Price, trade.Quantity, trade.EventTime)
	})

	cli.Subscribe([]string{topic})

	time.Sleep(10 * time.Second)

	cli.Close()

	select {}
}
