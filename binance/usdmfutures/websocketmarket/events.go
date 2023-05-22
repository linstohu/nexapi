package websocketmarket

import "github.com/chuckpreslar/emission"

type Listener func(any)

func (u *USDMarginedMarketStreamClient) AddListener(event string, listener Listener) *emission.Emitter {
	return u.emitter.On(event, listener)
}

func (u *USDMarginedMarketStreamClient) RemoveListener(event string, listener Listener) *emission.Emitter {
	return u.emitter.Off(u, listener)
}

func (u *USDMarginedMarketStreamClient) GetListeners(event string, argument any) *emission.Emitter {
	return u.emitter.Emit(event, argument)
}
