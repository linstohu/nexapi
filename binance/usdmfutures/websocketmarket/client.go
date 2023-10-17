/*
 * Copyright (c) 2023, LinstoHu
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package websocketmarket

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/chuckpreslar/emission"
	"github.com/go-playground/validator"
	"github.com/gorilla/websocket"
	"github.com/linstohu/nexapi/binance/utils"
	cmap "github.com/orcaman/concurrent-map/v2"
)

type USDMarginedMarketStreamClient struct {
	baseURL string
	// debug mode
	debug bool
	// logger
	logger *slog.Logger

	ctx         context.Context
	conn        *websocket.Conn
	mu          sync.RWMutex
	isConnected bool

	autoReconnect bool
	disconnect    chan struct{}

	sending       sync.Mutex
	subscriptions cmap.ConcurrentMap[string, struct{}]

	emitter *emission.Emitter
}

type USDMarginedMarketStreamCfg struct {
	BaseURL string `validate:"required"`
	Debug   bool
	// Logger
	Logger *slog.Logger
}

func NewMarketStreamClient(ctx context.Context, cfg *USDMarginedMarketStreamCfg) (*USDMarginedMarketStreamClient, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, err
	}

	cli := &USDMarginedMarketStreamClient{
		baseURL: cfg.BaseURL,
		debug:   cfg.Debug,
		logger:  cfg.Logger,

		ctx:           ctx,
		autoReconnect: true,

		subscriptions: cmap.New[struct{}](),
		emitter:       emission.NewEmitter(),
	}

	if cli.logger == nil {
		cli.logger = slog.Default()
	}

	err := cli.start()
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func (u *USDMarginedMarketStreamClient) start() error {
	u.conn = nil
	u.setIsConnected(false)
	u.disconnect = make(chan struct{})

	for i := 0; i < MaxTryTimes; i++ {
		conn, _, err := u.connect()
		if err != nil {
			u.logger.Info(fmt.Sprintf("connect error, times(%v), error: %s", i, err.Error()))
			tm := (i + 1) * 5
			time.Sleep(time.Duration(tm) * time.Second)
			continue
		}
		u.conn = conn
		break
	}
	if u.conn == nil {
		return errors.New("connect failed")
	}

	u.setIsConnected(true)

	u.resubscribe()

	if u.autoReconnect {
		go u.reconnect()
	}

	go u.readMessages()

	return nil
}

func (u *USDMarginedMarketStreamClient) connect() (*websocket.Conn, *http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, u.baseURL+CombinedStreamRouter, nil)
	if err == nil {
		conn.SetReadLimit(32768 * 64)
	}

	return conn, resp, err
}

func (u *USDMarginedMarketStreamClient) reconnect() {
	<-u.disconnect

	u.setIsConnected(false)

	u.logger.Info("disconnect, then reconnect...")

	time.Sleep(1 * time.Second)

	select {
	case <-u.ctx.Done():
		u.logger.Info(fmt.Sprintf("never reconnect, %s", u.ctx.Err()))
		return
	default:
		u.start()
	}
}

// close closes the websocket connection
func (u *USDMarginedMarketStreamClient) close() error {
	close(u.disconnect)

	err := u.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

// setIsConnected sets state for isConnected
func (u *USDMarginedMarketStreamClient) setIsConnected(state bool) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.isConnected = state
}

// IsConnected returns the WebSocket connection state
func (u *USDMarginedMarketStreamClient) IsConnected() bool {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return u.isConnected
}

func (u *USDMarginedMarketStreamClient) readMessages() {
	for {
		select {
		case <-u.ctx.Done():
			u.logger.Info(fmt.Sprintf("context done, error: %s", u.ctx.Err().Error()))

			if err := u.close(); err != nil {
				u.logger.Info(fmt.Sprintf("websocket connection closed error, %s", err.Error()))
			}

			return
		default:
			var msg utils.AnyMessage
			err := u.conn.ReadJSON(&msg)
			if err != nil {
				u.logger.Info(fmt.Sprintf("read object error, %s", err))

				if err := u.close(); err != nil {
					u.logger.Info(fmt.Sprintf("websocket connection closed error, %s", err.Error()))
				}

				return
			}

			switch {
			case msg.Response != nil:
				// todo
			case msg.SubscribedMessage != nil:
				err := u.handle(msg.SubscribedMessage)
				if err != nil {
					u.logger.Info(fmt.Sprintf("handle message error: %s", err.Error()))
				}
			}
		}
	}
}

func (u *USDMarginedMarketStreamClient) resubscribe() error {
	topics := u.subscriptions.Keys()

	if len(topics) == 0 {
		return nil
	}

	// do subscription
	err := u.send(&utils.Request{
		ID:     rand.Uint32(),
		Method: SUBSCRIBE,
		Params: topics,
	})

	if err != nil {
		return err
	}

	return nil
}

func (u *USDMarginedMarketStreamClient) subscribe(topics []string) error {
	ts := make([]string, 0)

	for _, topic := range topics {
		if u.subscriptions.Has(topic) {
			continue
		}
		ts = append(ts, topic)
	}

	if len(ts) == 0 {
		return nil
	}

	// do subscription
	err := u.send(&utils.Request{
		ID:     rand.Uint32(),
		Method: SUBSCRIBE,
		Params: ts,
	})

	if err != nil {
		return err
	}

	for _, v := range ts {
		u.subscriptions.Set(v, struct{}{})
	}

	return nil
}

func (u *USDMarginedMarketStreamClient) unsubscribe(topics []string) error {
	err := u.send(&utils.Request{
		ID:     rand.Uint32(),
		Method: UNSUBSCRIBE,
		Params: topics,
	})

	if err != nil {
		return err
	}

	for _, v := range topics {
		u.subscriptions.Remove(v)
	}

	return nil
}

func (u *USDMarginedMarketStreamClient) send(req *utils.Request) error {
	u.sending.Lock()
	defer u.sending.Unlock()

	if !u.IsConnected() {
		return errors.New("connection is closed")
	}

	return u.conn.WriteJSON(req)
}
