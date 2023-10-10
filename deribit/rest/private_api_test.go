package rest

import (
	"context"
	"os"
	"testing"

	"github.com/linstohu/nexapi/deribit/rest/types/account"
	"github.com/linstohu/nexapi/deribit/rest/types/auth"
	"github.com/stretchr/testify/assert"
)

func testNewDeribitRestPrivateClient(t *testing.T) *DeribitRestClient {
	cli, err := NewDeribitRestClient(&DeribitRestClientCfg{
		BaseURL: BaseURL,
		Debug:   true,
		Key:     os.Getenv("DERIBIT_KEY"),
		Secret:  os.Getenv("DERIBIT_SECRET"),
	})

	if err != nil {
		t.Fatalf("Could not create deribit client, %s", err)
	}

	return cli
}

func TestAuth(t *testing.T) {
	deribit := testNewDeribitRestPrivateClient(t)

	_, err := deribit.Auth(context.TODO(), auth.AuthParams{
		GrantType:    "client_credentials",
		ClientID:     deribit.key,
		ClientSecret: deribit.secret,
	})
	assert.Nil(t, err)
}

func TestGetAccountSummary(t *testing.T) {
	deribit := testNewDeribitRestPrivateClient(t)

	_, err := deribit.GetAccountSummary(context.TODO(), account.GetAccountSummaryParams{
		Currency: "USDC",
	})
	assert.Nil(t, err)
}

func TestGetPositions(t *testing.T) {
	deribit := testNewDeribitRestPrivateClient(t)

	_, err := deribit.GetPositions(context.TODO(), account.GetPositionsParams{
		Currency: "BTC",
		Kind:     "future",
	})
	assert.Nil(t, err)
}
