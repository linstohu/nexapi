package websocketmarket

import "github.com/chuckpreslar/emission"

type Listener func(any)

func (m *MarketStreamClient) AddListener(event string, listener Listener) *emission.Emitter {
	return m.emitter.On(event, listener)
}

func (m *MarketStreamClient) RemoveListener(event string, listener Listener) *emission.Emitter {
	return m.emitter.Off(m, listener)
}

func (m *MarketStreamClient) GetListeners(event string, argument any) *emission.Emitter {
	return m.emitter.Emit(event, argument)
}
