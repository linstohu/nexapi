package websocketuserdata

import "github.com/chuckpreslar/emission"

type Listener func(any)

func (o *OptionsUserDataStreamClient) AddListener(event string, listener Listener) *emission.Emitter {
	return o.emitter.On(event, listener)
}

func (o *OptionsUserDataStreamClient) RemoveListener(event string, listener Listener) *emission.Emitter {
	return o.emitter.Off(o, listener)
}

func (o *OptionsUserDataStreamClient) GetListeners(event string, argument any) *emission.Emitter {
	return o.emitter.Emit(event, argument)
}
