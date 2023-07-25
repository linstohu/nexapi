package margin

import (
	"context"
	"os"
	"testing"

	"github.com/linstohu/nexapi/binance/spot/margin/types"
	spotutils "github.com/linstohu/nexapi/binance/spot/utils"
	"github.com/stretchr/testify/assert"
)

func testNewSpotMarginClient(t *testing.T) *SpotMarginClient {
	cli, err := NewSpotMarginClient(&SpotMarginClientCfg{
		BaseURL: spotutils.BaseURL,
		Key:     os.Getenv("BINANCE_KEY"),
		Secret:  os.Getenv("BINANCE_SECRET"),
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create binance client, %s", err)
	}

	return cli
}

func TestGetInterestHistory(t *testing.T) {
	cli := testNewSpotMarginClient(t)

	// now := time.Now()
	// start := now.Add(-24 * time.Hour * 60)
	_, err := cli.GetInterestHistory(context.TODO(), types.GetInterestHistoryParam{
		IsolatedSymbol: "BTCUSDT",
	})
	assert.Nil(t, err)
}
