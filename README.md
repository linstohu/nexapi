# NexAPI – N EXchanges API Library

<p align="center">
<img align="center" width="150px" src="./docs/imgs/nexapi.svg">
</p>

NexAPI is a GO library that integrates official APIs from many well-known cryptocurrency exchanges.

<div align=center>

[![Go Report Card](https://goreportcard.com/badge/github.com/linstohu/nexapi)](https://goreportcard.com/report/github.com/linstohu/nexapi)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/linstohu/nexapi)](https://github.com/linstohu/nexapi/blob/main/go.mod)
[![Join Linsto Telegram](https://img.shields.io/badge/telegram-linsto-brightgreen?logo=telegram)](https://t.me/linstohu)

</div>

## 🤷‍ What is NexAPI?

The NexAPI library is a collection of APIs from many well-known cryptocurrency exchanges. It provides quick access to market data for storage, analysis, visualization, indicator development, algorithmic trading, strategy backtesting, bot programming, and related software engineering.

It is intended to be used by coders, developers, technically-skilled traders, data-scientists and financial analysts for building trading algorithms.

## 🔎 Usage

### Binance Example

#### Example 1: REST API

```go

package main

import (
	"context"
	"encoding/json"
	"fmt"

	bnspotmd "github.com/linstohu/nexapi/binance/spot/marketdata"
	bnspottypes "github.com/linstohu/nexapi/binance/spot/marketdata/types"
	bnspotutils "github.com/linstohu/nexapi/binance/spot/utils"
)

func main() {
	cli, err := bnspotmd.NewSpotMarketDataClient(&bnspotutils.SpotClientCfg{
		Debug:   true,
		BaseURL: bnspotutils.BaseURL,
	})
	if err != nil {
		panic(err)
	}

	orderbook, err := cli.GetOrderbook(context.TODO(), bnspottypes.GetOrderbookParams{
		Symbol: "BTCUSDT",
		Limit:  5,
	})
	if err != nil {
		panic(err)
	}

	limit := orderbook.Http.ApiRes.Header.Get("X-Mbx-Used-Weight-1m")

	fmt.Printf("Current used request weight: %v\n", limit)

	bytes, err := json.MarshalIndent(orderbook.Body, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

```

#### Example 2: Websocket

```go

package main

import (
	"fmt"
	"time"

	spotws "github.com/linstohu/nexapi/binance/spot/websocketmarket"
	"github.com/linstohu/nexapi/binance/spot/websocketmarket/types"
)

func main() {
	cli, err := spotws.NewSpotMarketStreamClient(&spotws.SpotMarketStreamCfg{
		Debug:         true,
		BaseURL:       spotws.SpotMarketStreamBaseURL,
		AutoReconnect: true,
	})
	if err != nil {
		panic(err)
	}

	err = cli.Open()
	if err != nil {
		panic(err)
	}

	topic, err := cli.GetAggTradeTopic("btcusdt")
	if err != nil {
		panic(err)
	}

	cli.AddListener(topic, func(e any) {
		trade, ok := e.(*types.AggregateTrade)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, Price: %v, Quantity: %v, Time: %v\n",
			topic, trade.Symbol, trade.Price, trade.Quantity, trade.EventTime)
	})

	cli.Subscribe([]string{topic})

	time.Sleep(20 * time.Second)

	cli.Close()

	select {}
}

```

## ⭐ Give a Star!

If you like or are using this project to learn or start your solution, please give it a star. Thanks!
