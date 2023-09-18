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
