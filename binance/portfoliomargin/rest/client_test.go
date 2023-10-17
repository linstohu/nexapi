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

package pmrest

import (
	"context"
	"os"
	"testing"

	pmutils "github.com/linstohu/nexapi/binance/portfoliomargin/utils"
	"github.com/stretchr/testify/assert"
)

func testNewAccountClient(t *testing.T) *PortfolioMarginAccountClient {
	cli, err := NewPortfolioMarginAccountClient(&pmutils.PortfolioMarginClientCfg{
		BaseURL: pmutils.PortfolioMarginBaseURL,
		Key:     os.Getenv("BINANCE_KEY"),
		Secret:  os.Getenv("BINANCE_SECRET"),
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create binance client, %s", err)
	}

	return cli
}

func TestGetBalance(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetBalance(context.TODO())
	assert.Nil(t, err)
}
