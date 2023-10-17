package websocketmarket

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/linstohu/nexapi/binance/spot/websocketmarket/types"
	"github.com/linstohu/nexapi/binance/utils"
)

func (m *SpotMarketStreamClient) Subscribe(topics []string) error {
	return m.subscribe(topics)
}

func (m *SpotMarketStreamClient) UnSubscribe(topics []string) error {
	return m.unsubscribe(topics)
}

func (m *SpotMarketStreamClient) handle(msg *utils.SubscribedMessage) error {
	if m.debug {
		m.logger.Info(fmt.Sprintf("subscribed message, stream: %s", msg.Stream))
	}

	switch {
	case strings.HasSuffix(msg.Stream, "@aggTrade"):
		var data types.AggregateTrade
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		m.GetListeners(msg.Stream, &data)
	case strings.HasSuffix(msg.Stream, "@trade"):
		var data types.Trade
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		m.GetListeners(msg.Stream, &data)
	case strings.Contains(msg.Stream, "@kline_"):
		var data types.Kline
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		m.GetListeners(msg.Stream, &data)
	case strings.HasSuffix(msg.Stream, "@miniTicker"):
		var data types.MiniTicker
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		m.GetListeners(msg.Stream, &data)
	case msg.Stream == "!miniTicker@arr":
		var data []*types.MiniTicker
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		m.GetListeners(msg.Stream, data)
	case strings.HasSuffix(msg.Stream, "@ticker"):
		var data types.Ticker
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		m.GetListeners(msg.Stream, &data)
	case msg.Stream == "!ticker@arr":
		var data []*types.Ticker
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		m.GetListeners(msg.Stream, data)
	case strings.HasSuffix(msg.Stream, "@bookTicker"):
		var data types.BookTicker
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}
		m.GetListeners(msg.Stream, &data)
	case strings.Contains(msg.Stream, "@depth"):
		switch {
		case strings.Contains(msg.Stream, "@depth5") ||
			strings.Contains(msg.Stream, "@depth10") ||
			strings.Contains(msg.Stream, "@depth20"):
			var data types.OrderbookDepth
			err := json.Unmarshal(msg.Data, &data)
			if err != nil {
				return err
			}
			m.GetListeners(msg.Stream, &data)
		default:
			var data types.OrderbookDiffDepth
			err := json.Unmarshal(msg.Data, &data)
			if err != nil {
				return err
			}
			m.GetListeners(msg.Stream, &data)
		}
	default:
		return fmt.Errorf("unknown message, topic: %s", msg.Stream)
	}

	return nil
}
