package websocketmarket

import "github.com/chuckpreslar/emission"

type Listener func(any)

func (u *CoinMarginedMarketStreamClient) AddListener(event string, listener Listener) *emission.Emitter {
	return u.emitter.On(event, listener)
}

func (u *CoinMarginedMarketStreamClient) RemoveListener(event string, listener Listener) *emission.Emitter {
	return u.emitter.Off(u, listener)
}

func (u *CoinMarginedMarketStreamClient) GetListeners(event string, argument any) *emission.Emitter {
	return u.emitter.Emit(event, argument)
}
