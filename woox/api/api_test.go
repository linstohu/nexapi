package api

import (
	"testing"

	"github.com/linstohu/nexapi/woox/api/types"
	"github.com/stretchr/testify/assert"
)

func testNewWooXClient(t *testing.T) *WooXClient {
	cli, err := NewWooXClient(&WooXCfg{
		BasePath: "https://api.woo.org",
		Debug:    false,
	})

	if err != nil {
		t.Fatalf("Could not create woox client, %s", err)
	}

	return cli
}

func TestGetPublicMarketTrades(t *testing.T) {
	woox := testNewWooXClient(t)

	_, err := woox.GetPublicMarketTrades(&types.GetMarketTradesParam{})
	assert.NotNil(t, err)

	_, err = woox.GetPublicMarketTrades(&types.GetMarketTradesParam{
		Symbol: "SPOT_ETH_USDT",
	})
	assert.Nil(t, err)
}
