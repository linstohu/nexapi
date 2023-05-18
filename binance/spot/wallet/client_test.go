package wallet

import (
	"context"
	"os"
	"testing"

	"github.com/linstohu/nexapi/binance/spot/utils"
	"github.com/stretchr/testify/assert"
)

func testNewSpotWalletClient(t *testing.T) *SpotWalletClient {
	cli, err := NewSpotWalletClient(&utils.SpotClientCfg{
		BaseURL: utils.BaseURL,
		Key:     os.Getenv("BINANCE_KEY"),
		Secret:  os.Getenv("BINANCE_SECRET"),
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create binance client, %s", err)
	}

	return cli
}

func TestGetSystemStatus(t *testing.T) {
	cli := testNewSpotWalletClient(t)

	_, err := cli.GetSystemStatus(context.TODO())
	assert.Nil(t, err)
}
