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

package accountws

import (
	"fmt"
	"os"
	"testing"

	"github.com/linstohu/nexapi/htx/spot/accountws/types"
	"github.com/stretchr/testify/assert"
)

func testNewAccountWsClient(t *testing.T, url string) *AccountWsClient {
	cli, err := NewAccountWsClient(&AccountWsClientCfg{
		Debug:         true,
		BaseURL:       url,
		AutoReconnect: true,
		Key:           os.Getenv("HTX_KEY"),
		Secret:        os.Getenv("HTX_SECRET"),
	})

	if err != nil {
		t.Fatalf("Could not create websocket client, %s", err)
	}

	return cli
}

func TestSubscribeAccountUpdate(t *testing.T) {
	cli := testNewAccountWsClient(t, GlobalWsBaseURL)

	topic, err := cli.GetAccountUpdateTopic(&AccountUpdateTopicParam{
		Mode: 2,
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		acc, ok := e.(*types.Account)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, AccountId: %v, Currency: %v, Balance: %v, Available: %v, ChangeType: %v, AccountType: %v, ChangeTime: %v, SeqNum: %v\n",
			topic, acc.AccountID, acc.Currency, acc.Balance, acc.Available, acc.ChangeType, acc.AccountType, acc.ChangeTime, acc.SeqNum)
	})

	cli.Subscribe(topic)

	select {}
}
