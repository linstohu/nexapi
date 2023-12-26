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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testNewMarketStreamClient(t *testing.T) *OptionsMarketStreamClient {
	cli, err := NewMarketStreamClient(&OptionsMarketStreamCfg{
		Debug:         true,
		BaseURL:       OptionsMarketStreamBaseURL,
		AutoReconnect: true,
	})

	if err != nil {
		t.Fatalf("Could not create websocket client, %s", err)
	}

	return cli
}

func TestSubscribeTrade(t *testing.T) {
	cli := testNewMarketStreamClient(t)

	topic, err := cli.GetTradeTopic("ETH")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		trade, ok := e.(*Trade)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, Price: %v, Quantity: %v, Time: %v\n",
			topic, trade.Symbol, trade.Price, trade.Quantity, trade.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeIndexPrice(t *testing.T) {
	cli := testNewMarketStreamClient(t)

	topic, err := cli.GetIndexPriceTopic("ETHUSDT")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		indexPrice, ok := e.(*IndexPrice)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, IndexPrice: %v, Time: %v\n",
			topic, indexPrice.Symbol, indexPrice.IndexPrice, indexPrice.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeMarkPrice(t *testing.T) {
	cli := testNewMarketStreamClient(t)

	topic, err := cli.GetMarkPriceTopic("ETH")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		s, ok := e.([]*MarkPrice)
		if !ok {
			return
		}

		for _, markprice := range s {
			fmt.Printf("Topic: %s, Symbol: %v, MarkPrice: %v, Time: %v\n",
				topic, markprice.Symbol, markprice.MarkPrice, markprice.EventTime)
		}
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeKline(t *testing.T) {
	cli := testNewMarketStreamClient(t)

	topic, err := cli.GetKlineTopic(&KlineTopicParam{
		Symbol:   "ETH-230525-1825-C",
		Interval: "1m",
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		kline, ok := e.(*Kline)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, OpenPrice: %v, ClosePrice: %v, Time: %v\n",
			topic, kline.Symbol, kline.Kline.OpenPrice, kline.Kline.ClosePrice, kline.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribe24HourTicker(t *testing.T) {
	cli := testNewMarketStreamClient(t)

	topic, err := cli.Get24HourTickerTopic("ETH-230525-1825-C")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		ticker, ok := e.(*Ticker)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, OpenPrice: %v, LowPrice: %v, Time: %v\n",
			topic, ticker.Symbol, ticker.OpeningPrice, ticker.LowPrice, ticker.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeUnderlying24HourTicker(t *testing.T) {
	cli := testNewMarketStreamClient(t)

	topic, err := cli.Get24HourTickerByUnderlyingAndexpirationTopic("ETH", "230525")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		tickers, ok := e.([]*Ticker)
		if !ok {
			return
		}

		for _, ticker := range tickers {
			fmt.Printf("Topic: %s, Symbol: %v, OpenPrice: %v, LowPrice: %v, Time: %v\n",
				topic, ticker.Symbol, ticker.OpeningPrice, ticker.LowPrice, ticker.EventTime)
		}
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeOpenInterest(t *testing.T) {
	cli := testNewMarketStreamClient(t)

	topic, err := cli.GetOpenInterestTopic("BTC", "230525")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		s, ok := e.([]*OpenInterest)
		if !ok {
			return
		}

		for _, openInterest := range s {
			fmt.Printf("Topic: %s, Symbol: %v, OpenInterestInContracts: %v, OpenInterestInUSD: %v, Time: %v\n",
				topic, openInterest.Symbol, openInterest.OpenInterestInContracts, openInterest.OpenInterestInUSDT, openInterest.EventTime)
		}
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeOrderbook(t *testing.T) {
	cli := testNewMarketStreamClient(t)

	topic, err := cli.GetBookDepthTopic(&BookDepthTopicParam{
		Symbol:      "BTC-230602-25000-P",
		Level:       10,
		UpdateSpeed: "500ms",
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		book, ok := e.(*OrderbookDepth)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, UpdateID: %v, BidsNum: %v, AsksNum: %v\n",
			topic, book.Symbol, book.UpdateID, len(book.Bids), len(book.Asks))
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeDiffOrderbook(t *testing.T) {
	cli := testNewMarketStreamClient(t)

	topic, err := cli.GetBookDiffDepthTopic("BTC-230602-25000-P")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		book, ok := e.(*OrderbookDepth)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, UpdateID: %v, BidsNum: %v, AsksNum: %v\n",
			topic, book.Symbol, book.UpdateID, len(book.Bids), len(book.Asks))
	})

	cli.Subscribe([]string{topic})

	select {}
}
