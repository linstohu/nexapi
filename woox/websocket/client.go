package websocket

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/chuckpreslar/emission"
	"github.com/go-playground/validator"
	cmap "github.com/orcaman/concurrent-map/v2"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type WooXStreamClient struct {
	addr                       string
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

	sending sync.Mutex

	subscriptions    []string
	subscriptionsMap cmap.ConcurrentMap[string, struct{}]

	emitter *emission.Emitter
}

type WooXStreamCfg struct {
	BasePath      string `validate:"required"`
	Key           string
	Secret        string
	ApplicationID string `validate:"required"`
	Debug         bool
	// Logger
	Logger *log.Logger
}

func NewWooXClient(cfg *WooXStreamCfg) (*WooXStreamClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := &WooXStreamClient{
		addr:          cfg.BasePath,
		key:           cfg.Key,
		secret:        cfg.Secret,
		applicationID: cfg.ApplicationID,
		debug:         cfg.Debug,
		logger:        cfg.Logger,

		ctx:           context.Background(),
		autoReconnect: true,

		subscriptions: make([]string, 0),
		emitter:       emission.NewEmitter(),
	}

	if cli.logger == nil {
		cli.logger = log.Default()
	}

	err = cli.start()
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func (w *WooXStreamClient) start() error {
	w.conn = nil
	w.setIsConnected(false)
	w.heartCancel = make(chan struct{})
	w.disconnect = make(chan struct{})
	w.subscriptionsMap = cmap.New[struct{}]()

	for i := 0; i < MaxTryTimes; i++ {
		conn, _, err := w.connect()
		if err != nil {
			log.Printf("woox connect error, times(%v), error: %s", i, err.Error())
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

func (w *WooXStreamClient) connect() (*websocket.Conn, *http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, resp, err := websocket.Dial(ctx, w.addr+w.applicationID, &websocket.DialOptions{})
	if err == nil {
		conn.SetReadLimit(32768 * 64)
	}

	return conn, resp, err
}

func (w *WooXStreamClient) reconnect() {
	<-w.disconnect

	w.setIsConnected(false)

	w.logger.Printf("woox disconnect, then reconnect...")

	close(w.heartCancel)

	time.Sleep(1 * time.Second)

	w.start()
}

// close closes the websocket connection
func (w *WooXStreamClient) close(cause error) error {
	close(w.disconnect)

	err := w.conn.Close(websocket.StatusNormalClosure, cause.Error())
	if err != nil {
		return err
	}

	w.logger.Printf("woox websocket connection closed, reason: %s", cause.Error())

	return nil
}

// setIsConnected sets state for isConnected
func (w *WooXStreamClient) setIsConnected(state bool) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.isConnected = state
}

// IsConnected returns the WebSocket connection state
func (w *WooXStreamClient) IsConnected() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.isConnected
}

// heartbeat sends ping every 5s to keep alive
func (w *WooXStreamClient) heartbeat() {
	t := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t.C:
			w.send(&Request{
				Event: "ping",
			})
		case <-w.heartCancel:
			return
		}
	}
}

func (w *WooXStreamClient) readMessages() {
	var err error

	for err == nil {
		var m AnyMessage
		err = w.readObject(&m)
		if err != nil {
			break
		}

		switch {
		case m.Response != nil:
		case m.SubscribedMessage != nil:
			w.subscriptionsProcess(m.SubscribedMessage)
		}
	}

	w.close(err)
}

func (w *WooXStreamClient) readObject(v interface{}) error {
	err := wsjson.Read(context.Background(), w.conn, v)
	if e, ok := err.(*websocket.CloseError); ok {
		if e.Code == websocket.StatusNormalClosure && e.Error() == io.ErrUnexpectedEOF.Error() {
			// unwrapping this error.
			err = io.ErrUnexpectedEOF
		}
	}
	return err
}

func (w *WooXStreamClient) resubscribe() error {
	var publicChannels []string

	for _, v := range w.subscriptions {
		if ok := w.subscriptionsMap.Has(v); ok {
			continue
		}

		publicChannels = append(publicChannels, v)
	}

	// do subscription

	for _, v := range publicChannels {
		err := w.send(&Request{
			ID:    genClientID(),
			Topic: v,
			Event: "subscribe",
		})

		if err != nil {
			return err
		}

		w.subscriptionsMap.Set(v, struct{}{})
	}

	return nil
}

func (w *WooXStreamClient) subscribe(channels []string) error {
	var publicChannels []string

	for _, v := range channels {
		if ok := w.subscriptionsMap.Has(v); ok {
			continue
		}

		publicChannels = append(publicChannels, v)
	}

	// do subscription

	for _, v := range publicChannels {
		err := w.send(&Request{
			ID:    genClientID(),
			Topic: v,
			Event: SUBSCRIBE,
		})

		if err != nil {
			return err
		}

		w.subscriptionsMap.Set(v, struct{}{})
	}

	return nil
}

func (w *WooXStreamClient) unsubscribe(channels []string) error {
	for _, v := range channels {
		err := w.send(&Request{
			ID:    genClientID(),
			Topic: v,
			Event: UNSUBSCRIBE,
		})

		if err != nil {
			return err
		}

		w.subscriptionsMap.Remove(v)
	}

	return nil
}

func (w *WooXStreamClient) send(req *Request) error {
	w.sending.Lock()
	defer w.sending.Unlock()

	if !w.IsConnected() {
		return errors.New("woox: connection is closed")
	}

	return wsjson.Write(context.Background(), w.conn, req)
}

func genClientID() string {
	return fmt.Sprintf("ClientID%d", rand.Intn(100))
}
