package websocketuserdata

import (
	"context"
	"fmt"
	"testing"
)

func testNewUserDataStreamClient(ctx context.Context, t *testing.T) *OptionsUserDataStreamClient {
	cli, err := NewUserDataStreamClient(ctx, &OptionsUserDataStreamCfg{
		BaseURL: OptionsUserDataStreamBaseURL,
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create websocket client for woox, %s", err)
	}

	return cli
}

func TestSubscribeAccountData(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewUserDataStreamClient(ctx, t)

	topic := cli.GenAccountDataTopic()

	cli.AddListener(topic, func(e any) {
		account, ok := e.(*AccountData)
		if !ok {
			return
		}

		for _, v := range account.Balances {
			fmt.Printf("Balance: Asset: %s, Balance: %v, UnPNL: %v\n", v.MarginAsset, v.AccountBalance, v.UnPNL)
		}
		for _, v := range account.Greek {
			fmt.Printf("Greek: Underlying: %s, Delta: %v, Theta: %v\n", v.Underlying, v.Delta, v.Theta)
		}
		for _, v := range account.Position {
			fmt.Printf("Position: Symbol: %s, PositionNum: %v, EntryPrice: %v\n", v.Symbol, v.PositionNum, v.EntryPrice)
		}
	})

	select {}
}

func TestSubscribeOrderUpdate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewUserDataStreamClient(ctx, t)

	topic := cli.GenOrderUpdateTopic()

	cli.AddListener(topic, func(e any) {
		orders, ok := e.(*OrderUpdate)
		if !ok {
			return
		}

		for _, order := range orders.Orders {
			fmt.Printf("Topic: %s, Symbol: %v, OrderType: %v, Price: %v, Quantity: %v, Time: %v\n",
				topic, order.Symbol, order.OrderType, order.OrderPrice, order.OrderQuantity, order.CreateTime)
		}
	})

	select {}
}
