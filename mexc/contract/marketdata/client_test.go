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
