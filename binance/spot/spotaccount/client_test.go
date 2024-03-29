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

package spotaccount

import (
	"context"
	"os"
	"testing"

	"github.com/linstohu/nexapi/binance/spot/spotaccount/types"
	spotutils "github.com/linstohu/nexapi/binance/spot/utils"
	"github.com/stretchr/testify/assert"
)

func testNewSpotAccountClient(t *testing.T) *SpotAccountClient {
	cli, err := NewSpotAccountClient(&SpotAccountClientCfg{
		BaseURL: spotutils.BaseURL,
		Key:     os.Getenv("BINANCE_KEY"),
		Secret:  os.Getenv("BINANCE_SECRET"),
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create binance client, %s", err)
	}

	return cli
}

func TestTestNewOrder(t *testing.T) {
	cli := testNewSpotAccountClient(t)

	err := cli.TestNewOrder(context.TODO(), types.NewOrderParam{
		Symbol:      "BTCUSDT",
		Side:        types.SideTypeBuy,
		Type:        types.Limit,
		TimeInForce: types.GTC,
		Price:       26000,
		Quantity:    0.001,
	})
	assert.Nil(t, err)
}

func TestNewOrder(t *testing.T) {
	cli := testNewSpotAccountClient(t)

	_, err := cli.NewOrder(context.TODO(), types.NewOrderParam{
		Symbol:      "BTCUSDT",
		Side:        types.SideTypeBuy,
		Type:        types.Limit,
		TimeInForce: types.GTC,
		Quantity:    0.001,
		Price:       25800,
	})
	assert.Nil(t, err)
}

func TestCancelOrder(t *testing.T) {
	cli := testNewSpotAccountClient(t)

	_, err := cli.CancelOrder(context.TODO(), types.CancelOrderParam{
		Symbol:  "BTCUSDT",
		OrderID: 1,
	})
	assert.Nil(t, err)
}

func TestCancelOrdersOnOneSymbol(t *testing.T) {
	cli := testNewSpotAccountClient(t)

	_, err := cli.CancelOrdersOnOneSymbol(context.TODO(), types.CancelOrdersOnOneSymbolParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)
}

func TestQueryOrder(t *testing.T) {
	cli := testNewSpotAccountClient(t)

	_, err := cli.QueryOrder(context.TODO(), types.QueryOrderParam{
		Symbol:  "BTCUSDT",
		OrderID: 1,
	})
	assert.Nil(t, err)
}

func TestGetOpenOrders(t *testing.T) {
	cli := testNewSpotAccountClient(t)

	_, err := cli.GetOpenOrders(context.TODO(), types.GetOpenOrdersParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)
}

func TestGetAllOrders(t *testing.T) {
	cli := testNewSpotAccountClient(t)

	_, err := cli.GetAllOrders(context.TODO(), types.GetAllOrdersParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)
}

func TestGetAccountInfo(t *testing.T) {
	cli := testNewSpotAccountClient(t)

	_, err := cli.GetAccountInfo(context.TODO())
	assert.Nil(t, err)
}

func TestGetTradeList(t *testing.T) {
	cli := testNewSpotAccountClient(t)

	_, err := cli.GetTradeList(context.TODO(), types.GetTradesParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)
}
