package spotaccount

import (
	"context"
	"os"
	"testing"

	spotutils "github.com/linstohu/nexapi/mexc/spot/utils"
	"github.com/stretchr/testify/assert"
)

func testNewAccountClient(t *testing.T) *SpotAccountClient {
	cli, err := NewSpotAccountClient(&SpotAccountClientCfg{
		BaseURL: spotutils.BaseURL,
		Key:     os.Getenv("MEXC_KEY"),
		Secret:  os.Getenv("MEXC_SECRET"),
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create mexc client, %s", err)
	}

	return cli
}

func TestGetAccountInfo(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetAccountInfo(context.TODO())
	assert.Nil(t, err)
}
