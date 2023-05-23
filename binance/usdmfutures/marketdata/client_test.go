package marketdata

import (
	"context"
	"testing"

	"github.com/linstohu/nexapi/binance/usdmfutures/marketdata/types"
	"github.com/linstohu/nexapi/binance/usdmfutures/utils"
	"github.com/stretchr/testify/assert"
)

func testUSDMFuturesMarketDataClient(t *testing.T) *USDMFuturesMarketDataClient {
	cli, err := NewUSDMFuturesMarketDataClient(&utils.USDMarginedClientCfg{
		BaseURL: utils.USDMarginedBaseURL,
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create binance client, %s", err)
	}

	return cli
}

func TestPing(t *testing.T) {
	cli := testUSDMFuturesMarketDataClient(t)

	err := cli.Ping(context.TODO())
	assert.Nil(t, err)
}

func TestGetServerTime(t *testing.T) {
	cli := testUSDMFuturesMarketDataClient(t)

	_, err := cli.GetServerTime(context.TODO())
	assert.Nil(t, err)
}

func TestGetExchangeInfo(t *testing.T) {
	cli := testUSDMFuturesMarketDataClient(t)

	_, err := cli.GetExchangeInfo(context.TODO())
	assert.Nil(t, err)
}

func TestGetOrderbook(t *testing.T) {
	cli := testUSDMFuturesMarketDataClient(t)

	_, err := cli.GetOrderbook(context.TODO(), types.GetOrderbookParams{
		Symbol: "BTCUSDT",
		Limit:  10,
	})
	assert.Nil(t, err)
}

func TestGetRecentTradeList(t *testing.T) {
	cli := testUSDMFuturesMarketDataClient(t)

	_, err := cli.GetRecentTradeList(context.TODO(), types.GetTradeParams{
		Symbol: "BTCUSDT",
		Limit:  10,
	})
	assert.Nil(t, err)
}

func TestGetAggTrades(t *testing.T) {
	cli := testUSDMFuturesMarketDataClient(t)

	_, err := cli.GetAggTrades(context.TODO(), types.GetAggTradesParam{
		Symbol: "BTCUSDT",
		Limit:  5,
	})
	assert.Nil(t, err)
}

func TestGetKlines(t *testing.T) {
	cli := testUSDMFuturesMarketDataClient(t)

	_, err := cli.GetKlines(context.TODO(), types.GetKlineParam{
		Symbol:   "BTCUSDT",
		Interval: utils.Minute1,
		Limit:    2,
	})
	assert.Nil(t, err)
}

func TestGetMarkPrice(t *testing.T) {
	cli := testUSDMFuturesMarketDataClient(t)

	_, err := cli.GetMarkPriceForSymbol(context.TODO(), types.GetMarkPriceParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)

	_, err = cli.GetMarkPriceForAllSymbols(context.TODO())
	assert.Nil(t, err)
}

func TestGetFundingRateHistory(t *testing.T) {
	cli := testUSDMFuturesMarketDataClient(t)

	_, err := cli.GetFundingRateHistory(context.TODO(), types.GetFundingRateParam{
		Symbol: "BTCUSDT",
		Limit:  2,
	})
	assert.Nil(t, err)
}

func TestGetTickerPrice(t *testing.T) {
	cli := testUSDMFuturesMarketDataClient(t)

	_, err := cli.GetTickerPriceForSymbol(context.TODO(), types.GetTickerPriceParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)

	_, err = cli.GetTickerPriceForAllSymbols(context.TODO())
	assert.Nil(t, err)
}

func TestGetBookTicker(t *testing.T) {
	cli := testUSDMFuturesMarketDataClient(t)

	_, err := cli.GetBookTickerForSymbol(context.TODO(), types.GetBookTickerForSymbolParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)

	_, err = cli.GetBookTickerForAllSymbols(context.TODO())
	assert.Nil(t, err)
}

func TestGetOpenInterestHistory(t *testing.T) {
	cli := testUSDMFuturesMarketDataClient(t)

	_, err := cli.GetOpenInterestHistory(context.TODO(), types.GetOpenInterestHistParam{
		Symbol: "BTCUSDT",
		Period: "5m",
		Limit:  2,
	})
	assert.Nil(t, err)
}
