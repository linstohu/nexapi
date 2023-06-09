package websocketmarket

import (
	"context"
	"errors"
	"log"
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
	logger *log.Logger

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

type SpotMarketStreamCfg struct {
	BaseURL string `validate:"required"`
	Debug   bool
	// Logger
	Logger *log.Logger
}

func NewSpotMarketStreamClient(ctx context.Context, cfg *SpotMarketStreamCfg) (*SpotMarketStreamClient, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, err
	}

	cli := &SpotMarketStreamClient{
		baseURL: cfg.BaseURL,
		debug:   cfg.Debug,
		logger:  cfg.Logger,

		ctx:           ctx,
		autoReconnect: true,

		subscriptions: cmap.New[struct{}](),
		emitter:       emission.NewEmitter(),
	}

	if cli.logger == nil {
		cli.logger = log.Default()
		cli.logger.SetPrefix("binance-spot-market-streams")
	}

	err := cli.start()
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func (m *SpotMarketStreamClient) start() error {
	m.conn = nil
	m.setIsConnected(false)
	m.disconnect = make(chan struct{})

	for i := 0; i < MaxTryTimes; i++ {
		conn, _, err := m.connect()
		if err != nil {
			m.logger.Printf("connect error, times(%v), error: %s", i, err.Error())
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

	m.logger.Println("disconnect, then reconnect...")

	time.Sleep(1 * time.Second)

	select {
	case <-m.ctx.Done():
		m.logger.Printf("never reconnect, %s", m.ctx.Err())
		return
	default:
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
		case <-m.ctx.Done():
			m.logger.Println(m.ctx.Err())

			if err := m.close(); err != nil {
				m.logger.Printf("websocket connection closed error, %s", err.Error())
			}

			return
		default:
			var msg utils.AnyMessage
			err := m.conn.ReadJSON(&msg)
			if err != nil {
				m.logger.Printf("read object error, %s", err)

				if err := m.close(); err != nil {
					m.logger.Printf("websocket connection closed error, %s", err.Error())
				}

				return
			}

			switch {
			case msg.Response != nil:
				// todo
			case msg.SubscribedMessage != nil:
				err := m.handle(msg.SubscribedMessage)
				if err != nil {
					m.logger.Printf("handle message error: %s", err.Error())
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
