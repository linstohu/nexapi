package pmrest

import (
	"context"
	"os"
	"testing"

	pmutils "github.com/linstohu/nexapi/binance/portfoliomargin/utils"
	"github.com/stretchr/testify/assert"
)

func testNewAccountClient(t *testing.T) *PortfolioMarginAccountClient {
	cli, err := NewPortfolioMarginAccountClient(&pmutils.PortfolioMarginClientCfg{
		BaseURL: pmutils.PortfolioMarginBaseURL,
		Key:     os.Getenv("BINANCE_KEY"),
		Secret:  os.Getenv("BINANCE_SECRET"),
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create binance client, %s", err)
	}

	return cli
}

func TestGetBalance(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetBalance(context.TODO())
	assert.Nil(t, err)
}
