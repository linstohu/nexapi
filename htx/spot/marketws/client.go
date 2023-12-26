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

package marketws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/chuckpreslar/emission"
	"github.com/go-playground/validator"
	"github.com/gorilla/websocket"
	htxutils "github.com/linstohu/nexapi/htx/utils"
	cmap "github.com/orcaman/concurrent-map/v2"
)

type MarketWsClient struct {
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

type MarketWsClientCfg struct {
	Debug         bool
	BaseURL       string `validate:"required"`
	AutoReconnect bool   `validate:"required"`

	Logger *slog.Logger
}

func NewMarketWsClient(cfg *MarketWsClientCfg) (*MarketWsClient, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, err
	}

	cli := &MarketWsClient{
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

func (m *MarketWsClient) Open() error {
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

func (m *MarketWsClient) Close() error {
	if m.stopCtx == nil {
		return fmt.Errorf("%s: ws is not open", logPrefix)
	}

	m.cancel()

	return nil
}

func (m *MarketWsClient) start() error {
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

func (m *MarketWsClient) connect() (*websocket.Conn, *http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, m.baseURL, nil)
	if err == nil {
		conn.SetReadLimit(32768 * 64)
	}

	return conn, resp, err
}

func (m *MarketWsClient) reconnect() {
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
func (m *MarketWsClient) close() error {
	close(m.disconnect)

	err := m.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

// setIsConnected sets state for isConnected
func (m *MarketWsClient) setIsConnected(state bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.isConnected = state
}

// IsConnected returns the WebSocket connection state
func (m *MarketWsClient) IsConnected() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.isConnected
}

func (m *MarketWsClient) readMessages() {
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
			msgType, buf, err := m.conn.ReadMessage()
			if err != nil {
				m.logger.Info(fmt.Sprintf("%s: read message error, %s", logPrefix, err))
				time.Sleep(TimerIntervalSecond * time.Second)
				continue
			}

			// decompress gzip data if it is binary message
			if msgType == websocket.BinaryMessage {
				message, err := htxutils.GZipDecompress(buf)
				if err != nil {
					m.logger.Info(fmt.Sprintf("%s: ungzip data error: %s", logPrefix, err))
					m.logger.Info(fmt.Sprintf("%s: ready to close...", logPrefix))

					if err := m.close(); err != nil {
						m.logger.Error(fmt.Sprintf("%s: connection closed error, %s", logPrefix, err.Error()))
						return
					}

					m.logger.Info(fmt.Sprintf("%s: connection closed success", logPrefix))
					return
				}

				var msg AnyMessage

				err = json.Unmarshal([]byte(message), &msg)
				if err != nil {
					m.logger.Info(fmt.Sprintf("%s: read object error, %s", logPrefix, err))
					m.logger.Info(fmt.Sprintf("%s: ready to close...", logPrefix))

					if err := m.close(); err != nil {
						m.logger.Error(fmt.Sprintf("%s: connection closed error, %s", logPrefix, err.Error()))
						return
					}

					m.logger.Info(fmt.Sprintf("%s: connection closed success", logPrefix))
					return
				}

				switch {
				case msg.Ping != nil:
					err := m.pong(&PongMessage{
						Pong: msg.Ping.Ping,
					})
					if err != nil {
						m.logger.Error(fmt.Sprintf("%s: handle ping error: %s", logPrefix, err.Error()))
					}
				case msg.Response != nil:
					// todo
				case msg.SubscribedMessage != nil:
					err := m.handle(msg.SubscribedMessage)
					if err != nil {
						m.logger.Error(fmt.Sprintf("%s: handle message error: %s", logPrefix, err.Error()))
					}
				}
			}
		}
	}
}

func (m *MarketWsClient) resubscribe() error {
	topics := m.subscriptions.Keys()

	if len(topics) == 0 {
		return nil
	}

	redo := make([]string, 0)

	for _, v := range topics {
		// do subscription
		err := m.send(&Request{
			ID:  fmt.Sprintf("%v", rand.Uint32()),
			Sub: v,
		})

		if err != nil {
			redo = append(redo, v)
			continue
		}
	}

	if len(redo) != 0 {
		return fmt.Errorf("resubscribe error: %s", strings.Join(redo, ","))
	}

	return nil
}

func (m *MarketWsClient) subscribe(topic string) error {
	if m.subscriptions.Has(topic) {
		return nil
	}

	// do subscription

	err := m.send(&Request{
		ID:  fmt.Sprintf("%v", rand.Uint32()),
		Sub: topic,
	})

	if err != nil {
		return err
	}

	m.subscriptions.Set(topic, struct{}{})

	return nil
}

func (m *MarketWsClient) unsubscribe(topic string) error {
	err := m.send(&Request{
		ID:    fmt.Sprintf("%v", rand.Uint32()),
		UnSub: topic,
	})

	if err != nil {
		return err
	}

	m.subscriptions.Remove(topic)

	return nil
}

func (m *MarketWsClient) send(req *Request) error {
	m.sending.Lock()

	// Rate Limit: https://www.htx.com/en-us/opend/newApiPages/?id=662
	defer time.Sleep(100 * time.Millisecond)
	defer m.sending.Unlock()

	if !m.IsConnected() {
		return errors.New("connection is closed")
	}

	return m.conn.WriteJSON(req)
}

func (m *MarketWsClient) pong(ping *PongMessage) error {
	m.sending.Lock()

	// Rate Limit: https://www.htx.com/en-us/opend/newApiPages/?id=662
	defer time.Sleep(100 * time.Millisecond)
	defer m.sending.Unlock()

	if !m.IsConnected() {
		return errors.New("connection is closed")
	}

	return m.conn.WriteJSON(ping)
}
