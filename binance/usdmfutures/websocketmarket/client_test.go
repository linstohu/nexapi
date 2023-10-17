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
	"context"
	"fmt"
	"testing"
	"time"

	spottypes "github.com/linstohu/nexapi/binance/spot/websocketmarket/types"
	"github.com/linstohu/nexapi/binance/usdmfutures/websocketmarket/types"
	"github.com/stretchr/testify/assert"
)

func testNewMarketStreamClient(ctx context.Context, t *testing.T) *USDMarginedMarketStreamClient {
	cli, err := NewMarketStreamClient(ctx, &USDMarginedMarketStreamCfg{
		BaseURL: USDMarginedMarketStreamBaseURL,
		Debug:   false,
	})

	if err != nil {
		t.Fatalf("Could not create websocket client, %s", err)
	}

	return cli
}

func TestSubscribeAggTrade(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

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

	time.Sleep(5 * time.Second)

	cli.UnSubscribe([]string{topic})

	select {}
}

func TestSubscribeMarkPrice(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetMarkPriceTopic(&MarkPriceTopicParam{
		Symbol:      "btcusdt",
		UpdateSpeed: "1s",
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		trade, ok := e.(*types.MarkPrice)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, MarkPrice: %v, IndexPrice: %v, Time: %v\n",
			topic, trade.Symbol, trade.MarkPrice, trade.IndexPrice, trade.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeAllMarkPrice(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetAllMarketPriceTopic(&AllMarkPriceTopicParam{
		UpdateSpeed: "1s",
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		trades, ok := e.([]*types.MarkPrice)
		if !ok {
			return
		}

		for _, trade := range trades {
			fmt.Printf("Topic: %s, Symbol: %v, MarkPrice: %v, IndexPrice: %v, Time: %v\n",
				topic, trade.Symbol, trade.MarkPrice, trade.IndexPrice, trade.EventTime)
		}
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeKline(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetKlineTopic(&KlineTopicParam{
		Symbol:   "btcusdt",
		Interval: "1m",
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		kline, ok := e.(*spottypes.Kline)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, OpenPrice: %v, ClosePrice: %v, Time: %v\n",
			topic, kline.Symbol, kline.Kline.OpenPrice, kline.Kline.ClosePrice, kline.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeMiniTicker(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetMiniTickerTopic("btcusdt")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		ticker, ok := e.(*spottypes.MiniTicker)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, OpenPrice: %v, ClosePrice: %v, Time: %v\n",
			topic, ticker.Symbol, ticker.OpenPrice, ticker.ClosePrice, ticker.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeAllMiniTicker(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetAllMarketMiniTickersTopic()
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		tickers, ok := e.([]*spottypes.MiniTicker)
		if !ok {
			return
		}

		for _, ticker := range tickers {
			fmt.Printf("Topic: %s, Symbol: %v, OpenPrice: %v, ClosePrice: %v, Time: %v\n",
				topic, ticker.Symbol, ticker.OpenPrice, ticker.ClosePrice, ticker.EventTime)
		}
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeTicker(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetTickerTopic("btcusdt")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		ticker, ok := e.(*types.Ticker)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, OpenPrice: %v, LastPrice: %v, Time: %v\n",
			topic, ticker.Symbol, ticker.OpenPrice, ticker.LastPrice, ticker.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeAllTicker(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetAllMarketTickersTopic()
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		tickers, ok := e.([]*types.Ticker)
		if !ok {
			return
		}

		for _, ticker := range tickers {
			fmt.Printf("Topic: %s, Symbol: %v, OpenPrice: %v, LastPrice: %v, Time: %v\n",
				topic, ticker.Symbol, ticker.OpenPrice, ticker.LastPrice, ticker.EventTime)
		}
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeBookTicker(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetBookTickerTopic("btcusdt")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		book, ok := e.(*types.BookTicker)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, BestBidPrice: %v, BestBidQty: %v, BestAskPrice: %v, BestAskQty: %v\n",
			topic, book.Symbol, book.BestBidPrice, book.BestBidQty, book.BestAskPrice, book.BestAskQty)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeAllBookTickers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetAllBookTickersTopic()
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		book, ok := e.(*types.BookTicker)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, BestBidPrice: %v, BestBidQty: %v, BestAskPrice: %v, BestAskQty: %v\n",
			topic, book.Symbol, book.BestBidPrice, book.BestBidQty, book.BestAskPrice, book.BestAskQty)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeLiquidationOrder(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetLiquidationOrderTopic("btcusdt")
	assert.Nil(t, err)

	fmt.Println(topic)

	cli.AddListener(topic, func(e any) {
		order, ok := e.(*types.LiquidationOrder)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, Side: %v, Price: %v, Quanty: %v\n",
			topic, order.Order.Symbol, order.Order.Side, order.Order.Price, order.Order.FilledAccumulatedQuantity)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeAllLiquidationOrders(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetAllLiquidationOrdersTopic()
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		order, ok := e.(*types.LiquidationOrder)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, Side: %v, Price: %v, Quanty: %v\n",
			topic, order.Order.Symbol, order.Order.Side, order.Order.Price, order.Order.FilledAccumulatedQuantity)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeBookDepth(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetBookDepthTopic(&BookDepthTopicParam{
		Symbol:      "btcusdt",
		Level:       5,
		UpdateSpeed: "500ms",
	})
	assert.Nil(t, err)

	fmt.Println(topic)

	cli.AddListener(topic, func(e any) {
		book, ok := e.(*types.OrderbookDepth)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, FirstID: %v, FinalID: %v, Bids-Num: %v, Asks-Num: %v\n",
			topic, book.Symbol, book.FirstUpdateID, book.FinalUpdateID, len(book.Bids), len(book.Asks))
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeBookDiffDepth(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetBookDiffDepthTopic(&BookDiffDepthTopicParam{
		Symbol:      "btcusdt",
		UpdateSpeed: "500ms",
	})
	assert.Nil(t, err)

	fmt.Println(topic)

	cli.AddListener(topic, func(e any) {
		book, ok := e.(*types.OrderbookDepth)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, FirstID: %v, FinalID: %v, Bids-Num: %v, Asks-Num: %v\n",
			topic, book.Symbol, book.FirstUpdateID, book.FinalUpdateID, len(book.Bids), len(book.Asks))
	})

	cli.Subscribe([]string{topic})

	select {}
}
