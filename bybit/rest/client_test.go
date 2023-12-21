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

package rest

import (
	"fmt"
	"os"
	"testing"

	"github.com/linstohu/nexapi/bybit/utils"
	"github.com/stretchr/testify/assert"
)

func testNewClient(t *testing.T) *BybitClient {
	cli, err := NewBybitClient(&utils.BybitClientCfg{
		Debug:   true,
		BaseURL: utils.TestBaseURL,
		Key:     os.Getenv("BYBIT_KEY"),
		Secret:  os.Getenv("BYBIT_SECRET"),
	})

	if err != nil {
		t.Fatalf("Could not create bybit client, %s", err)
	}

	return cli
}

func TestGetUnifiedAccountBalance(t *testing.T) {
	cli := testNewClient(t)

	resp, err := cli.GetUnifiedAccountBalance()
	assert.Nil(t, err)

	for _, v := range resp.Body.Result.List {
		fmt.Printf("%+v\n", v)
	}
}

func TestGetUnifiedAccountContractBalance(t *testing.T) {
	cli := testNewClient(t)

	resp, err := cli.GetUnifiedAccountContractBalance()
	assert.Nil(t, err)

	for _, v := range resp.Body.Result.List {
		fmt.Printf("%+v\n", v)
	}
}

func TestGetFundAccountBalance(t *testing.T) {
	cli := testNewClient(t)

	resp, err := cli.GetFundAccountBalance()
	assert.Nil(t, err)

	for _, v := range resp.Body.Result.Balance {
		fmt.Printf("%+v\n", v)
	}
}
