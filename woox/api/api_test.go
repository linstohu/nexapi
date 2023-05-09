package api

import (
	"testing"

	"github.com/linstohu/nexapi/woox/api/types"
	"github.com/stretchr/testify/assert"
)

func TestGetPublicMarketTrades(t *testing.T) {
	woox, err := NewWooXClient(&WooXCfg{
		BasePath: "https://api.woo.org",
		Debug:    false,
	})
	assert.Nil(t, err)

	_, err = woox.GetPublicMarketTrades(&types.GetMarketTradesParam{})
	assert.NotNil(t, err)

	_, err = woox.GetPublicMarketTrades(&types.GetMarketTradesParam{
		Symbol: "SPOT_ETH_USDT",
	})
	assert.Nil(t, err)
}
