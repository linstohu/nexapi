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
	"net/http"
	"os"
	"testing"

	"github.com/linstohu/nexapi/woox/rest/types"
	"github.com/stretchr/testify/assert"
)

func testNewWooXRestPrivateClient(t *testing.T) *WooXRestClient {
	cli, err := NewWooXRestClient(&WooXRestClientCfg{
		BaseURL: TestNetBaseURL,
		Key:     os.Getenv("WOOX_KEY"),    // required
		Secret:  os.Getenv("WOOX_SECRET"), // required
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create woox client, %s", err)
	}

	return cli
}

func TestNormalizeV1RequestContent(t *testing.T) {
	req := types.HTTPRequest{
		Path:   "/v1/order",
		Method: http.MethodPost,
		Body: types.SendOrderReq{
			Symbol:        "SPOT_BTC_USDT",
			OrderType:     "LIMIT",
			OrderPrice:    9000,
			OrderQuantity: 0.11,
			Side:          "BUY",
		},
	}

	content, err := normalizeV1RequestContent(req)
	assert.Nil(t, err)
	assert.Equal(t, "order_price=9000&order_quantity=0.11&order_type=LIMIT&side=BUY&symbol=SPOT_BTC_USDT", content)
}

func TestSendOrder(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.SendOrder(context.TODO(), types.SendOrderReq{
		Symbol:        "SPOT_BTC_USDT",
		OrderType:     MarketOrderType,
		Side:          BUY,
		OrderQuantity: 0.1,
	})
	assert.Nil(t, err)
}

func TestGetOrders(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetOrders(context.TODO(), types.GetOrdersParam{})
	assert.Nil(t, err)
}

func TestGetTradeHistory(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetTradeHistory(context.TODO(), types.GetTradeHistoryParam{})
	assert.Nil(t, err)
}

func TestGetBalances(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetBalances(context.TODO())
	assert.Nil(t, err)
}

func TestGetAccountInfo(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetAccountInfo(context.TODO())
	assert.Nil(t, err)
}

func TestGetAssetHisotry(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetAssetHisotry(context.TODO(), types.GetAssetHisotryParam{})
	assert.Nil(t, err)
}

func TestGetSubAccounts(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetSubAccounts(context.TODO())
	assert.Nil(t, err)
}

func TestTransferAsset(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.TransferAsset(context.TODO(), types.TransferAssetParam{
		Token:     "USDT",
		Amount:    100,
		FromAppID: "",
		ToAppID:   "",
	})
	assert.Nil(t, err)
}

func TestUpdateAccountMode(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.UpdateAccountMode(context.TODO(), types.UpdateAccountModeParam{
		AccountMode: "FUTURES", // PURE_SPOT, MARGIN, FUTURES
	})
	assert.Nil(t, err)
}

func TestGetIPRestriction(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetIPRestriction(context.TODO())
	assert.Nil(t, err)
}

func TestGetOnePositionInfo(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetOnePositionInfo(context.TODO(), "PERP_BTC_USDT")
	assert.Nil(t, err)
}

func TestGetAllPositionInfo(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetAllPositionInfo(context.TODO())
	assert.Nil(t, err)
}
