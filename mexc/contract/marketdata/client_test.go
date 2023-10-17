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

	"github.com/linstohu/nexapi/mexc/contract/marketdata/types"
	ctutils "github.com/linstohu/nexapi/mexc/contract/utils"
	"github.com/stretchr/testify/assert"
)

func testNewContractMarketDataClient(t *testing.T) *ContractMarketDataClient {
	cli, err := NewContractMarketDataClient(&ctutils.ContractClientCfg{
		BaseURL: ctutils.BaseURL,
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create mexc client, %s", err)
	}

	return cli
}

func TestPing(t *testing.T) {
	cli := testNewContractMarketDataClient(t)

	_, err := cli.GetServerTime(context.TODO())
	assert.Nil(t, err)
}

func TestGetContractDetails(t *testing.T) {
	cli := testNewContractMarketDataClient(t)

	_, err := cli.GetContractDetails(context.TODO(), types.GetContractDetailsParams{})
	assert.Nil(t, err)
}

func TestGetTickerForSymbol(t *testing.T) {
	cli := testNewContractMarketDataClient(t)

	_, err := cli.GetTickerForSymbol(context.TODO(), types.GetTickerForSymbolParam{
		Symbol: "BTC_USDT",
	})
	assert.Nil(t, err)
}

func TestGetTickerForAllSymbols(t *testing.T) {
	cli := testNewContractMarketDataClient(t)

	_, err := cli.GetTickerForAllSymbols(context.TODO())
	assert.Nil(t, err)
}
