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

	"github.com/linstohu/nexapi/binance/coinmfutures/marketdata/types"
	cmutils "github.com/linstohu/nexapi/binance/coinmfutures/utils"
	usdmtypes "github.com/linstohu/nexapi/binance/usdmfutures/marketdata/types"
	usdmutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	"github.com/stretchr/testify/assert"
)

func testCoinMFuturesMarketDataClient(t *testing.T) *CoinMFuturesMarketDataClient {
	cli, err := NewCoinMFuturesMarketDataClient(&cmutils.CoinMarginedClientCfg{
		BaseURL: cmutils.CoinMarginedBaseURL,
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create binance client, %s", err)
	}

	return cli
}

func TestPing(t *testing.T) {
	cli := testCoinMFuturesMarketDataClient(t)

	err := cli.Ping(context.TODO())
	assert.Nil(t, err)
}

func TestGetServerTime(t *testing.T) {
	cli := testCoinMFuturesMarketDataClient(t)

	_, err := cli.GetServerTime(context.TODO())
	assert.Nil(t, err)
}

func TestGetExchangeInfo(t *testing.T) {
	cli := testCoinMFuturesMarketDataClient(t)

	_, err := cli.GetExchangeInfo(context.TODO())
	assert.Nil(t, err)
}

func TestGetOrderbook(t *testing.T) {
	cli := testCoinMFuturesMarketDataClient(t)

	_, err := cli.GetOrderbook(context.TODO(), types.GetOrderbookParams{
		Symbol: "ETHUSD_PERP",
		Limit:  10,
	})
	assert.Nil(t, err)
}

func TestGetRecentTradeList(t *testing.T) {
	cli := testCoinMFuturesMarketDataClient(t)

	_, err := cli.GetRecentTradeList(context.TODO(), types.GetTradeParams{
		Symbol: "ETHUSD_PERP",
		Limit:  10,
	})
	assert.Nil(t, err)
}

func TestGetAggTrades(t *testing.T) {
	cli := testCoinMFuturesMarketDataClient(t)

	_, err := cli.GetAggTrades(context.TODO(), usdmtypes.GetAggTradesParam{
		Symbol: "ETHUSD_PERP",
		Limit:  5,
	})
	assert.Nil(t, err)
}

func TestGetMarkPrice(t *testing.T) {
	cli := testCoinMFuturesMarketDataClient(t)

	resp, err := cli.GetMarkPrice(context.TODO(), types.GetMarkPriceParam{
		Symbol: "ETHUSD_PERP",
	})
	assert.Nil(t, err)

	assert.Equal(t, 1, len(resp.Body))

	_, err = cli.GetMarkPrice(context.TODO(), types.GetMarkPriceParam{})
	assert.Nil(t, err)
}

func TestGetFundingRateHistory(t *testing.T) {
	cli := testCoinMFuturesMarketDataClient(t)

	_, err := cli.GetFundingRateHistory(context.TODO(), types.GetFundingRateParam{
		Symbol: "ETHUSD_PERP",
		Limit:  2,
	})
	assert.Nil(t, err)
}

func TestGetKlines(t *testing.T) {
	cli := testCoinMFuturesMarketDataClient(t)

	_, err := cli.GetKlines(context.TODO(), usdmtypes.GetKlineParam{
		Symbol:   "ETHUSD_PERP",
		Interval: usdmutils.Minute1,
		Limit:    2,
	})
	assert.Nil(t, err)
}

func TestGetTickerPrice(t *testing.T) {
	cli := testCoinMFuturesMarketDataClient(t)

	_, err := cli.GetTickerPrice(context.TODO(), types.GetPriceTickerParam{
		Symbol: "ETHUSD_PERP",
	})
	assert.Nil(t, err)
}

func TestGetBookTicker(t *testing.T) {
	cli := testCoinMFuturesMarketDataClient(t)

	_, err := cli.GetBookTicker(context.TODO(), types.GetBookTickerParam{
		Symbol: "ETHUSD_PERP",
	})
	assert.Nil(t, err)
}

func TestGetOpenInterest(t *testing.T) {
	cli := testCoinMFuturesMarketDataClient(t)

	_, err := cli.GetOpenInterest(context.TODO(), types.GetOpenInterestParam{
		Symbol: "ETHUSD_PERP",
	})
	assert.Nil(t, err)
}

func TestGetOpenInterestHistory(t *testing.T) {
	cli := testCoinMFuturesMarketDataClient(t)

	_, err := cli.GetOpenInterestHistory(context.TODO(), types.GetOpenInterestHistParam{
		Pair:         "ETHUSD",
		ContractType: "CURRENT_QUARTER",
		Period:       "5m",
		Limit:        2,
	})
	assert.Nil(t, err)
}
