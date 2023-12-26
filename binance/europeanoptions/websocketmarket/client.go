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

type OptionsMarketStreamClient struct {
	baseURL string
	// debug mode
	debug bool
	// logger
	logger *slog.Logger

	stopCtx context.Context
	cancel  context.CancelFunc

	conn        *websocket.Conn
	mu          sync.RWMutex
	isConnected bool

	autoReconnect bool
	disconnect    chan struct{}

	sending       sync.Mutex
	subscriptions cmap.ConcurrentMap[string, struct{}]

	emitter *emission.Emitter
}

type OptionsMarketStreamCfg struct {
	Debug         bool
	BaseURL       string `validate:"required"`
	AutoReconnect bool   `validate:"required"`

	Logger *slog.Logger
}

func NewMarketStreamClient(cfg *OptionsMarketStreamCfg) (*OptionsMarketStreamClient, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, err
	}

	cli := &OptionsMarketStreamClient{
		baseURL: cfg.BaseURL,
		debug:   cfg.Debug,
		logger:  cfg.Logger,

		autoReconnect: cfg.AutoReconnect,

		subscriptions: cmap.New[struct{}](),
		emitter:       emission.NewEmitter(),
	}

	if cli.logger == nil {
		cli.logger = slog.Default()
	}

	return cli, nil
}

func (o *OptionsMarketStreamClient) Open() error {
	if o.stopCtx != nil {
		return fmt.Errorf("%s: ws is already open", logPrefix)
	}

	o.stopCtx, o.cancel = context.WithCancel(context.Background())

	err := o.start()
	if err != nil {
		return err
	}

	return nil
}

func (o *OptionsMarketStreamClient) Close() error {
	if o.stopCtx == nil {
		return fmt.Errorf("%s: ws is not open", logPrefix)
	}

	o.cancel()

	return nil
}

func (o *OptionsMarketStreamClient) start() error {
	o.conn = nil
	o.setIsConnected(false)
	o.disconnect = make(chan struct{})

	for i := 0; i < MaxTryTimes; i++ {
		conn, _, err := o.connect()
		if err != nil {
			o.logger.Info(fmt.Sprintf("%s: connect error, times(%v), error: %s", logPrefix, i, err.Error()))
			tm := (i + 1) * 5
			time.Sleep(time.Duration(tm) * time.Second)
			continue
		}
		o.conn = conn
		break
	}
	if o.conn == nil {
		return errors.New("connect failed")
	}

	o.logger.Info(fmt.Sprintf("%s: connect success, base_url: %s", logPrefix, o.baseURL))

	o.setIsConnected(true)

	o.resubscribe()

	if o.autoReconnect {
		go o.reconnect()
	}

	go o.readMessages()

	return nil
}

func (o *OptionsMarketStreamClient) connect() (*websocket.Conn, *http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, o.baseURL+CombinedStreamRouter, nil)
	if err == nil {
		conn.SetReadLimit(32768 * 64)
	}

	return conn, resp, err
}

func (o *OptionsMarketStreamClient) reconnect() {
	<-o.disconnect

	o.setIsConnected(false)

	time.Sleep(1 * time.Second)

	select {
	case <-o.stopCtx.Done():
		o.logger.Info(fmt.Sprintf("%s: reconnection exits", logPrefix))
		return
	default:
		o.logger.Info(fmt.Sprintf("%s: try to reconnect...", logPrefix))
		o.start()
	}
}

// close closes the websocket connection
func (o *OptionsMarketStreamClient) close() error {
	close(o.disconnect)

	err := o.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

// setIsConnected sets state for isConnected
func (o *OptionsMarketStreamClient) setIsConnected(state bool) {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.isConnected = state
}

// IsConnected returns the WebSocket connection state
func (o *OptionsMarketStreamClient) IsConnected() bool {
	o.mu.RLock()
	defer o.mu.RUnlock()

	return o.isConnected
}

func (o *OptionsMarketStreamClient) readMessages() {
	for {
		select {
		case <-o.stopCtx.Done():
			o.logger.Info(fmt.Sprintf("%s: ready to close...", logPrefix))

			if err := o.close(); err != nil {
				o.logger.Error(fmt.Sprintf("%s: connection closed error, %s", logPrefix, err.Error()))
				return
			}

			o.logger.Info(fmt.Sprintf("%s: connection closed success", logPrefix))
			return
		default:
			var msg utils.AnyMessage
			err := o.conn.ReadJSON(&msg)
			if err != nil {
				o.logger.Info(fmt.Sprintf("read message error, %s", err))

				if err := o.close(); err != nil {
					o.logger.Error(fmt.Sprintf("%s: connection closed error, %s", logPrefix, err.Error()))
					return
				}

				o.logger.Info(fmt.Sprintf("%s: connection closed success", logPrefix))
				return
			}

			switch {
			case msg.Response != nil:
				// todo
			case msg.SubscribedMessage != nil:
				err := o.handle(msg.SubscribedMessage)
				if err != nil {
					o.logger.Info(fmt.Sprintf("%s: handle message error: %s", logPrefix, err.Error()))
				}
			}
		}
	}
}

func (o *OptionsMarketStreamClient) resubscribe() error {
	topics := o.subscriptions.Keys()

	if len(topics) == 0 {
		return nil
	}

	// do subscription
	err := o.send(&utils.Request{
		ID:     rand.Uint32(),
		Method: SUBSCRIBE,
		Params: topics,
	})

	if err != nil {
		return err
	}

	return nil
}

func (o *OptionsMarketStreamClient) subscribe(topics []string) error {
	ts := make([]string, 0)

	for _, topic := range topics {
		if o.subscriptions.Has(topic) {
			continue
		}
		ts = append(ts, topic)
	}

	if len(ts) == 0 {
		return nil
	}

	// do subscription
	err := o.send(&utils.Request{
		ID:     rand.Uint32(),
		Method: SUBSCRIBE,
		Params: ts,
	})

	if err != nil {
		return err
	}

	for _, v := range ts {
		o.subscriptions.Set(v, struct{}{})
	}

	return nil
}

func (o *OptionsMarketStreamClient) unsubscribe(topics []string) error {
	err := o.send(&utils.Request{
		ID:     rand.Uint32(),
		Method: UNSUBSCRIBE,
		Params: topics,
	})

	if err != nil {
		return err
	}

	for _, v := range topics {
		o.subscriptions.Remove(v)
	}

	return nil
}

func (o *OptionsMarketStreamClient) send(req *utils.Request) error {
	o.sending.Lock()
	defer o.sending.Unlock()

	if !o.IsConnected() {
		return errors.New("connection is closed")
	}

	return o.conn.WriteJSON(req)
}
