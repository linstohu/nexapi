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

package websocketuserdata

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/chuckpreslar/emission"
	"github.com/go-playground/validator"
	"github.com/gorilla/websocket"
	eoutils "github.com/linstohu/nexapi/binance/europeanoptions/utils"
)

type OptionsUserDataStreamClient struct {
	// debug mode
	debug bool
	// logger
	logger *slog.Logger

	baseURL     string
	key, secret string

	stopCtx context.Context
	cancel  context.CancelFunc

	conn        *websocket.Conn
	mu          sync.RWMutex
	isConnected bool

	autoReconnect bool
	disconnect    chan struct{}
	heartCancel   chan struct{}

	emitter *emission.Emitter
}

type OptionsUserDataStreamCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL       string `validate:"required"`
	Key           string `validate:"required"`
	Secret        string `validate:"required"`
	AutoReconnect bool   `validate:"required"`
}

func NewUserDataStreamClient(cfg *OptionsUserDataStreamCfg) (*OptionsUserDataStreamClient, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, err
	}

	cli := &OptionsUserDataStreamClient{
		debug:  cfg.Debug,
		logger: cfg.Logger,

		baseURL: cfg.BaseURL,
		key:     cfg.Key,
		secret:  cfg.Secret,

		autoReconnect: cfg.AutoReconnect,

		emitter: emission.NewEmitter(),
	}

	if cli.logger == nil {
		cli.logger = slog.Default()
	}

	return cli, nil
}

func (o *OptionsUserDataStreamClient) Open() error {
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

func (o *OptionsUserDataStreamClient) Close() error {
	if o.stopCtx == nil {
		return fmt.Errorf("%s: ws is not open", logPrefix)
	}

	o.cancel()

	return nil
}

func (o *OptionsUserDataStreamClient) start() error {
	o.conn = nil
	o.setIsConnected(false)
	o.disconnect = make(chan struct{})
	o.heartCancel = make(chan struct{})

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

	if o.autoReconnect {
		go o.reconnect()
	}

	go o.heartbeat()

	go o.readMessages()

	return nil
}

func (o *OptionsUserDataStreamClient) connect() (*websocket.Conn, *http.Response, error) {
	listenKey, err := o.genListenKey()
	if err != nil {
		return nil, nil, err
	}

	baseURL := fmt.Sprintf("%s%s%s", o.baseURL, UserDataStreamRouter, listenKey)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, baseURL, nil)
	if err == nil {
		conn.SetReadLimit(32768 * 64)
	}

	return conn, resp, err
}

func (o *OptionsUserDataStreamClient) genListenKey() (string, error) {
	cli, err := newHttpAuthClient(&httpAuthClientCfg{
		Debug:      o.debug,
		Logger:     o.logger,
		BaseURL:    eoutils.OptionsBaseURL,
		Key:        o.key,
		Secret:     o.secret,
		RecvWindow: 5000,
	})
	if err != nil {
		return "", err
	}

	resp, err := cli.genListenKey(context.TODO())
	if err != nil {
		return "", err
	}

	if resp.Body == nil {
		return "", fmt.Errorf("unknown error")
	}

	return resp.Body.ListenKey, nil
}

func (o *OptionsUserDataStreamClient) updateListenKey() error {
	cli, err := newHttpAuthClient(&httpAuthClientCfg{
		Debug:      o.debug,
		Logger:     o.logger,
		BaseURL:    eoutils.OptionsBaseURL,
		Key:        o.key,
		Secret:     o.secret,
		RecvWindow: 5000,
	})
	if err != nil {
		return err
	}

	err = cli.updateListenKey(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func (o *OptionsUserDataStreamClient) reconnect() {
	<-o.disconnect

	o.setIsConnected(false)

	close(o.heartCancel)

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
func (o *OptionsUserDataStreamClient) close() error {
	close(o.disconnect)

	err := o.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

// setIsConnected sets state for isConnected
func (o *OptionsUserDataStreamClient) setIsConnected(state bool) {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.isConnected = state
}

// IsConnected returns the WebSocket connection state
func (o *OptionsUserDataStreamClient) IsConnected() bool {
	o.mu.RLock()
	defer o.mu.RUnlock()

	return o.isConnected
}

func (o *OptionsUserDataStreamClient) heartbeat() {
	t := time.NewTicker(50 * time.Minute)
	for {
		select {
		case <-t.C:
			err := o.updateListenKey()
			if err != nil {
				o.logger.Info(fmt.Sprintf("%s: update listen-key error, %s", logPrefix, err.Error()))
			}
		case <-o.heartCancel:
			return
		}
	}
}

func (o *OptionsUserDataStreamClient) readMessages() {
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
			_, bytes, err := o.conn.ReadMessage()
			if err != nil {
				o.logger.Info(fmt.Sprintf("read message error, %s", err))

				if err := o.close(); err != nil {
					o.logger.Error(fmt.Sprintf("%s: connection closed error, %s", logPrefix, err.Error()))
					return
				}

				o.logger.Info(fmt.Sprintf("%s: connection closed success", logPrefix))
				return
			}

			err = o.handle(bytes)
			if err != nil {
				o.logger.Info(fmt.Sprintf("%s: handle message error: %s", logPrefix, err.Error()))
			}
		}
	}
}
