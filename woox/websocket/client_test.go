package websocket

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWooXClient(t *testing.T) {
	cli, err := NewWooXClient(&WooXStreamCfg{
		BasePath:      PublicProdBaseEndpoint,
		ApplicationID: "",
		Debug:         true,
	})
	assert.Nil(t, err)

	err = cli.send(&Request{
		ID:    genClientID(),
		Topic: "SPOT_WOO_USDT@orderbook",
		Event: SUBSCRIBE,
	})
	assert.Nil(t, err)

	time.Sleep(1 * time.Minute)
}
