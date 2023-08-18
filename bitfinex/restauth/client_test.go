package bitfinexrestauth

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
