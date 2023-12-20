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

package accountws

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/chuckpreslar/emission"
	"github.com/go-playground/validator"
	"github.com/gorilla/websocket"
	htxutils "github.com/linstohu/nexapi/htx/utils"
	cmap "github.com/orcaman/concurrent-map/v2"
)

type AccountWsClient struct {
	host, path string
	baseURL    string
	// debug mode
	debug bool
	// logger
	logger *slog.Logger

	key, secret string

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

type AccountWsClientCfg struct {
	Debug bool
	// Logger
	Logger  *slog.Logger
	BaseURL string `validate:"required"`
	Key     string `validate:"required"`
	Secret  string `validate:"required"`
}

func NewAccountWsClient(ctx context.Context, cfg *AccountWsClientCfg) (*AccountWsClient, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, err
	}

	baseURL, err := url.Parse(cfg.BaseURL)
	if err != nil {
		return nil, err
	}

	cli := &AccountWsClient{
		debug:   cfg.Debug,
		logger:  cfg.Logger,
		baseURL: cfg.BaseURL,
		host:    baseURL.Host,
		path:    baseURL.Path,

		key:    cfg.Key,
		secret: cfg.Secret,

		ctx:           ctx,
		autoReconnect: true,

		subscriptions: cmap.New[struct{}](),
		emitter:       emission.NewEmitter(),
	}

	if cli.logger == nil {
		cli.logger = slog.Default()
	}

	err = cli.start()
	if err != nil {
		return nil, err
	}

	time.Sleep(100 * time.Millisecond)

	return cli, nil
}

func (m *AccountWsClient) start() error {
	m.conn = nil
	m.setIsConnected(false)
	m.disconnect = make(chan struct{})

	for i := 0; i < MaxTryTimes; i++ {
		conn, _, err := m.connect()
		if err != nil {
			m.logger.Info(fmt.Sprintf("connect error, times(%v), error: %s", i, err.Error()))
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

	m.setIsConnected(true)

	m.resubscribe()

	if m.autoReconnect {
		go m.reconnect()
	}

	m.auth()

	go m.readMessages()

	return nil
}

func (m *AccountWsClient) connect() (*websocket.Conn, *http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, m.baseURL, nil)
	if err == nil {
		conn.SetReadLimit(32768 * 64)
	}

	return conn, resp, err
}

func (m *AccountWsClient) reconnect() {
	<-m.disconnect

	m.setIsConnected(false)

	m.logger.Info("disconnect, then reconnect...")

	time.Sleep(1 * time.Second)

	select {
	case <-m.ctx.Done():
		m.logger.Info(fmt.Sprintf("never reconnect, %s", m.ctx.Err()))
		return
	default:
		m.start()
	}
}

// close closes the websocket connection
func (m *AccountWsClient) close() error {
	close(m.disconnect)

	err := m.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

// setIsConnected sets state for isConnected
func (m *AccountWsClient) setIsConnected(state bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.isConnected = state
}

// IsConnected returns the WebSocket connection state
func (m *AccountWsClient) IsConnected() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.isConnected
}

func (m *AccountWsClient) auth() error {
	parameters := url.Values{}

	parameters.Add("AccessKeyId", m.key)
	parameters.Add("SignatureMethod", "HmacSHA256")
	parameters.Add("SignatureVersion", "2")

	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05")
	parameters.Add("Timestamp", timestamp)

	var sb strings.Builder
	sb.WriteString(http.MethodGet)
	sb.WriteString("\n")
	sb.WriteString(m.host)
	sb.WriteString("\n")
	sb.WriteString(m.path)
	sb.WriteString("\n")
	sb.WriteString(parameters.Encode())

	hm := hmac.New(sha256.New, []byte(m.secret))
	hm.Write([]byte(sb.String()))
	sign := base64.StdEncoding.EncodeToString(hm.Sum(nil))

	msg := AuthRequest{
		Operation:        "auth",
		Type:             "api",
		AccessKeyId:      m.key,
		SignatureMethod:  "HmacSHA256",
		SignatureVersion: "2",
		Timestamp:        timestamp,
		Signature:        sign,
	}

	m.sending.Lock()
	defer m.sending.Unlock()

	if !m.IsConnected() {
		return errors.New("connection is closed")
	}

	return m.conn.WriteJSON(msg)
}

func (m *AccountWsClient) readMessages() {
	for {
		select {
		case <-m.ctx.Done():
			m.logger.Info(fmt.Sprintf("context done, error: %s", m.ctx.Err().Error()))

			if err := m.close(); err != nil {
				m.logger.Info(fmt.Sprintf("websocket connection closed error, %s", err.Error()))
			}

			return
		default:
			msgType, buf, err := m.conn.ReadMessage()
			if err != nil {
				m.logger.Error(fmt.Sprintf("read message error, %s", err))
				time.Sleep(TimerIntervalSecond * time.Second)
				continue
			}

			// decompress gzip data if it is binary message
			if msgType == websocket.BinaryMessage {
				message, err := htxutils.GZipDecompress(buf)
				if err != nil {
					m.logger.Error(fmt.Sprintf("ungzip data error: %s", err))

					if err := m.close(); err != nil {
						m.logger.Error(fmt.Sprintf("websocket connection closed error, %s", err.Error()))
					}

					return
				}

				var msg Message

				err = json.Unmarshal([]byte(message), &msg)
				if err != nil {
					m.logger.Error(fmt.Sprintf("read object error, %s", err))

					if err := m.close(); err != nil {
						m.logger.Error(fmt.Sprintf("websocket connection closed error, %s", err.Error()))
					}

					return
				}

				switch {
				case msg.Operation == PING:
					err := m.pong(&Message{
						Operation: PONG,
						Ts:        msg.Ts,
					})
					if err != nil {
						m.logger.Error(fmt.Sprintf("handle ping error: %s", err.Error()))
					}
				case msg.Operation == SUB:
					if msg.ErrCode != 0 {
						m.logger.Error(fmt.Sprintf("sub websocket error, op: %s, topic: %s, err-code: %v, err-msg: %v", msg.Operation, msg.Topic, msg.ErrCode, msg.ErrMsg))
					}
				case msg.Operation == AUTH:
					if msg.ErrCode != 0 {
						m.logger.Error(fmt.Sprintf("auth websocket error, op: %s, err-code: %v, err-msg: %v", msg.Operation, msg.ErrCode, msg.ErrMsg))

						if err := m.close(); err != nil {
							m.logger.Error(fmt.Sprintf("websocket connection closed error, %s", err.Error()))
						}

						return
					} else {
						m.logger.Info(fmt.Sprintf("auth websocket success, op: %s, err-code: %v", msg.Operation, msg.ErrCode))
					}
				case msg.Operation == "notify":
					err := m.handle(&msg)
					if err != nil {
						m.logger.Error(fmt.Sprintf("handle message error: %s", err.Error()))
					}
				}
			}
		}
	}
}

func (m *AccountWsClient) resubscribe() error {
	topics := m.subscriptions.Keys()

	if len(topics) == 0 {
		return nil
	}

	redo := make([]string, 0)

	for _, v := range topics {
		// do subscription
		err := m.send(&Message{
			Operation: SUB,
			Topic:     v,
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

func (m *AccountWsClient) subscribe(topic string) error {
	if m.subscriptions.Has(topic) {
		return nil
	}

	// do subscription

	err := m.send(&Message{
		Operation: SUB,
		Topic:     topic,
	})

	if err != nil {
		return err
	}

	m.subscriptions.Set(topic, struct{}{})

	return nil
}

func (m *AccountWsClient) unsubscribe(topic string) error {
	err := m.send(&Message{
		Operation: SUB,
		Topic:     topic,
	})

	if err != nil {
		return err
	}

	m.subscriptions.Remove(topic)

	return nil
}

func (m *AccountWsClient) send(req *Message) error {
	m.sending.Lock()
	defer m.sending.Unlock()

	if !m.IsConnected() {
		return errors.New("connection is closed")
	}

	return m.conn.WriteJSON(req)
}

func (m *AccountWsClient) pong(ping *Message) error {
	m.sending.Lock()

	// Rate Limit: https://www.htx.com/en-us/opend/newApiPages/?id=662
	defer time.Sleep(100 * time.Millisecond)
	defer m.sending.Unlock()

	if !m.IsConnected() {
		return errors.New("connection is closed")
	}

	return m.conn.WriteJSON(ping)
}
