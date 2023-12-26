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

package websocket

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
	"github.com/linstohu/nexapi/woox/websocket/types"
	cmap "github.com/orcaman/concurrent-map/v2"
)

type WooXWebsocketClient struct {
	baseURL                    string
	key, secret, applicationID string
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
	heartCancel   chan struct{}
	disconnect    chan struct{}

	sending       sync.Mutex
	subscriptions cmap.ConcurrentMap[string, struct{}]

	emitter *emission.Emitter
}

type WooXWebsocketCfg struct {
	Debug         bool
	BaseURL       string `validate:"required"`
	AutoReconnect bool   `validate:"required"`

	Key           string
	Secret        string
	ApplicationID string `validate:"required"`

	// Logger
	Logger *slog.Logger
}

func NewWooXWebsocketClient(cfg *WooXWebsocketCfg) (*WooXWebsocketClient, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, err
	}

	cli := &WooXWebsocketClient{
		debug:   cfg.Debug,
		baseURL: cfg.BaseURL,

		key:           cfg.Key,
		secret:        cfg.Secret,
		applicationID: cfg.ApplicationID,

		logger: cfg.Logger,

		autoReconnect: cfg.AutoReconnect,

		subscriptions: cmap.New[struct{}](),
		emitter:       emission.NewEmitter(),
	}

	if cli.logger == nil {
		cli.logger = slog.Default()
	}

	return cli, nil
}

func (w *WooXWebsocketClient) Open() error {
	if w.stopCtx != nil {
		return fmt.Errorf("%s: ws is already open", logPrefix)
	}

	w.stopCtx, w.cancel = context.WithCancel(context.Background())

	err := w.start()
	if err != nil {
		return err
	}

	return nil
}

func (w *WooXWebsocketClient) Close() error {
	if w.stopCtx == nil {
		return fmt.Errorf("%s: ws is not open", logPrefix)
	}

	w.cancel()
	w.stopCtx = nil

	return nil
}

func (w *WooXWebsocketClient) start() error {
	w.conn = nil
	w.setIsConnected(false)
	w.heartCancel = make(chan struct{})
	w.disconnect = make(chan struct{})

	for i := 0; i < MaxTryTimes; i++ {
		conn, _, err := w.connect()
		if err != nil {
			w.logger.Info(fmt.Sprintf("%s: connect error, times(%v), error: %s", logPrefix, i, err.Error()))
			tm := (i + 1) * 5
			time.Sleep(time.Duration(tm) * time.Second)
			continue
		}
		w.conn = conn
		break
	}
	if w.conn == nil {
		return errors.New("connect failed")
	}

	w.logger.Info(fmt.Sprintf("%s: connect success, base_url: %s", logPrefix, w.baseURL))

	w.setIsConnected(true)

	w.resubscribe()

	if w.autoReconnect {
		go w.reconnect()
	}

	go w.heartbeat()

	go w.readMessages()

	return nil
}

func (w *WooXWebsocketClient) connect() (*websocket.Conn, *http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, w.baseURL+w.applicationID, nil)
	if err == nil {
		conn.SetReadLimit(32768 * 64)
	}

	return conn, resp, err
}

func (w *WooXWebsocketClient) reconnect() {
	<-w.disconnect

	w.setIsConnected(false)

	close(w.heartCancel)

	time.Sleep(1 * time.Second)

	select {
	case <-w.stopCtx.Done():
		w.logger.Info(fmt.Sprintf("%s: reconnection exits", logPrefix))
		return
	default:
		w.logger.Info(fmt.Sprintf("%s: try to reconnect...", logPrefix))
		w.start()
	}
}

// close closes the websocket connection
func (w *WooXWebsocketClient) close() error {
	close(w.disconnect)

	err := w.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

// setIsConnected sets state for isConnected
func (w *WooXWebsocketClient) setIsConnected(state bool) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.isConnected = state
}

// IsConnected returns the WebSocket connection state
func (w *WooXWebsocketClient) IsConnected() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.isConnected
}

// heartbeat sends ping every 5s to keep alive
func (w *WooXWebsocketClient) heartbeat() {
	t := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t.C:
			w.send(&types.Request{
				Event: "ping",
			})
		case <-w.heartCancel:
			return
		}
	}
}

func (w *WooXWebsocketClient) readMessages() {
	for {
		select {
		case <-w.stopCtx.Done():
			w.logger.Info(fmt.Sprintf("%s: ready to close...", logPrefix))

			if err := w.close(); err != nil {
				w.logger.Error(fmt.Sprintf("%s: connection closed error, %s", logPrefix, err.Error()))
				return
			}

			w.logger.Info(fmt.Sprintf("%s: connection closed success", logPrefix))
			return
		default:
			var msg types.AnyMessage
			err := w.conn.ReadJSON(&msg)
			if err != nil {
				w.logger.Info(fmt.Sprintf("%s: read message error, %s", logPrefix, err))
				w.logger.Info(fmt.Sprintf("%s: ready to close...", logPrefix))

				if err := w.close(); err != nil {
					w.logger.Error(fmt.Sprintf("%s: connection closed error, %s", logPrefix, err.Error()))
					return
				}

				w.logger.Info(fmt.Sprintf("%s: connection closed success", logPrefix))

				return
			}

			switch {
			case msg.Response != nil:
				// todo
			case msg.SubscribedMessage != nil:
				err := w.handle(msg.SubscribedMessage)
				if err != nil {
					w.logger.Info(fmt.Sprintf("%s: handle message error: %s", logPrefix, err.Error()))
				}
			}
		}
	}
}

func (w *WooXWebsocketClient) resubscribe() error {
	topics := w.subscriptions.Keys()

	if len(topics) == 0 {
		return nil
	}

	// do subscription
	for _, v := range topics {
		err := w.send(&types.Request{
			ID:    fmt.Sprintf("ClientID%d", rand.Intn(100)),
			Topic: v,
			Event: SUBSCRIBE,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (w *WooXWebsocketClient) subscribe(topics []string) error {
	ts := make([]string, 0)

	for _, topic := range topics {
		if w.subscriptions.Has(topic) {
			continue
		}
		ts = append(ts, topic)
	}

	if len(ts) == 0 {
		return nil
	}

	// do subscription
	for _, v := range ts {
		err := w.send(&types.Request{
			ID:    fmt.Sprintf("ClientID%d", rand.Intn(100)),
			Topic: v,
			Event: SUBSCRIBE,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (w *WooXWebsocketClient) unsubscribe(channels []string) error {
	for _, v := range channels {
		err := w.send(&types.Request{
			ID:    fmt.Sprintf("ClientID%d", rand.Intn(100)),
			Topic: v,
			Event: UNSUBSCRIBE,
		})

		if err != nil {
			return err
		}

		w.subscriptions.Remove(v)
	}

	return nil
}

func (w *WooXWebsocketClient) send(req *types.Request) error {
	w.sending.Lock()
	defer w.sending.Unlock()

	if !w.IsConnected() {
		return errors.New("connection is closed")
	}

	return w.conn.WriteJSON(req)
}
