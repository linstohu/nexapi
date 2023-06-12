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

### Intro

### Binance Example

#### Example 1: REST API

```go
package main

import (
	"context"
	"fmt"
	"log"

	bnUsdmMarket "github.com/linstohu/nexapi/binance/usdmfutures/marketdata"
	umtypes "github.com/linstohu/nexapi/binance/usdmfutures/marketdata/types"
	umutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
)

func main() {
	cli, err := bnUsdmMarket.NewUSDMFuturesMarketDataClient(&umutils.USDMarginedClientCfg{
		Debug:   false,
		Logger:  log.Default(),
		BaseURL: umutils.USDMarginedBaseURL,
	})

	if err != nil {
		panic(err)
	}

	trades, err := cli.GetRecentTradeList(context.TODO(), umtypes.GetTradeParams{
		Symbol: "BTCUSDT",
		Limit:  10,
	})
	if err != nil {
		panic(err)
	}

	for i, v := range trades {
		fmt.Printf("Index-%v, %+v\n", i, v)
	}
}
```

#### Example 2: Websocket API

```go

package main

import (
	"context"
	"fmt"
	"log"

	bnUsdmWsMarket "github.com/linstohu/nexapi/binance/usdmfutures/websocketmarket"
	bnUsdmWsTypes "github.com/linstohu/nexapi/binance/usdmfutures/websocketmarket/types"
)

func main() {
	cli, err := bnUsdmWsMarket.NewMarketStreamClient(context.TODO(), &bnUsdmWsMarket.USDMarginedMarketStreamCfg{
		Debug:   false,
		Logger:  log.Default(),
		BaseURL: bnUsdmWsMarket.USDMarginedMarketStreamBaseURL,
	})

	if err != nil {
		panic(err)
	}

	topic, err := cli.GetAggTradeTopic("btcusdt")
	if err != nil {
		panic(err)
	}

	cli.AddListener(topic, func(e any) {
		trade, ok := e.(*bnUsdmWsTypes.AggregateTrade)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, Price: %v, Quantity: %v, Time: %v\n",
			topic, trade.Symbol, trade.Price, trade.Quantity, trade.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

```

### Gate.io Example

## ⭐ Give a Star!

If you like or are using this project to learn or start your solution, please give it a star. Thanks!