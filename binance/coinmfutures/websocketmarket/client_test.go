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

	coinmws "github.com/linstohu/nexapi/binance/coinmfutures/websocketmarket"
	spottypes "github.com/linstohu/nexapi/binance/spot/websocketmarket/types"
	usdmtypes "github.com/linstohu/nexapi/binance/usdmfutures/websocketmarket/types"
	"github.com/stretchr/testify/assert"
)

func testNewMarketStreamClient(t *testing.T) *coinmws.CoinMarginedMarketStreamClient {
	cli, err := coinmws.NewMarketStreamClient(&coinmws.CoinMarginedMarketStreamCfg{
		Debug:         false,
		BaseURL:       coinmws.CoinMarginedMarketStreamBaseURL,
		AutoReconnect: true,
	})

	if err != nil {
		t.Fatalf("Could not create websocket client, %s", err)
	}

	return cli
}

func TestSubscribeAggTrade(t *testing.T) {
	cli := testNewMarketStreamClient(t)
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetAggTradeTopic("btcusd_perp")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		trade, ok := e.(*usdmtypes.AggregateTrade)
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
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetIndexPriceTopic(&coinmws.IndexPriceTopicParam{
		Pair:        "btcusd",
		UpdateSpeed: "1s",
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		indexPrice, ok := e.(*coinmws.IndexPrice)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Pair: %v, IndexPrice: %v, Time: %v\n",
			topic, indexPrice.Pair, indexPrice.IndexPrice, indexPrice.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeMarkPrice(t *testing.T) {
	cli := testNewMarketStreamClient(t)
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetMarketPriceTopic(&coinmws.MarkPriceTopicParam{
		Symbol:      "btcusd_perp",
		UpdateSpeed: "1s",
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		markprice, ok := e.(*coinmws.MarkPrice)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, MarkPrice: %v, SettlePrice: %v, Time: %v\n",
			topic, markprice.Symbol, markprice.MarkPrice, markprice.EstimatedSettlePrice, markprice.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribePairMarkPrice(t *testing.T) {
	cli := testNewMarketStreamClient(t)
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetPairMarketPriceTopic(&coinmws.PairMarkPriceTopicParam{
		Pair:        "btcusd",
		UpdateSpeed: "1s",
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		markprices, ok := e.([]*coinmws.MarkPrice)
		if !ok {
			return
		}

		for _, markprice := range markprices {
			fmt.Printf("Topic: %s, Symbol: %v, MarkPrice: %v, SettlePrice: %v, Time: %v\n",
				topic, markprice.Symbol, markprice.MarkPrice, markprice.EstimatedSettlePrice, markprice.EventTime)
		}

	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeKline(t *testing.T) {
	cli := testNewMarketStreamClient(t)
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetKlineTopic(&coinmws.KlineTopicParam{
		Symbol:   "btcusd_perp",
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
	cli := testNewMarketStreamClient(t)
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetMiniTickerTopic("btcusd_perp")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		ticker, ok := e.(*coinmws.MiniTicker)
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
	cli := testNewMarketStreamClient(t)
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetAllMarketMiniTickersTopic()
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		tickers, ok := e.([]*coinmws.MiniTicker)
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
	cli := testNewMarketStreamClient(t)
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetTickerTopic("btcusd_perp")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		ticker, ok := e.(*coinmws.Ticker)
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
	cli := testNewMarketStreamClient(t)
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetAllMarketTickersTopic()
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		tickers, ok := e.([]*coinmws.Ticker)
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
	cli := testNewMarketStreamClient(t)
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetBookTickerTopic("btcusd_perp")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		book, ok := e.(*coinmws.BookTicker)
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
	cli := testNewMarketStreamClient(t)
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetAllBookTickersTopic()
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		book, ok := e.(*coinmws.BookTicker)
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
	cli := testNewMarketStreamClient(t)
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetLiquidationOrderTopic("btcusd_perp")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		order, ok := e.(*coinmws.LiquidationOrder)
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
	cli := testNewMarketStreamClient(t)
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetAllLiquidationOrdersTopic()
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		order, ok := e.(*coinmws.LiquidationOrder)
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
	cli := testNewMarketStreamClient(t)
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetBookDepthTopic(&coinmws.BookDepthTopicParam{
		Symbol:      "btcusd_perp",
		Level:       5,
		UpdateSpeed: "500ms",
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		book, ok := e.(*coinmws.OrderbookDepth)
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
	cli := testNewMarketStreamClient(t)
	err := cli.Open()
	assert.Nil(t, err)

	topic, err := cli.GetBookDiffDepthTopic(&coinmws.BookDiffDepthTopicParam{
		Symbol:      "btcusd_perp",
		UpdateSpeed: "500ms",
	})
	assert.Nil(t, err)

	fmt.Println(topic)

	cli.AddListener(topic, func(e any) {
		book, ok := e.(*coinmws.OrderbookDepth)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, FirstID: %v, FinalID: %v, Bids-Num: %v, Asks-Num: %v\n",
			topic, book.Symbol, book.FirstUpdateID, book.FinalUpdateID, len(book.Bids), len(book.Asks))
	})

	cli.Subscribe([]string{topic})

	select {}
}
