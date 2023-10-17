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

package publicdata

import (
	"context"
	"os"
	"testing"

	"github.com/linstohu/nexapi/okx/tradingaccount/types"
	okxutils "github.com/linstohu/nexapi/okx/utils"
	"github.com/stretchr/testify/assert"
)

func testNewTradingAccountClient(t *testing.T) *TradingAccountClient {
	cli, err := NewTradingAccountClient(&TradingAccountClientCfg{
		Debug:      true,
		BaseURL:    okxutils.RestURL,
		Key:        os.Getenv("OKX_KEY"),
		Secret:     os.Getenv("OKX_SECRET"),
		Passphrase: os.Getenv("OKX_PASS"),
	})

	if err != nil {
		t.Fatalf("Could not create okx private client, %s", err)
	}

	return cli
}

func TestGetBalance(t *testing.T) {
	cli := testNewTradingAccountClient(t)

	_, err := cli.GetBalance(context.TODO(), types.GetBalanceParam{})
	assert.Nil(t, err)
}
