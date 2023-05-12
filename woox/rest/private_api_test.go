package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/linstohu/nexapi/woox/rest/types"
	"github.com/stretchr/testify/assert"
)

func testNewWooXRestPrivateClient(t *testing.T) *WooXRestClient {
	cli, err := NewWooXRestClient(&WooXRestClientCfg{
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
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.SendOrder(context.TODO(), types.SendOrderReq{
		Symbol:        "SPOT_BTC_USDT",
		OrderType:     MarketOrderType,
		Side:          BUY,
		OrderQuantity: 0.1,
	})
	assert.Nil(t, err)
}

func TestGetOrders(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetOrders(context.TODO(), types.GetOrdersParam{})
	assert.Nil(t, err)
}

func TestGetTradeHistory(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetTradeHistory(context.TODO(), types.GetTradeHistoryParam{})
	assert.Nil(t, err)
}

func TestGetBalances(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetBalances(context.TODO())
	assert.Nil(t, err)
}

func TestGetAccountInfo(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetAccountInfo(context.TODO())
	assert.Nil(t, err)
}

func TestGetAssetHisotry(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetAssetHisotry(context.TODO(), types.GetAssetHisotryParam{})
	assert.Nil(t, err)
}

func TestGetSubAccounts(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetSubAccounts(context.TODO())
	assert.Nil(t, err)
}

func TestTransferAsset(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.TransferAsset(context.TODO(), types.TransferAssetParam{
		Token:     "USDT",
		Amount:    100,
		FromAppID: "",
		ToAppID:   "",
	})
	assert.Nil(t, err)
}

func TestGetIPRestriction(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetIPRestriction(context.TODO())
	assert.Nil(t, err)
}

func TestGetOnePositionInfo(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetOnePositionInfo(context.TODO(), "PERP_BTC_USDT")
	assert.Nil(t, err)
}

func TestGetAllPositionInfo(t *testing.T) {
	woox := testNewWooXRestPrivateClient(t)

	_, err := woox.GetAllPositionInfo(context.TODO())
	assert.Nil(t, err)
}
