package marketdata

import (
	"context"
	"testing"

	"github.com/linstohu/nexapi/binance/spot/marketdata/types"
	"github.com/linstohu/nexapi/binance/spot/utils"
	"github.com/stretchr/testify/assert"
)

func testNewSpotMarketDataClient(t *testing.T) *SpotMarketDataClient {
	cli, err := NewSpotMarketDataClient(&utils.SpotClientCfg{
		BaseURL: utils.BaseURL,
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create binance client, %s", err)
	}

	return cli
}

func TestPing(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	err := cli.Ping(context.TODO())
	assert.Nil(t, err)
}

func TestGetServerTime(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetServerTime(context.TODO())
	assert.Nil(t, err)
}

func TestGetExchangeInfo(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetExchangeInfo(context.TODO(), types.GetExchangeInfoParam{
		Symbols: []string{},
	})
	assert.Nil(t, err)

	_, err = cli.GetExchangeInfo(context.TODO(), types.GetExchangeInfoParam{
		Symbols: []string{"BTCUSDT"},
	})
	assert.Nil(t, err)

	_, err = cli.GetExchangeInfo(context.TODO(), types.GetExchangeInfoParam{
		Symbols: []string{"BTCUSDT", "DOGEUSDT"},
	})
	assert.Nil(t, err)

	_, err = cli.GetExchangeInfo(context.TODO(), types.GetExchangeInfoParam{
		Permissions: []string{},
	})
	assert.Nil(t, err)

	_, err = cli.GetExchangeInfo(context.TODO(), types.GetExchangeInfoParam{
		Permissions: []string{"SPOT"},
	})
	assert.Nil(t, err)

	_, err = cli.GetExchangeInfo(context.TODO(), types.GetExchangeInfoParam{
		Permissions: []string{"SPOT", "MARGIN", "LEVERAGED"},
	})
	assert.Nil(t, err)
}

func TestGetOrderbook(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetOrderbook(context.TODO(), types.GetOrderbookParams{
		Symbol: "BTCUSDT",
		Limit:  10,
	})
	assert.Nil(t, err)
}

func TestGetRecentTradeList(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetRecentTradeList(context.TODO(), types.GetTradeParams{
		Symbol: "BTCUSDT",
		Limit:  10,
	})
	assert.Nil(t, err)
}

func TestGetAggTrades(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetAggTrades(context.TODO(), types.GetAggTradesParam{
		Symbol: "BTCUSDT",
		Limit:  5,
	})
	assert.Nil(t, err)
}

func TestGetKlines(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetKlines(context.TODO(), types.GetKlineParam{
		Symbol:   "BTCUSDT",
		Interval: utils.Minute1,
		Limit:    1,
	})
	assert.Nil(t, err)
}

func TestGetAvgPrice(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetAvgPrice(context.TODO(), types.GetAvgPriceParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)
}

func TestGetTickerForSymbol(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetTickerForSymbol(context.TODO(), types.GetTickerForSymbolParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)
}

func TestGetTickerForSymbols(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetTickerForSymbols(context.TODO(), types.GetTickerForSymbolsParam{
		Symbols: []string{"BTCUSDT", "DOGEUSDT"},
	})
	assert.Nil(t, err)
}

func TestGetTickerPriceForSymbol(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetTickerPriceForSymbol(context.TODO(), types.GetTickerPriceForSymbolParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)
}

func TestGetTickerPriceForSymbols(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetTickerPriceForSymbols(context.TODO(), types.GetTickerPriceForSymbolsParam{
		Symbols: []string{"BTCUSDT", "DOGEUSDT"},
	})
	assert.Nil(t, err)
}

func TestGetBookTickerForSymbol(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetBookTickerForSymbol(context.TODO(), types.GetBookTickerForSymbolParam{
		Symbol: "BTCUSDT",
	})
	assert.Nil(t, err)
}

func TestGetBookTickerForSymbols(t *testing.T) {
	cli := testNewSpotMarketDataClient(t)

	_, err := cli.GetBookTickerForSymbols(context.TODO(), types.GetBookTickerForSymbolsParam{
		Symbols: []string{"BTCUSDT", "DOGEUSDT"},
	})
	assert.Nil(t, err)
}
