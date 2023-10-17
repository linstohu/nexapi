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

package marketdata

import (
	"context"
	"testing"

	"github.com/linstohu/nexapi/mexc/spot/marketdata/types"
	spotutils "github.com/linstohu/nexapi/mexc/spot/utils"
	"github.com/stretchr/testify/assert"
)

func testNewSpotMarketDataClient(t *testing.T) *SpotMarketDataClient {
	cli, err := NewSpotMarketDataClient(&spotutils.SpotClientCfg{
		BaseURL: spotutils.BaseURL,
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create mexc client, %s", err)
	}

	return cli
}

func TestPing(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	err := cli.Ping(context.TODO())
	assert.Nil(t, err)
}

func TestGetServerTime(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetServerTime(context.TODO())
	assert.Nil(t, err)
}

func TestGetSymbols(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetSymbols(context.TODO())
	assert.Nil(t, err)
}

func TestGetExchangeInfo(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetExchangeInfo(context.TODO(), types.GetExchangeInfoParam{})
	assert.Nil(t, err)

	_, err = cli.GetExchangeInfo(context.TODO(), types.GetExchangeInfoParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)
}

func TestGetOrderbook(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetOrderbook(context.TODO(), types.GetOrderbookParams{
		Symbol: "BTCUSDT",
		Limit:  10,
	})
	assert.Nil(t, err)
}

func TestGetRecentTradeList(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetRecentTradeList(context.TODO(), types.GetTradeParams{
		Symbol: "BTCUSDT",
		Limit:  10,
	})
	assert.Nil(t, err)
}

func TestGetAggTrades(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetAggTrades(context.TODO(), types.GetAggTradesParam{
		Symbol: "BTCUSDT",
		Limit:  5,
	})
	assert.Nil(t, err)
}

func TestGetKlines(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetKlines(context.TODO(), types.GetKlineParam{
		Symbol:   "BTCUSDT",
		Interval: spotutils.Minute1,
		Limit:    1,
	})
	assert.Nil(t, err)
}

func TestGetAvgPrice(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetAvgPrice(context.TODO(), types.GetAvgPriceParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)
}

func TestGetTickerForSymbol(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetTickerForSymbol(context.TODO(), types.GetTickerForSymbolParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)
}

func TestGetTickerForAllSymbols(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetTickerForAllSymbols(context.TODO())
	assert.Nil(t, err)
}

func TestGetTickerPriceForSymbol(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetTickerPriceForSymbol(context.TODO(), types.GetTickerPriceForSymbolParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)
}

func TestGetTickerPriceForAllSymbols(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetTickerPriceForAllSymbols(context.TODO())
	assert.Nil(t, err)
}

func TestGetBookTickerForSymbol(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetBookTickerForSymbol(context.TODO(), types.GetBookTickerParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)
}

func TestGetBookTickerForAllSymbols(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetBookTickerForSymbols(context.TODO())
	assert.Nil(t, err)
}
