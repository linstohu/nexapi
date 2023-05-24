package websocketmarket

import "github.com/chuckpreslar/emission"

type Listener func(any)

func (o *OptionsMarketStreamClient) AddListener(event string, listener Listener) *emission.Emitter {
	return o.emitter.On(event, listener)
}

func (o *OptionsMarketStreamClient) RemoveListener(event string, listener Listener) *emission.Emitter {
	return o.emitter.Off(o, listener)
}

func (o *OptionsMarketStreamClient) GetListeners(event string, argument any) *emission.Emitter {
	return o.emitter.Emit(event, argument)
}
