package websocket

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func testNewWooXStreamClient(ctx context.Context, t *testing.T) *WooXStreamClient {
	cli, err := NewWooXClient(ctx, &WooXStreamCfg{
		BasePath:      PublicProdBaseEndpoint,
		ApplicationID: "", // required
		Debug:         true,
	})

	if err != nil {
		t.Fatalf("Could not create websocket client for woox, %s", err)
	}

	return cli
}

func TestWooXStreamClientConnection(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewWooXStreamClient(ctx, t)

	err := cli.Subscribe([]string{"SPOT_WOO_USDT@orderbook"})
	assert.Nil(t, err)

	time.Sleep(1 * time.Minute)
}
