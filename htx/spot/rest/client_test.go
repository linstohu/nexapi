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

package rest

import (
	"context"
	"os"
	"testing"

	"github.com/linstohu/nexapi/htx/spot/rest/types"
	"github.com/linstohu/nexapi/htx/utils"
	"github.com/stretchr/testify/assert"
)

func testNewSpotClient(t *testing.T) *SpotClient {
	cli, err := NewSpotClient(&SpotClientCfg{
		BaseURL: utils.SpotAWSBaseURL,
		Key:     os.Getenv("HTX_KEY"),
		Secret:  os.Getenv("HTX_SECRET"),
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create htx client, %s", err)
	}

	return cli
}

func TestGetSymbols(t *testing.T) {
	cli := testNewSpotClient(t)

	_, err := cli.GetSymbols(context.TODO(), types.GetSymbolsParam{})

	assert.Nil(t, err)
}

func TestGetMergedMarketTicker(t *testing.T) {
	cli := testNewSpotClient(t)

	_, err := cli.GetMergedMarketTicker(context.TODO(), types.GetMergedMarketTickerParam{
		Symbol: "btcusdt",
	})

	assert.Nil(t, err)
}

func TestGetAccountInfo(t *testing.T) {
	cli := testNewSpotClient(t)

	_, err := cli.GetAccountInfo(context.TODO())

	assert.Nil(t, err)
}

func TestGetAccountValuation(t *testing.T) {
	cli := testNewSpotClient(t)

	_, err := cli.GetAccountValuation(context.TODO(),
		types.GetAccountValuationParam{},
	)

	assert.Nil(t, err)
}

func TestNewOrder(t *testing.T) {
	cli := testNewSpotClient(t)

	_, err := cli.NewOrder(context.TODO(), types.NewOrderParam{
		AccountID: "",
		Symbol:    "usdcusdt",
		Type:      "buy-limit",
		Amount:    "12",
		Price:     "0.9",
	})

	assert.Nil(t, err)
}

func TestCancelOrder(t *testing.T) {
	cli := testNewSpotClient(t)

	_, err := cli.CancelOrder(context.TODO(), types.CancelOrderParams{
		OrderID: "",
	})

	assert.Nil(t, err)
}

func TestGetOpenOrders(t *testing.T) {
	cli := testNewSpotClient(t)

	_, err := cli.GetOpenOrders(context.TODO(), types.GetOpenOrdersParam{
		AccountID: "",
	})

	assert.Nil(t, err)
}

func TestCancelOrders(t *testing.T) {
	cli := testNewSpotClient(t)

	_, err := cli.CancelOrders(context.TODO(), types.CancelOrdersParam{
		AccountID: "",
	})

	assert.Nil(t, err)
}

func TestSearchMatchResults(t *testing.T) {
	cli := testNewSpotClient(t)

	_, err := cli.SearchMatchResults(context.TODO(), types.SearchMatchResultsParam{
		Symbol: "usdcusdt",
	})

	assert.Nil(t, err)
}
