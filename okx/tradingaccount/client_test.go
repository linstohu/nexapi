package publicdata

import (
	"context"
	"os"
	"testing"

	"github.com/linstohu/nexapi/okx/tradingaccount/types"
	okxutils "github.com/linstohu/nexapi/okx/utils"
	"github.com/stretchr/testify/assert"
)

func testNewTradingAccountClient(t *testing.T) *TradingAccountClient {
	cli, err := NewTradingAccountClient(&TradingAccountClientCfg{
		Debug:      true,
		BaseURL:    okxutils.RestURL,
		Key:        os.Getenv("OKX_KEY"),
		Secret:     os.Getenv("OKX_SECRET"),
		Passphrase: os.Getenv("OKX_PASS"),
	})

	if err != nil {
		t.Fatalf("Could not create okx private client, %s", err)
	}

	return cli
}

func TestGetBalance(t *testing.T) {
	cli := testNewTradingAccountClient(t)

	_, err := cli.GetBalance(context.TODO(), types.GetBalanceParam{})
	assert.Nil(t, err)
}
