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

package account

import (
	"context"
	"os"
	"testing"

	"github.com/linstohu/nexapi/binance/europeanoptions/account/types"
	eoutils "github.com/linstohu/nexapi/binance/europeanoptions/utils"
	"github.com/stretchr/testify/assert"
)

func testNewAccountClient(t *testing.T) *OptionsAccountClient {
	cli, err := NewOptionsAccountClient(&eoutils.OptionsClientCfg{
		BaseURL: eoutils.OptionsBaseURL,
		Key:     os.Getenv("BINANCE_KEY"),
		Secret:  os.Getenv("BINANCE_SECRET"),
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create binance client, %s", err)
	}

	return cli
}

func TestGetAccountInfo(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetAccountInfo(context.TODO())
	assert.Nil(t, err)
}

func TestNewOrder(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.NewOrder(context.TODO(), types.NewOrderParam{
		Symbol:      "ETH-230609-2100-C",
		Side:        eoutils.BuySide,
		Type:        eoutils.Limit,
		Quantity:    1,
		Price:       4.5,
		TimeInForce: eoutils.GTC,
	})
	assert.Nil(t, err)
}

func TestGetSingleOrder(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetSingleOrder(context.TODO(), types.GetSingleOrderParam{
		Symbol:  "ETH-230609-2100-C",
		OrderID: 0,
	})
	assert.Nil(t, err)
}

func TestCancelOrder(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.CancelOrder(context.TODO(), types.CancelOrderParam{
		Symbol:  "ETH-230609-2100-C",
		OrderID: 0,
	})
	assert.Nil(t, err)
}

func TestCancelAllOrdersBySymbol(t *testing.T) {
	cli := testNewAccountClient(t)

	err := cli.CancelAllOrdersBySymbol(context.TODO(), types.CancelAllOrdersParam{
		Symbol: "ETH-230609-2100-C",
	})
	assert.Nil(t, err)
}

func TestCancelAllOrdersByUnderlying(t *testing.T) {
	cli := testNewAccountClient(t)

	err := cli.CancelAllOrdersByUnderlying(context.TODO(), types.CancelAllOrdersByUnderlyingParam{
		Underlying: "ETHUSDT",
	})
	assert.Nil(t, err)
}

func TestGetOpenOrders(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetOpenOrders(context.TODO(), types.GetCurrentOpenOrdersParam{
		Symbol: "ETH-230609-2100-C",
	})
	assert.Nil(t, err)
}

func TestGetOrderHistory(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetOrderHistory(context.TODO(), types.GetOrderHistoryParam{
		Symbol: "ETH-230609-2100-C",
	})
	assert.Nil(t, err)
}

func TestGetPositionInfo(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetPositionInfo(context.TODO(), types.GetPositionInfoParam{
		Symbol: "ETH-230609-2100-C",
	})
	assert.Nil(t, err)
}

func TestGetTradeList(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetTradeList(context.TODO(), types.GetTradeListParam{})
	assert.Nil(t, err)
}

func TestGetExerciseRecord(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetExerciseRecord(context.TODO(), types.GetExerciseRecordParam{})
	assert.Nil(t, err)
}

func TestGetFundingFlow(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetFundingFlow(context.TODO(), types.GetFundingFlowParam{
		Currency: "USDT",
	})
	assert.Nil(t, err)
}
