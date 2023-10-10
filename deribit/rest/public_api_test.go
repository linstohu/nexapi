package rest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testNewDeribitRestPublicClient(t *testing.T) *DeribitRestClient {
	cli, err := NewDeribitRestClient(&DeribitRestClientCfg{
		BaseURL: BaseURL,
		Debug:   false,
	})

	if err != nil {
		t.Fatalf("Could not create deribit client, %s", err)
	}

	return cli
}

func TestMe(t *testing.T) {
	deribit := testNewDeribitRestPublicClient(t)

	_, err := deribit.Test(context.TODO())
	assert.Nil(t, err)
}
