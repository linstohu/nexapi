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

	"github.com/linstohu/nexapi/htx/usdm/rest/types"
	"github.com/linstohu/nexapi/htx/utils"
	"github.com/stretchr/testify/assert"
)

func testNewUsdmClient(t *testing.T) *UsdmClient {
	cli, err := NewUsdmClient(&UsdmClientCfg{
		BaseURL: utils.UsdmBaseURL,
		Key:     os.Getenv("HTX_KEY"),
		Secret:  os.Getenv("HTX_SECRET"),
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create htx client, %s", err)
	}

	return cli
}

func TestGetContractInfo(t *testing.T) {
	cli := testNewUsdmClient(t)

	_, err := cli.GetContractInfo(context.TODO(), types.GetContractInfoParam{})

	assert.Nil(t, err)
}

func TestGetAssetValuation(t *testing.T) {
	cli := testNewUsdmClient(t)

	_, err := cli.GetAssetValuation(context.TODO(), types.GetAssetValuationParam{
		ValuationAsset: "USDT",
	})

	assert.Nil(t, err)
}

func TestPlaceIsolatedOrder(t *testing.T) {
	cli := testNewUsdmClient(t)

	_, err := cli.PlaceIsolatedOrder(context.TODO(), types.PlaceIsolatedOrderParam{
		ContractCode:   "XRP-USDT",
		Price:          0.5,
		Volume:         1,
		Direction:      "buy",
		LeverRate:      1,
		OrderPriceType: "limit",
	})

	assert.Nil(t, err)
}

func TestPlaceCrossOrder(t *testing.T) {
	cli := testNewUsdmClient(t)

	_, err := cli.PlaceCrossOrder(context.TODO(), types.PlaceCrossOrderParam{
		ContractCode:   "XRP-USDT",
		Price:          0.5,
		Volume:         1,
		Direction:      "buy",
		LeverRate:      1,
		OrderPriceType: "limit",
	})

	assert.Nil(t, err)
}

func TestCancelIsolatedOrder(t *testing.T) {
	cli := testNewUsdmClient(t)

	_, err := cli.CancelIsolatedOrder(context.TODO(), types.CancelIsolatedOrderParam{
		OrderID:      "",
		ContractCode: "XRP-USDT",
	})

	assert.Nil(t, err)
}

func TestCancelCrossOrder(t *testing.T) {
	cli := testNewUsdmClient(t)

	_, err := cli.CancelCrossOrder(context.TODO(), types.CancelCrossOrderParam{
		OrderID: "",
	})

	assert.Nil(t, err)
}

func TestCancelAllIsolatedOrders(t *testing.T) {
	cli := testNewUsdmClient(t)

	_, err := cli.CancelAllIsolatedOrders(context.TODO(), types.CancelAllIsolatedOrdersParam{
		ContractCode: "XRP-USDT",
	})

	assert.Nil(t, err)
}

func TestCancelAllCrossOrders(t *testing.T) {
	cli := testNewUsdmClient(t)

	_, err := cli.CancelAllCrossOrders(context.TODO(), types.CancelAllCrossOrdersParam{
		ContractCode: "XRP-USDT",
	})

	assert.Nil(t, err)
}

func TestGetIsolatedOpenOrders(t *testing.T) {
	cli := testNewUsdmClient(t)

	_, err := cli.GetIsolatedOpenOrders(context.TODO(), types.GetIsolatedOpenOrdersParam{
		ContractCode: "XRP-USDT",
	})

	assert.Nil(t, err)
}

func TestGetCrossOpenOrders(t *testing.T) {
	cli := testNewUsdmClient(t)

	_, err := cli.GetCrossOpenOrders(context.TODO(), types.GetCrossOpenOrdersParam{
		ContractCode: "XRP-USDT",
	})

	assert.Nil(t, err)
}

func TestGetIsolatedHistoryMatchResults(t *testing.T) {
	cli := testNewUsdmClient(t)

	_, err := cli.GetIsolatedHistoryMatchResults(context.TODO(), types.GetIsolatedHistoryMatchResultsParam{
		TradeType: 0,
	})

	assert.Nil(t, err)
}

func TestGetCrossHistoryMatchResults(t *testing.T) {
	cli := testNewUsdmClient(t)

	_, err := cli.GetCrossHistoryMatchResults(context.TODO(), types.GetCrossHistoryMatchResultsParam{
		Contract: "XRP-USDT",
	})

	assert.Nil(t, err)
}
