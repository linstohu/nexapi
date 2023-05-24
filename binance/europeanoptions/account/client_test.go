package account

import (
	"context"
	"os"
	"testing"

	"github.com/linstohu/nexapi/binance/europeanoptions/account/types"
	"github.com/linstohu/nexapi/binance/europeanoptions/utils"
	"github.com/stretchr/testify/assert"
)

func testNewAccountClient(t *testing.T) *OptionsAccountClient {
	cli, err := NewOptionsAccountClient(&utils.OptionsClientCfg{
		BaseURL: utils.OptionsBaseURL,
		Key:     os.Getenv("BINANCE_KEY"),
		Secret:  os.Getenv("BINANCE_SECRET"),
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create binance client, %s", err)
	}

	return cli
}

func TestGetAccountInfo(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetAccountInfo(context.TODO())
	assert.Nil(t, err)
}

func TestNewOrder(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.NewOrder(context.TODO(), types.NewOrderParam{
		Symbol:      "ETH-230609-2100-C",
		Side:        utils.BuySide,
		Type:        utils.Limit,
		Quantity:    1,
		Price:       4.5,
		TimeInForce: utils.GTC,
	})
	assert.Nil(t, err)
}

func TestGetSingleOrder(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetSingleOrder(context.TODO(), types.GetSingleOrderParam{
		Symbol:  "ETH-230609-2100-C",
		OrderID: 0,
	})
	assert.Nil(t, err)
}

func TestCancelOrder(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.CancelOrder(context.TODO(), types.CancelOrderParam{
		Symbol:  "ETH-230609-2100-C",
		OrderID: 0,
	})
	assert.Nil(t, err)
}

func TestCancelAllOrdersBySymbol(t *testing.T) {
	cli := testNewAccountClient(t)

	err := cli.CancelAllOrdersBySymbol(context.TODO(), types.CancelAllOrdersParam{
		Symbol: "ETH-230609-2100-C",
	})
	assert.Nil(t, err)
}

func TestCancelAllOrdersByUnderlying(t *testing.T) {
	cli := testNewAccountClient(t)

	err := cli.CancelAllOrdersByUnderlying(context.TODO(), types.CancelAllOrdersByUnderlyingParam{
		Underlying: "ETHUSDT",
	})
	assert.Nil(t, err)
}

func TestGetOpenOrders(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetOpenOrders(context.TODO(), types.GetCurrentOpenOrdersParam{
		Symbol: "ETH-230609-2100-C",
	})
	assert.Nil(t, err)
}

func TestGetOrderHistory(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetOrderHistory(context.TODO(), types.GetOrderHistoryParam{
		Symbol: "ETH-230609-2100-C",
	})
	assert.Nil(t, err)
}

func TestGetPositionInfo(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetPositionInfo(context.TODO(), types.GetPositionInfoParam{
		Symbol: "ETH-230609-2100-C",
	})
	assert.Nil(t, err)
}

func TestGetTradeList(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetTradeList(context.TODO(), types.GetTradeListParam{})
	assert.Nil(t, err)
}

func TestGetExerciseRecord(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetExerciseRecord(context.TODO(), types.GetExerciseRecordParam{})
	assert.Nil(t, err)
}

func TestGetFundingFlow(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetFundingFlow(context.TODO(), types.GetFundingFlowParam{
		Currency: "USDT",
	})
	assert.Nil(t, err)
}
