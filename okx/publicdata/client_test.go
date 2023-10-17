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
	"testing"

	"github.com/linstohu/nexapi/okx/publicdata/types"
	okxutils "github.com/linstohu/nexapi/okx/utils"
	"github.com/stretchr/testify/assert"
)

func testNewPublicDataClient(t *testing.T) *PublicDataClient {
	cli, err := NewPublicDataClient(&okxutils.OKXRestClientCfg{
		BaseURL: okxutils.RestURL,
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create okx client, %s", err)
	}

	return cli
}

func TestGetInstruments(t *testing.T) {
	cli := testNewPublicDataClient(t)

	_, err := cli.GetInstruments(context.TODO(), types.GetInstrumentsParam{
		InstType: okxutils.Spot,
	})
	assert.Nil(t, err)
}
