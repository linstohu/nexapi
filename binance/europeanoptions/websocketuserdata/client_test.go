/*
 * Copyright (c) 2023, LinstoHu
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package websocketuserdata

import (
	"fmt"
	"os"
	"testing"
)

func testNewUserDataStreamClient(t *testing.T) *OptionsUserDataStreamClient {
	cli, err := NewUserDataStreamClient(&OptionsUserDataStreamCfg{
		Debug:         true,
		BaseURL:       OptionsUserDataStreamBaseURL,
		AutoReconnect: true,
		Key:           os.Getenv("BINANCE_KEY"),
		Secret:        os.Getenv("BINANCE_SECRET"),
	})

	if err != nil {
		t.Fatalf("Could not create websocket client for woox, %s", err)
	}

	return cli
}

func TestSubscribeAccountData(t *testing.T) {
	cli := testNewUserDataStreamClient(t)

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
	cli := testNewUserDataStreamClient(t)

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
