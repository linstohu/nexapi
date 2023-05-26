package account

import (
	"context"
	"os"
	"testing"

	"github.com/linstohu/nexapi/binance/coinmfutures/account/types"
	cmutils "github.com/linstohu/nexapi/binance/coinmfutures/utils"
	umutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	"github.com/stretchr/testify/assert"
)

func testNewAccountClient(t *testing.T) *CoinMFuturesAccountClient {
	cli, err := NewCoinMFuturesAccountClient(&cmutils.CoinMarginedClientCfg{
		BaseURL: umutils.USDMarginedBaseURL,
		Key:     os.Getenv("BINANCE_KEY"),
		Secret:  os.Getenv("BINANCE_SECRET"),
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create binance client, %s", err)
	}

	return cli
}

func TestChangePositionMode(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.ChangePositionMode(context.TODO(), types.ChangePositionModeParam{
		DualSidePosition: "true",
	})
	assert.Nil(t, err)
}

func TestGetPositionMode(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetPositionMode(context.TODO())
	assert.Nil(t, err)
}
