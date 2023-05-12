package api

import (
	"context"
	"fmt"
	"testing"

	"github.com/linstohu/nexapi/woox/rest/types"
	"github.com/stretchr/testify/assert"
)

func testNewWooXRestPublicClient(t *testing.T) *WooXRestClient {
	cli, err := NewWooXRestClient(&WooXRestClientCfg{
		BasePath: "https://api.woo.org",
		Debug:    false,
	})

	if err != nil {
		t.Fatalf("Could not create woox client, %s", err)
	}

	return cli
}

func TestGetPublicInfo(t *testing.T) {
	woox := testNewWooXRestPublicClient(t)

	_, err := woox.GetPublicInfo(context.TODO())
	assert.Nil(t, err)
}

func TestGetPublicInfoForSymbol(t *testing.T) {
	woox := testNewWooXRestPublicClient(t)

	resp, err := woox.GetPublicInfoForSymbol(context.TODO(), "SPOT_ETH_USDT")
	assert.Nil(t, err)

	fmt.Printf("%+v\n", resp)
}

func TestGetPublicMarketTrades(t *testing.T) {
	woox := testNewWooXRestPublicClient(t)

	_, err := woox.GetPublicMarketTrades(context.TODO(), &types.GetMarketTradesParam{})
	assert.NotNil(t, err)

	_, err = woox.GetPublicMarketTrades(context.TODO(), &types.GetMarketTradesParam{
		Symbol: "SPOT_ETH_USDT",
	})
	assert.Nil(t, err)
}

func TestGetPublicOrderbook(t *testing.T) {
	woox := testNewWooXRestPublicClient(t)

	_, err := woox.GetPublicOrderbook(context.TODO(), "SPOT_ETH_USDT", &types.GetOrderbookParam{})
	assert.Nil(t, err)
}

func TestGetPublicKline(t *testing.T) {
	woox := testNewWooXRestPublicClient(t)

	_, err := woox.GetPublicKline(context.TODO(), &types.GetKlineParam{
		Symbol: "SPOT_ETH_USDT",
		Type:   "5m",
		Limit:  100,
	})
	assert.Nil(t, err)
}

func TestGetPublicTokens(t *testing.T) {
	woox := testNewWooXRestPublicClient(t)

	_, err := woox.GetPublicTokens(context.TODO())
	assert.Nil(t, err)
}

func TestGetPublicFundingRates(t *testing.T) {
	woox := testNewWooXRestPublicClient(t)

	_, err := woox.GetPublicFundingRates(context.TODO())
	assert.Nil(t, err)
}

func TestGetPublicFundingRateForSymbol(t *testing.T) {
	woox := testNewWooXRestPublicClient(t)

	_, err := woox.GetPublicFundingRateForSymbol(context.TODO(), "PERP_ETC_USDT")
	assert.Nil(t, err)
}

func TestGetPublicFuturesInfo(t *testing.T) {
	woox := testNewWooXRestPublicClient(t)

	_, err := woox.GetPublicAllFuturesInfo(context.TODO())
	assert.Nil(t, err)
}

func TestGetPublicFuturesInfoForSymbol(t *testing.T) {
	woox := testNewWooXRestPublicClient(t)

	_, err := woox.GetPublicFuturesInfoForSymbol(context.TODO(), "PERP_ETC_USDT")
	assert.Nil(t, err)
}
