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
	"testing"

	"github.com/linstohu/nexapi/deribit/rest/types/marketdata"
	"github.com/stretchr/testify/assert"
)

func testNewDeribitRestPublicClient(t *testing.T) *DeribitRestClient {
	cli, err := NewDeribitRestClient(&DeribitRestClientCfg{
		BaseURL: BaseURL,
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create deribit client, %s", err)
	}

	return cli
}

func TestMe(t *testing.T) {
	deribit := testNewDeribitRestPublicClient(t)

	_, err := deribit.Test(context.TODO())
	assert.Nil(t, err)
}

func TestGetBookSummaryByCurrency(t *testing.T) {
	deribit := testNewDeribitRestPublicClient(t)

	_, err := deribit.GetBookSummaryByCurrency(context.TODO(), marketdata.GetBookSummaryByCurrencyParams{
		Currency: "BTC",
	})
	assert.Nil(t, err)
}

func TestGetBookSummaryByInstrument(t *testing.T) {
	deribit := testNewDeribitRestPublicClient(t)

	_, err := deribit.GetBookSummaryByInstrument(context.TODO(), marketdata.GetBookSummaryByInstrumentParams{
		InstrumentName: "BTC-PERPETUAL",
	})
	assert.Nil(t, err)
}

func TestGetContractSize(t *testing.T) {
	deribit := testNewDeribitRestPublicClient(t)

	_, err := deribit.GetContractSize(context.TODO(), marketdata.GetContractSizeParams{
		InstrumentName: "BTC-PERPETUAL",
	})
	assert.Nil(t, err)
}

func TestGetCurrencies(t *testing.T) {
	deribit := testNewDeribitRestPublicClient(t)

	_, err := deribit.GetCurrencies(context.TODO())
	assert.Nil(t, err)
}

func TestGetFundingRate(t *testing.T) {
	deribit := testNewDeribitRestPublicClient(t)

	_, err := deribit.GetFundingRate(context.TODO(), marketdata.GetFundingRateParams{
		InstrumentName: "ETH-PERPETUAL",
		StartTimestamp: 1569888000000,
		EndTimestamp:   1569974400000,
	})
	assert.Nil(t, err)
}

func TestGetIndexPrice(t *testing.T) {
	deribit := testNewDeribitRestPublicClient(t)

	_, err := deribit.GetIndexPrice(context.TODO(), marketdata.GetIndexPriceParams{
		IndexName: "btc_usd",
	})
	assert.Nil(t, err)
}

func TestGetOneInstrument(t *testing.T) {
	deribit := testNewDeribitRestPublicClient(t)

	_, err := deribit.GetInstrument(context.TODO(), marketdata.GetInstrumentParams{
		InstrumentName: "ETH-PERPETUAL",
	})
	assert.Nil(t, err)
}

func TestGetInstruments(t *testing.T) {
	deribit := testNewDeribitRestPublicClient(t)

	_, err := deribit.GetInstruments(context.TODO(), marketdata.GetInstrumentsParams{
		Currency: "BTC",
		Kind:     "future",
	})
	assert.Nil(t, err)
}

func TestGetLastTradesByInstrumentAndTime(t *testing.T) {
	deribit := testNewDeribitRestPublicClient(t)

	_, err := deribit.GetLastTradesByInstrumentAndTime(context.TODO(), marketdata.GetLastTradesByInstrumentAndTimeParams{
		InstrumentName: "ETH-PERPETUAL",
		EndTimestamp:   1696953839705,
	})
	assert.Nil(t, err)
}

func TestGetOrderBook(t *testing.T) {
	deribit := testNewDeribitRestPublicClient(t)

	_, err := deribit.GetOrderBook(context.TODO(), marketdata.GetOrderBookParams{
		InstrumentName: "ETH-PERPETUAL",
		Depth:          5,
	})
	assert.Nil(t, err)
}

func TestGetTradingviewChartData(t *testing.T) {
	deribit := testNewDeribitRestPublicClient(t)

	_, err := deribit.GetTradingviewChartData(context.TODO(), marketdata.GetTradingviewChartDataParams{
		InstrumentName: "ETH-PERPETUAL",
		StartTimestamp: 1696947123000,
		EndTimestamp:   1696950723000,
		Resolution:     "30",
	})
	assert.Nil(t, err)
}

func TestGetTicker(t *testing.T) {
	deribit := testNewDeribitRestPublicClient(t)

	_, err := deribit.GetTicker(context.TODO(), marketdata.TickerParams{
		InstrumentName: "ETH-PERPETUAL",
	})
	assert.Nil(t, err)
}
