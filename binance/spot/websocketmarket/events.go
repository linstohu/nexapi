package websocketmarket

import "github.com/chuckpreslar/emission"

type Listener func(any)

func (m *SpotMarketStreamClient) AddListener(event string, listener Listener) *emission.Emitter {
	return m.emitter.On(event, listener)
}

func (m *SpotMarketStreamClient) RemoveListener(event string, listener Listener) *emission.Emitter {
	return m.emitter.Off(m, listener)
}

func (m *SpotMarketStreamClient) GetListeners(event string, argument any) *emission.Emitter {
	return m.emitter.Emit(event, argument)
}
