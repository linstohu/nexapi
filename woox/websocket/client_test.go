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

package websocket

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/linstohu/nexapi/woox/websocket/types"
	"github.com/stretchr/testify/assert"
)

func testNewWooXWebsocketClient(t *testing.T) *WooXWebsocketClient {
	cli, err := NewWooXWebsocketClient(&WooXWebsocketCfg{
		Debug:         true,
		BaseURL:       TestNetPublicBaseURL,
		AutoReconnect: true,
		Key:           os.Getenv("WOOX_KEY"),
		Secret:        os.Getenv("WOOX_SECRET"),
		ApplicationID: os.Getenv("WOOX_APP_ID"), // required
	})

	if err != nil {
		t.Fatalf("Could not create websocket client, %s", err)
	}

	return cli
}

func TestWooXWebsocketClientConnection(t *testing.T) {
	cli := testNewWooXWebsocketClient(t)

	err := cli.Subscribe([]string{"SPOT_WOO_USDT@orderbook"})
	assert.Nil(t, err)

	time.Sleep(1 * time.Minute)
}

func TestSubscribeOrderbook(t *testing.T) {
	cli := testNewWooXWebsocketClient(t)

	topic, err := cli.GetOrderbookTopic("PERP_BTC_USDT")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		book, ok := e.(*types.Orderbook)
		if !ok {
			return
		}

		if len(book.Data.Asks) == 0 || len(book.Data.Bids) == 0 {
			return
		}

		fmt.Printf("获取到新的Orderbook, 币对: %v, 时间戳: %v, 买单数量: %v, 卖单数量: %v\n",
			book.Data.Symbol, book.Ts, len(book.Data.Bids), len(book.Data.Asks))
	})

	cli.Subscribe([]string{topic})

	time.Sleep(3 * time.Second)

	cli.UnSubscribe([]string{topic})

	time.Sleep(5 * time.Second)

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeTrade(t *testing.T) {
	cli := testNewWooXWebsocketClient(t)

	topic, err := cli.GetTradeTopic("PERP_BTC_USDT")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		trade, ok := e.(*types.Trade)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Ts: %v, Symbol: %v, Price: %v, Size: %v, Side: %v\n",
			trade.Topic, trade.Ts, trade.Data.Symbol, trade.Data.Price, trade.Data.Size, trade.Data.Side)
	})

	cli.Subscribe([]string{topic})

	time.Sleep(5 * time.Second)

	cli.UnSubscribe([]string{topic})

	select {}
}

func TestSubscribeTickerForSymbol(t *testing.T) {
	cli := testNewWooXWebsocketClient(t)

	topic, err := cli.GetTickerTopic("PERP_BTC_USDT")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		ticker, ok := e.(*types.Ticker24H)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Open: %v, Close: %v, High: %v, Low: %v\n",
			ticker.Topic, ticker.Data.Open, ticker.Data.Close, ticker.Data.High, ticker.Data.Low)
	})

	cli.Subscribe([]string{topic})

	time.Sleep(5 * time.Second)

	cli.UnSubscribe([]string{topic})

	select {}
}

func TestSubscribeTickers(t *testing.T) {
	cli := testNewWooXWebsocketClient(t)

	topic, err := cli.GetAllTickersTopic()
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		tickers, ok := e.(*types.Tickers)
		if !ok {
			return
		}

		for _, ticker := range tickers.Data {
			fmt.Printf("Symbol: %s, Open: %v, Close: %v, High: %v, Low: %v\n",
				ticker.Symbol, ticker.Open, ticker.Close, ticker.High, ticker.Low)
		}
	})

	cli.Subscribe([]string{topic})

	time.Sleep(5 * time.Second)

	cli.UnSubscribe([]string{topic})

	select {}
}

func TestSubscribeBBOForSymbol(t *testing.T) {
	cli := testNewWooXWebsocketClient(t)

	topic, err := cli.GetBboTopic("PERP_BTC_USDT")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		bbo, ok := e.(*types.BBO)
		if !ok {
			return
		}

		fmt.Printf("Symbol: %s, Ask: %v, AskSize: %v, Bid: %v, BidSize: %v\n",
			bbo.Data.Symbol, bbo.Data.Ask, bbo.Data.AskSize, bbo.Data.Bid, bbo.Data.BidSize)
	})

	cli.Subscribe([]string{topic})

	time.Sleep(5 * time.Second)

	cli.UnSubscribe([]string{topic})

	select {}
}

func TestSubscribeBBOs(t *testing.T) {
	cli := testNewWooXWebsocketClient(t)

	topic, err := cli.GetAllBbosTopic()
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		bbos, ok := e.(*types.AllBBO)
		if !ok {
			return
		}

		for _, bbo := range bbos.Data {
			fmt.Printf("Symbol: %s, Ask: %v, AskSize: %v, Bid: %v, BidSize: %v\n",
				bbo.Symbol, bbo.Ask, bbo.AskSize, bbo.Bid, bbo.BidSize)
		}
	})

	cli.Subscribe([]string{topic})

	time.Sleep(5 * time.Second)

	cli.UnSubscribe([]string{topic})

	select {}
}

func TestSubscribeKline(t *testing.T) {
	cli := testNewWooXWebsocketClient(t)

	topic, err := cli.GetKlineTopic(&KlineTopicParam{
		Symbol: "PERP_BTC_USDT",
		Time:   "1m",
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		kline, ok := e.(*types.Kline)
		if !ok {
			return
		}

		fmt.Printf("Symbol: %s, Type: %v, Open: %v, Close: %v\n",
			kline.Data.Symbol, kline.Data.Type, kline.Data.Open, kline.Data.Close)
	})

	cli.Subscribe([]string{topic})

	time.Sleep(5 * time.Second)

	cli.UnSubscribe([]string{topic})

	select {}
}

func TestSubscribeIndexPrice(t *testing.T) {
	cli := testNewWooXWebsocketClient(t)

	topic, err := cli.GetIndexPriceTopic("SPOT_ETH_USDT")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		indexPrice, ok := e.(*types.IndexPrice)
		if !ok {
			return
		}

		fmt.Printf("Symbol: %s, Price: %v\n", indexPrice.Data.Symbol, indexPrice.Data.Price)
	})

	cli.Subscribe([]string{topic})

	time.Sleep(5 * time.Second)

	cli.UnSubscribe([]string{topic})

	select {}
}

func TestSubscribeMarkPrice(t *testing.T) {
	cli := testNewWooXWebsocketClient(t)

	topic, err := cli.GetMarkPriceTopic("PERP_BTC_USDT")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		mp, ok := e.(*types.MarkPrice)
		if !ok {
			return
		}

		fmt.Printf("Symbol: %s, Price: %v\n", mp.Data.Symbol, mp.Data.Price)
	})

	cli.Subscribe([]string{topic})

	time.Sleep(5 * time.Second)

	cli.UnSubscribe([]string{topic})

	select {}
}

func TestSubscribeAllMarkPrice(t *testing.T) {
	cli := testNewWooXWebsocketClient(t)

	topic, err := cli.GetMarkPricesTopic()
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		s, ok := e.(*types.MarkPrices)
		if !ok {
			return
		}

		for _, mp := range s.Data {
			fmt.Printf("Symbol: %s, Price: %v\n", mp.Symbol, mp.Price)
		}
	})

	cli.Subscribe([]string{topic})

	time.Sleep(5 * time.Second)

	cli.UnSubscribe([]string{topic})

	select {}
}

func TestSubscribeOpenInterest(t *testing.T) {
	cli := testNewWooXWebsocketClient(t)

	topic, err := cli.GetOpenInterestTopic("PERP_BTC_USDT")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		oi, ok := e.(*types.OpenInterest)
		if !ok {
			return
		}

		fmt.Printf("Symbol: %s, OpenInterest: %v\n", oi.Data.Symbol, oi.Data.OpenInterest)
	})

	cli.Subscribe([]string{topic})

	time.Sleep(5 * time.Second)

	cli.UnSubscribe([]string{topic})

	select {}
}

func TestSubscribeEstFundingRate(t *testing.T) {
	cli := testNewWooXWebsocketClient(t)

	topic, err := cli.GetEstFundingRateTopic("PERP_BTC_USDT")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		rate, ok := e.(*types.EstFundingRate)
		if !ok {
			return
		}

		fmt.Printf("Symbol: %s, FundingRate: %v, TS: %v\n", rate.Data.Symbol, rate.Data.FundingRate, rate.Data.FundingTs)
	})

	cli.Subscribe([]string{topic})

	time.Sleep(5 * time.Second)

	cli.UnSubscribe([]string{topic})

	select {}
}
