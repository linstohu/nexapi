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
	"context"
	"os"
	"testing"

	"github.com/linstohu/nexapi/htx/utils"
	"github.com/stretchr/testify/assert"
)

func testNewSpotClient(t *testing.T) *SpotClient {
	cli, err := NewSpotClient(&SpotClientCfg{
		BaseURL:     utils.ProdAWSBaseURL,
		Key:         os.Getenv("HTX_KEY"),
		Secret:      os.Getenv("HTX_SECRET"),
		SignVersion: utils.ApiKeyVersionV2,
		Debug:       true,
	})

	if err != nil {
		t.Fatalf("Could not create htx client, %s", err)
	}

	return cli
}

func TestGetAccountInfo(t *testing.T) {
	cli := testNewSpotClient(t)

	_, err := cli.GetAccountInfo(context.TODO())

	assert.Nil(t, err)
}
