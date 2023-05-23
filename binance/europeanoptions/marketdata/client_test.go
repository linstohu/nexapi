package marketdata

import (
	"context"
	"testing"

	"github.com/linstohu/nexapi/binance/europeanoptions/marketdata/types"
	"github.com/linstohu/nexapi/binance/europeanoptions/utils"
	usdmtypes "github.com/linstohu/nexapi/binance/usdmfutures/marketdata/types"
	usdmutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	"github.com/stretchr/testify/assert"
)

func testOptionsMarketDataClient(t *testing.T) *OptionsMarketDataClient {
	cli, err := NewOptionsMarketDataClient(&utils.OptionsClientCfg{
		BaseURL: utils.OptionsBaseURL,
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create binance client, %s", err)
	}

	return cli
}

func TestPing(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	err := cli.Ping(context.TODO())
	assert.Nil(t, err)
}

func TestGetServerTime(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetServerTime(context.TODO())
	assert.Nil(t, err)
}

func TestGetExchangeInfo(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetExchangeInfo(context.TODO())
	assert.Nil(t, err)
}

func TestGetOrderbook(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetOrderbook(context.TODO(), types.GetOrderbookParams{
		Symbol: "BTC-230630-25000-P",
		Limit:  10,
	})
	assert.Nil(t, err)
}

func TestGetRecentTradesList(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetRecentTradesList(context.TODO(), types.GetTradeParams{
		Symbol: "BTC-230630-25000-P",
		Limit:  10,
	})
	assert.Nil(t, err)
}

func TestGetKlines(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetKlines(context.TODO(), usdmtypes.GetKlineParam{
		Symbol:   "BTC-230630-25000-P",
		Interval: usdmutils.Minute1,
		Limit:    2,
	})
	assert.Nil(t, err)
}

func TestGetMarkPrice(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetMarkPrice(context.TODO(), types.GetMarkPriceParam{
		Symbol: "BTC-230630-25000-P",
	})
	assert.Nil(t, err)
}

func TestGetTickerPrice(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetTickerPrice(context.TODO(), types.GetTickerPriceParam{
		Symbol: "BTC-230630-25000-P",
	})
	assert.Nil(t, err)
}

func TestGetUnderlyingIndexPrice(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetUnderlyingIndexPrice(context.TODO(), types.GetUnderlyingIndexPriceParams{
		Underlying: "BTCUSDT",
	})
	assert.Nil(t, err)
}

func TestGetOpenInterest(t *testing.T) {
	cli := testOptionsMarketDataClient(t)

	_, err := cli.GetOpenInterest(context.TODO(), types.GetOpenInterestParam{
		UnderlyingAsset: "BTC",
		Expiration:      "230630",
	})
	assert.Nil(t, err)
}
