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

type SpotMarketStreamClient struct {
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

type SpotMarketStreamCfg struct {
	Debug         bool
	BaseURL       string `validate:"required"`
	AutoReconnect bool   `validate:"required"`

	Logger *slog.Logger
}

func NewSpotMarketStreamClient(cfg *SpotMarketStreamCfg) (*SpotMarketStreamClient, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, err
	}

	cli := &SpotMarketStreamClient{
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

func (m *SpotMarketStreamClient) Open() error {
	if m.stopCtx != nil {
		return fmt.Errorf("%s: ws is already open", logPrefix)
	}

	m.stopCtx, m.cancel = context.WithCancel(context.Background())

	err := m.start()
	if err != nil {
		return err
	}

	return nil
}

func (m *SpotMarketStreamClient) Close() error {
	if m.stopCtx == nil {
		return fmt.Errorf("%s: ws is not open", logPrefix)
	}

	m.cancel()

	return nil
}

func (m *SpotMarketStreamClient) start() error {
	m.conn = nil
	m.setIsConnected(false)
	m.disconnect = make(chan struct{})

	for i := 0; i < MaxTryTimes; i++ {
		conn, _, err := m.connect()
		if err != nil {
			m.logger.Info(fmt.Sprintf("%s: connect error, times(%v), error: %s", logPrefix, i, err.Error()))
			tm := (i + 1) * 5
			time.Sleep(time.Duration(tm) * time.Second)
			continue
		}
		m.conn = conn
		break
	}
	if m.conn == nil {
		return errors.New("connect failed")
	}

	m.logger.Info(fmt.Sprintf("%s: connect success, base_url: %s", logPrefix, m.baseURL))

	m.setIsConnected(true)

	m.resubscribe()

	if m.autoReconnect {
		go m.reconnect()
	}

	go m.readMessages()

	return nil
}

func (m *SpotMarketStreamClient) connect() (*websocket.Conn, *http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, m.baseURL+CombinedStreamRouter, nil)
	if err == nil {
		conn.SetReadLimit(32768 * 64)
	}

	return conn, resp, err
}

func (m *SpotMarketStreamClient) reconnect() {
	<-m.disconnect

	m.setIsConnected(false)

	time.Sleep(1 * time.Second)

	select {
	case <-m.stopCtx.Done():
		m.logger.Info(fmt.Sprintf("%s: reconnection exits", logPrefix))
		return
	default:
		m.logger.Info(fmt.Sprintf("%s: try to reconnect...", logPrefix))
		m.start()
	}
}

// close closes the websocket connection
func (m *SpotMarketStreamClient) close() error {
	close(m.disconnect)

	err := m.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

// setIsConnected sets state for isConnected
func (m *SpotMarketStreamClient) setIsConnected(state bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.isConnected = state
}

// IsConnected returns the WebSocket connection state
func (m *SpotMarketStreamClient) IsConnected() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.isConnected
}

func (m *SpotMarketStreamClient) readMessages() {
	for {
		select {
		case <-m.stopCtx.Done():
			m.logger.Info(fmt.Sprintf("%s: ready to close...", logPrefix))

			if err := m.close(); err != nil {
				m.logger.Error(fmt.Sprintf("%s: connection closed error, %s", logPrefix, err.Error()))
				return
			}

			m.logger.Info(fmt.Sprintf("%s: connection closed success", logPrefix))
			return
		default:
			var msg utils.AnyMessage
			err := m.conn.ReadJSON(&msg)
			if err != nil {
				m.logger.Info(fmt.Sprintf("%s: read message error, %s", logPrefix, err))

				if err := m.close(); err != nil {
					m.logger.Error(fmt.Sprintf("%s: connection closed error, %s", logPrefix, err.Error()))
					return
				}

				m.logger.Info(fmt.Sprintf("%s: connection closed success", logPrefix))
				return
			}

			switch {
			case msg.Response != nil:
				// todo
			case msg.SubscribedMessage != nil:
				err := m.handle(msg.SubscribedMessage)
				if err != nil {
					m.logger.Info(fmt.Sprintf("%s: handle message error: %s", logPrefix, err.Error()))
				}
			}
		}
	}
}

func (m *SpotMarketStreamClient) resubscribe() error {
	topics := m.subscriptions.Keys()

	if len(topics) == 0 {
		return nil
	}

	// do subscription
	err := m.send(&utils.Request{
		ID:     rand.Uint32(),
		Method: SUBSCRIBE,
		Params: topics,
	})

	if err != nil {
		return err
	}

	return nil
}

func (m *SpotMarketStreamClient) subscribe(topics []string) error {
	ts := make([]string, 0)

	for _, topic := range topics {
		if m.subscriptions.Has(topic) {
			continue
		}
		ts = append(ts, topic)
	}

	if len(ts) == 0 {
		return nil
	}

	// do subscription
	err := m.send(&utils.Request{
		ID:     rand.Uint32(),
		Method: SUBSCRIBE,
		Params: ts,
	})

	if err != nil {
		return err
	}

	for _, v := range ts {
		m.subscriptions.Set(v, struct{}{})
	}

	return nil
}

func (m *SpotMarketStreamClient) unsubscribe(topics []string) error {
	err := m.send(&utils.Request{
		ID:     rand.Uint32(),
		Method: UNSUBSCRIBE,
		Params: topics,
	})

	if err != nil {
		return err
	}

	for _, v := range topics {
		m.subscriptions.Remove(v)
	}

	return nil
}

func (m *SpotMarketStreamClient) send(req *utils.Request) error {
	m.sending.Lock()
	defer m.sending.Unlock()

	if !m.IsConnected() {
		return errors.New("connection is closed")
	}

	return m.conn.WriteJSON(req)
}
