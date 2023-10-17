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

	"github.com/linstohu/nexapi/binance/europeanoptions/marketdata/types"
	eoutils "github.com/linstohu/nexapi/binance/europeanoptions/utils"
	usdmtypes "github.com/linstohu/nexapi/binance/usdmfutures/marketdata/types"
	usdmutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	"github.com/stretchr/testify/assert"
)

func testOptionsMarketDataClient(t *testing.T) *OptionsMarketDataClient {
	cli, err := NewOptionsMarketDataClient(&eoutils.OptionsClientCfg{
		BaseURL: eoutils.OptionsBaseURL,
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create binance client, %s", err)
	}

	return cli
}

func TestPing(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	err := cli.Ping(context.TODO())
	assert.Nil(t, err)
}

func TestGetServerTime(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetServerTime(context.TODO())
	assert.Nil(t, err)
}

func TestGetExchangeInfo(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetExchangeInfo(context.TODO())
	assert.Nil(t, err)
}

func TestGetOrderbook(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetOrderbook(context.TODO(), types.GetOrderbookParams{
		Symbol: "BTC-230630-25000-P",
		Limit:  10,
	})
	assert.Nil(t, err)
}

func TestGetRecentTradesList(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetRecentTradesList(context.TODO(), types.GetTradeParams{
		Symbol: "BTC-230630-25000-P",
		Limit:  10,
	})
	assert.Nil(t, err)
}

func TestGetKlines(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetKlines(context.TODO(), usdmtypes.GetKlineParam{
		Symbol:   "BTC-230630-25000-P",
		Interval: usdmutils.Minute1,
		Limit:    2,
	})
	assert.Nil(t, err)
}

func TestGetMarkPrice(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetMarkPrice(context.TODO(), types.GetMarkPriceParam{
		Symbol: "BTC-230630-25000-P",
	})
	assert.Nil(t, err)
}

func TestGetTickerPrice(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetTickerPrice(context.TODO(), types.GetTickerPriceParam{
		Symbol: "BTC-230630-25000-P",
	})
	assert.Nil(t, err)
}

func TestGetUnderlyingIndexPrice(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetUnderlyingIndexPrice(context.TODO(), types.GetUnderlyingIndexPriceParams{
		Underlying: "BTCUSDT",
	})
	assert.Nil(t, err)
}

func TestGetOpenInterest(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetOpenInterest(context.TODO(), types.GetOpenInterestParam{
		UnderlyingAsset: "BTC",
		Expiration:      "230630",
	})
	assert.Nil(t, err)
}
