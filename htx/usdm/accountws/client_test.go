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
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/linstohu/nexapi/htx/usdm/accountws/types"
	"github.com/stretchr/testify/assert"
)

func testNewAccountWsClient(ctx context.Context, t *testing.T, url string) *AccountWsClient {
	cli, err := NewAccountWsClient(ctx, &AccountWsClientCfg{
		BaseURL: url,
		Debug:   true,
		Key:     os.Getenv("HTX_KEY"),
		Secret:  os.Getenv("HTX_SECRET"),
	})

	if err != nil {
		t.Fatalf("Could not create websocket client, %s", err)
	}

	return cli
}

func TestSubscribeAccountUpdate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewAccountWsClient(ctx, t, GlobalOrderWsBaseURL)

	topic, err := cli.GetCrossAccountUpdateTopic("ETH-USDT")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		acc, ok := e.(*types.CrossAccount)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Data: %+v\n",
			topic, acc.Data)
	})

	cli.Subscribe(topic)

	select {}
}
