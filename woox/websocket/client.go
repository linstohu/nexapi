package websocket

import (
	"context"
	"errors"
	"fmt"
	"log"
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
	logger *log.Logger

	ctx         context.Context
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
	BaseURL       string `validate:"required"`
	Key           string
	Secret        string
	ApplicationID string `validate:"required"`
	Debug         bool
	// Logger
	Logger *log.Logger
}

func NewWooXWebsocketClient(ctx context.Context, cfg *WooXWebsocketCfg) (*WooXWebsocketClient, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, err
	}

	cli := &WooXWebsocketClient{
		baseURL:       cfg.BaseURL,
		key:           cfg.Key,
		secret:        cfg.Secret,
		applicationID: cfg.ApplicationID,
		debug:         cfg.Debug,
		logger:        cfg.Logger,

		ctx:           ctx,
		autoReconnect: true,

		subscriptions: cmap.New[struct{}](),
		emitter:       emission.NewEmitter(),
	}

	if cli.logger == nil {
		cli.logger = log.Default()
		cli.logger.SetPrefix("woox-websocket")
	}

	err := cli.start()
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func (w *WooXWebsocketClient) start() error {
	w.conn = nil
	w.setIsConnected(false)
	w.heartCancel = make(chan struct{})
	w.disconnect = make(chan struct{})

	for i := 0; i < MaxTryTimes; i++ {
		conn, _, err := w.connect()
		if err != nil {
			w.logger.Printf("connect error, times(%v), error: %s", i, err.Error())
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

	w.logger.Printf("disconnect, then reconnect...")

	close(w.heartCancel)

	time.Sleep(1 * time.Second)

	select {
	case <-w.ctx.Done():
		w.logger.Printf("never reconnect, %s", w.ctx.Err())
		return
	default:
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
		case <-w.ctx.Done():
			w.logger.Println(w.ctx.Err())

			if err := w.close(); err != nil {
				w.logger.Printf("websocket connection closed error, %s", err.Error())
			}

			return
		default:
			var msg types.AnyMessage
			err := w.conn.ReadJSON(&msg)
			if err != nil {
				w.logger.Printf("read object error, %s", err)

				if err := w.close(); err != nil {
					w.logger.Printf("websocket connection closed error, %s", err.Error())
				}

				return
			}

			switch {
			case msg.Response != nil:
				// todo
			case msg.SubscribedMessage != nil:
				err := w.handle(msg.SubscribedMessage)
				if err != nil {
					w.logger.Printf("read messages error: %s", err.Error())
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
