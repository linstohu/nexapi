package api

import (
	"net/http"
	"testing"

	"github.com/linstohu/nexapi/woox/api/types"
	"github.com/stretchr/testify/assert"
)

func testNewWooXPrivateClient(t *testing.T) *WooXClient {
	cli, err := NewWooXClient(&WooXCfg{
		BasePath: "https://api.woo.org",
		Key:      "", // required
		Secret:   "", // required
		Debug:    false,
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

	content, err := NormalizeV1RequestContent(req)
	assert.Nil(t, err)
	assert.Equal(t, "order_price=9000&order_quantity=0.11&order_type=LIMIT&side=BUY&symbol=SPOT_BTC_USDT", content)
}
