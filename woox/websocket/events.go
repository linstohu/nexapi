package websocket

import "github.com/chuckpreslar/emission"

type Listener func(interface{})

func (w *WooXWebsocketClient) AddListener(event string, listener Listener) *emission.Emitter {
	return w.emitter.On(event, listener)
}

func (w *WooXWebsocketClient) RemoveListener(event string, listener Listener) *emission.Emitter {
	return w.emitter.Off(w, listener)
}

func (w *WooXWebsocketClient) GetListeners(event string, argument interface{}) *emission.Emitter {
	return w.emitter.Emit(event, argument)
}
