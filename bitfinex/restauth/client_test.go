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

package restauth

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testRestAuthClient(t *testing.T) *RestAuthClient {
	cli, err := NewRestAuthClient(&BitfinexClientCfg{
		BaseURL: BaseURL,
		Debug:   true,
		Key:     os.Getenv("BITFINEX_KEY"),
		Secret:  os.Getenv("BBITFINEX_SECRET"),
	})

	if err != nil {
		t.Fatalf("Could not create bitfinex authenticated client, %s", err)
	}

	return cli
}

func TestGetWallets(t *testing.T) {
	cli := testRestAuthClient(t)

	err := cli.GetWallets(context.TODO())
	assert.Nil(t, err)
}
