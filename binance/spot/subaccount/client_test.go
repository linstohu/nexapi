package subaccount

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/linstohu/nexapi/binance/spot/subaccount/types"
	spotutils "github.com/linstohu/nexapi/binance/spot/utils"
	"github.com/stretchr/testify/assert"
)

func testNewSpotSubAccountClient(t *testing.T) *SpotSubAccountClient {
	cli, err := NewSpotSubAccountClient(&SpotSubAccountClientCfg{
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

func TestGetSubAccountTransferHistory(t *testing.T) {
	cli := testNewSpotSubAccountClient(t)

	now := time.Now()
	start := now.Add(-24 * time.Hour * 60)
	_, err := cli.GetSubAccountTransferHistory(context.TODO(), types.GetSubAccountTransferHistoryParam{
		Type:      1,
		StartTime: start.UnixMilli(),
		EndTime:   now.UnixMilli(),
	})
	assert.Nil(t, err)
}
