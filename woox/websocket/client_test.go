package websocket

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/linstohu/nexapi/woox/websocket/types"
	"github.com/stretchr/testify/assert"
)

func testNewWooXStreamClient(ctx context.Context, t *testing.T) *WooXStreamClient {
	cli, err := NewWooXClient(ctx, &WooXStreamCfg{
		BasePath:      PublicTestnetBaseEndpoint,
		ApplicationID: os.Getenv("WOOX_APP_ID"), // required
		Debug:         false,
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

func TestSubscribeOrderbook(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewWooXStreamClient(ctx, t)

	topic, err := cli.GetOrderbookTopic("PERP_BTC_USDT")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e interface{}) {
		book, ok := e.(*types.Orderbook)
		if !ok {
			return
		}

		if len(book.Data.Asks) == 0 || len(book.Data.Bids) == 0 {
			return
		}

		fmt.Printf("获取到新的Orderbook, 币对: %v, 时间戳: %v,  买单数量: %v, 卖单数量: %v\n",
			book.Data.Symbol, book.Ts, len(book.Data.Bids), len(book.Data.Asks))
	})

	cli.Subscribe([]string{topic})

	select {}
}
