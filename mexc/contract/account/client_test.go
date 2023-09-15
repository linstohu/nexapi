package account

import (
	"context"
	"os"
	"testing"

	ctutils "github.com/linstohu/nexapi/mexc/contract/utils"
	"github.com/stretchr/testify/assert"
)

func testNewAccountClient(t *testing.T) *ContractAccountClient {
	cli, err := NewContractAccountClient(&ctutils.ContractClientCfg{
		BaseURL: ctutils.BaseURL,
		Key:     os.Getenv("MEXC_KEY"),
		Secret:  os.Getenv("MEXC_SECRET"),
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create mexc client, %s", err)
	}

	return cli
}

func TestGetAccountAsset(t *testing.T) {
	cli := testNewAccountClient(t)

	_, err := cli.GetAccountAsset(context.TODO())
	assert.Nil(t, err)
}
