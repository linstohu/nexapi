package rest

import (
	"context"
	"os"
	"testing"

	"github.com/linstohu/nexapi/deribit/rest/types/account"
	"github.com/linstohu/nexapi/deribit/rest/types/auth"
	"github.com/linstohu/nexapi/deribit/rest/types/trading"
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

func TestGetTransactionLog(t *testing.T) {
	deribit := testNewDeribitRestPrivateClient(t)

	_, err := deribit.GetTransactionLog(context.TODO(), account.GetTransactionLogParams{
		Currency:       "USDC",
		StartTimestamp: 1683743286000,
		EndTimestamp:   1696976886000,
		Query:          "buy",
		Count:          30,
		Continuation:   0,
	})
	assert.Nil(t, err)
}

func TestBuy(t *testing.T) {
	deribit := testNewDeribitRestPrivateClient(t)

	_, err := deribit.Buy(context.TODO(), trading.BuyParams{
		InstrumentName: "BTC_USDC",
		Amount:         0.0007,
		Type:           "limit",
		Price:          25000,
	})
	assert.Nil(t, err)
}

func TestSell(t *testing.T) {
	deribit := testNewDeribitRestPrivateClient(t)

	_, err := deribit.Sell(context.TODO(), trading.SellParams{
		InstrumentName: "BTC_USDC",
		Amount:         0.0011,
		Type:           "limit",
		Price:          30000,
	})
	assert.Nil(t, err)
}

func TestCancel(t *testing.T) {
	deribit := testNewDeribitRestPrivateClient(t)

	_, err := deribit.Cancel(context.TODO(), trading.CancelParams{
		OrderID: "",
	})
	assert.Nil(t, err)
}

func TestCancelAll(t *testing.T) {
	deribit := testNewDeribitRestPrivateClient(t)

	_, err := deribit.CancelAll(context.TODO())
	assert.Nil(t, err)
}

func TestCancelAllByInstrument(t *testing.T) {
	deribit := testNewDeribitRestPrivateClient(t)

	_, err := deribit.CancelAllByInstrument(context.TODO(), trading.CancelAllByInstrumentParams{
		InstrumentName: "BTC_USDC",
	})
	assert.Nil(t, err)
}

func TestClosePosition(t *testing.T) {
	deribit := testNewDeribitRestPrivateClient(t)

	_, err := deribit.ClosePosition(context.TODO(), trading.ClosePositionParams{
		InstrumentName: "BTC-PERPETUAL",
		Type:           "market",
	})
	assert.Nil(t, err)
}

func TestGetOpenOrdersByCurrency(t *testing.T) {
	deribit := testNewDeribitRestPrivateClient(t)

	_, err := deribit.GetOpenOrdersByCurrency(context.TODO(), trading.GetOpenOrdersByCurrencyParams{
		Currency: "BTC",
		Kind:     "spot",
	})
	assert.Nil(t, err)
}

func TestGetOrderState(t *testing.T) {
	deribit := testNewDeribitRestPrivateClient(t)

	_, err := deribit.GetOrderState(context.TODO(), trading.GetOrderStateParams{
		OrderID: "",
	})
	assert.Nil(t, err)
}

func TestGetUserTradesByCurrency(t *testing.T) {
	deribit := testNewDeribitRestPrivateClient(t)

	_, err := deribit.GetUserTradesByCurrency(context.TODO(), trading.GetUserTradesByCurrencyParams{
		Currency: "BTC",
		Kind:     "spot",
	})
	assert.Nil(t, err)
}

func TestGetSettlementHistoryByCurrency(t *testing.T) {
	deribit := testNewDeribitRestPrivateClient(t)

	_, err := deribit.GetSettlementHistoryByCurrency(context.TODO(), trading.GetSettlementHistoryByCurrencyParams{
		Currency: "BTC",
		Type:     "delivery",
	})
	assert.Nil(t, err)
}
