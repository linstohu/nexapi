package bitfinexrestpub

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testNewRestPubClient(t *testing.T) *RestPubClient {
	cli, err := NewRestPubClient(&BitfinexClientCfg{
		BaseURL: BaseURL,
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create bitfinex client, %s", err)
	}

	return cli
}

func TestPlatformStatus(t *testing.T) {
	cli := testNewRestPubClient(t)

	err := cli.PlatformStatus(context.TODO())
	assert.Nil(t, err)
}
