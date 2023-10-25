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

package wallet

import (
	"context"
	"os"
	"testing"

	spotutils "github.com/linstohu/nexapi/binance/spot/utils"
	"github.com/linstohu/nexapi/binance/spot/wallet/types"
	"github.com/stretchr/testify/assert"
)

func testNewSpotWalletClient(t *testing.T) *SpotWalletClient {
	cli, err := NewSpotWalletClient(&SpotWalletClientCfg{
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

func TestGetSystemStatus(t *testing.T) {
	cli := testNewSpotWalletClient(t)

	_, err := cli.GetSystemStatus(context.TODO())
	assert.Nil(t, err)
}

func TestGetAllCoinsInfo(t *testing.T) {
	cli := testNewSpotWalletClient(t)

	_, err := cli.GetAllCoinsInfo(context.TODO())
	assert.Nil(t, err)
}

func TestGetAssetDetail(t *testing.T) {
	cli := testNewSpotWalletClient(t)

	_, err := cli.GetAssetDetail(context.TODO(), types.GetAssetDetailParam{})
	assert.Nil(t, err)
}

func TestGetTradeFee(t *testing.T) {
	cli := testNewSpotWalletClient(t)

	_, err := cli.GetTradeFee(context.TODO(), types.GetTradeFeeParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)
}

func TestUniversalTransfer(t *testing.T) {
	cli := testNewSpotWalletClient(t)

	_, err := cli.UniversalTransfer(context.TODO(), types.UniversalTransferParam{
		Type:   types.MAIN_UMFUTURE,
		Asset:  "USDT",
		Amount: 10,
	})
	assert.Nil(t, err)
}

func TestGetUniversalTransferHistory(t *testing.T) {
	cli := testNewSpotWalletClient(t)

	_, err := cli.GetUniversalTransferHistory(context.TODO(), types.GetUniversalTransferHistoryParam{
		Type: types.MAIN_UMFUTURE,
	})
	assert.Nil(t, err)
}

func TestGetFundingAsset(t *testing.T) {
	cli := testNewSpotWalletClient(t)

	_, err := cli.GetFundingAsset(context.TODO(), types.GetFundingAssetParam{})
	assert.Nil(t, err)
}

func TestGetUserAsset(t *testing.T) {
	cli := testNewSpotWalletClient(t)

	_, err := cli.GetUserAsset(context.TODO(), types.GetUserAssetParam{})
	assert.Nil(t, err)
}

func TestGetApiRestrictions(t *testing.T) {
	cli := testNewSpotWalletClient(t)

	_, err := cli.GetApiRestrictions(context.TODO())
	assert.Nil(t, err)
}

func TestGetWalletBalance(t *testing.T) {
	cli := testNewSpotWalletClient(t)

	_, err := cli.GetWalletBalance(context.TODO())
	assert.Nil(t, err)
}

func TestGetEarnAccount(t *testing.T) {
	cli := testNewSpotWalletClient(t)

	_, err := cli.GetEarnAccount(context.TODO())
	assert.Nil(t, err)
}
