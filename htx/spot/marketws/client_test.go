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
	"fmt"
	"testing"

	"github.com/linstohu/nexapi/htx/spot/marketws/types"
	resttypes "github.com/linstohu/nexapi/htx/spot/rest/types"
	"github.com/stretchr/testify/assert"
)

func testNewMarketWsClient(t *testing.T, url string) *MarketWsClient {
	cli, err := NewMarketWsClient(&MarketWsClientCfg{
		Debug:         true,
		BaseURL:       url,
		AutoReconnect: true,
	})

	if err != nil {
		t.Fatalf("Could not create websocket client, %s", err)
	}

	return cli
}

func TestSubscribeKline(t *testing.T) {
	cli := testNewMarketWsClient(t, GlobalWsBaseURL)

	topic, err := cli.GetKlineTopic(&KlineTopicParam{
		Symbol:   "btcusdt",
		Interval: resttypes.Minute1,
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		kline, ok := e.(*types.Kline)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Open: %v, Close: %v, Low: %v, High: %v, Amount: %v\n",
			topic, kline.Open, kline.Close, kline.Low, kline.High, kline.Amount)
	})

	cli.Subscribe(topic)

	select {}
}

func TestSubscribeMBPUpdateDepth(t *testing.T) {
	cli := testNewMarketWsClient(t, MBPWsBaseURL)

	topic, err := cli.GetMBPDepthUpdateTopic(&MBPDepthUpdateTopicParam{
		Symbol: "btcusdt",
		Level:  5,
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		depth, ok := e.(*types.MBPRefreshDepth)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, ", topic)
		for _, v := range depth.Bids {
			fmt.Printf("Bids: %v", v)
		}
		for _, v := range depth.Asks {
			fmt.Printf("Bids: %v", v)
		}
	})

	cli.Subscribe(topic)

	select {}
}

func TestSubscribeMBPRefreshDepth(t *testing.T) {
	cli := testNewMarketWsClient(t, GlobalWsBaseURL)

	topic, err := cli.GetMBPRefreshDepthTopic(&MBPDepthRefreshTopicParam{
		Symbol: "btcusdt",
		Level:  5,
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		depth, ok := e.(*types.MBPRefreshDepth)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, ", topic)
		for _, v := range depth.Bids {
			fmt.Printf("Bids: %v", v)
		}
		for _, v := range depth.Asks {
			fmt.Printf("Bids: %v", v)
		}
	})

	cli.Subscribe(topic)

	select {}
}
