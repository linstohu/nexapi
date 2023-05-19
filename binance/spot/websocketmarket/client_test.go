package websocketmarket

import (
	"context"
	"fmt"
	"testing"

	"github.com/linstohu/nexapi/binance/spot/websocketmarket/types"
	"github.com/stretchr/testify/assert"
)

func testNewMarketStreamClient(ctx context.Context, t *testing.T) *MarketStreamClient {
	cli, err := NewMarketStreamClient(ctx, &MarketStreamCfg{
		BaseURL: MarketStreamBaseURL,
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create websocket client for woox, %s", err)
	}

	return cli
}

func TestSubscribeAggTrade(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetAggTradeTopic("btcusdt")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		trade, ok := e.(*types.AggregateTrade)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, Price: %v, Quantity: %v, Time: %v\n",
			topic, trade.Symbol, trade.Price, trade.Quantity, trade.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}
