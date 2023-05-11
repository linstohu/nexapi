package websocket

import "github.com/chuckpreslar/emission"

type Listener func(interface{})

func (w *WooXStreamClient) AddListener(event string, listener Listener) *emission.Emitter {
	return w.emitter.On(event, listener)
}

func (w *WooXStreamClient) RemoveListener(event string, listener Listener) *emission.Emitter {
	return w.emitter.Off(w, listener)
}

func (w *WooXStreamClient) GetListeners(event string, argument interface{}) *emission.Emitter {
	return w.emitter.Emit(event, argument)
}
