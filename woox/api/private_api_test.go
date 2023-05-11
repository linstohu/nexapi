package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/linstohu/nexapi/woox/api/types"
	"github.com/stretchr/testify/assert"
)

func testNewWooXPrivateClient(t *testing.T) *WooXClient {
	cli, err := NewWooXClient(&WooXCfg{
		BasePath: "https://api.staging.woo.org",
		Key:      os.Getenv("WOOX_KEY"),    // required
		Secret:   os.Getenv("WOOX_SECRET"), // required
		Debug:    true,
	})

	if err != nil {
		t.Fatalf("Could not create woox client, %s", err)
	}

	return cli
}

func TestNormalizeV1RequestContent(t *testing.T) {
	req := types.HTTPRequest{
		Path:   "/v1/order",
		Method: http.MethodPost,
		Body: types.SendOrderReq{
			Symbol:        "SPOT_BTC_USDT",
			OrderType:     "LIMIT",
			OrderPrice:    9000,
			OrderQuantity: 0.11,
			Side:          "BUY",
		},
	}

	content, err := normalizeV1RequestContent(req)
	assert.Nil(t, err)
	assert.Equal(t, "order_price=9000&order_quantity=0.11&order_type=LIMIT&side=BUY&symbol=SPOT_BTC_USDT", content)
}

func TestSendOrder(t *testing.T) {
	woox := testNewWooXPrivateClient(t)

	_, err := woox.SendOrder(context.TODO(), types.SendOrderReq{
		Symbol:        "SPOT_BTC_USDT",
		OrderType:     MarketOrderType,
		Side:          BUY,
		OrderQuantity: 0.1,
	})
	assert.Nil(t, err)
}

func TestGetPrivateBalances(t *testing.T) {
	woox := testNewWooXPrivateClient(t)

	_, err := woox.GetPrivateBalances(context.TODO())
	assert.Nil(t, err)
}

func TestGetAssetHisotry(t *testing.T) {
	woox := testNewWooXPrivateClient(t)

	_, err := woox.GetAssetHisotry(context.TODO(), types.GetAssetHisotryParam{})
	assert.Nil(t, err)
}

func TestGetIPRestriction(t *testing.T) {
	woox := testNewWooXPrivateClient(t)

	_, err := woox.GetIPRestriction(context.TODO())
	assert.Nil(t, err)
}
