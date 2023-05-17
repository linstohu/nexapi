package types

import (
	"encoding/json"
	"errors"

	"github.com/valyala/fastjson"
)

type Request struct {
	ID    string `json:"id,omitempty"`
	Topic string `json:"topic,omitempty"`
	Event string `json:"event,omitempty"`
}

// AnyMessage represents either a JSON Response or SubscribedMessage.
type AnyMessage struct {
	Response          *Response
	SubscribedMessage *SubscribedMessage
}

type Response struct {
	ID        string `json:"id"`
	Event     string `json:"event"`
	Success   bool   `json:"success"`
	Timestamp int64  `json:"ts"`
}

type SubscribedMessage struct {
	OriginData []byte
	Topic      string          `json:"topic"`
	Timestamp  int64           `json:"ts"`
	Data       json.RawMessage `json:"data"`
}

func (m AnyMessage) MarshalJSON() ([]byte, error) {
	var v interface{}

	switch {
	case m.Response != nil && m.SubscribedMessage == nil:
		v = m.Response
	case m.Response == nil && m.SubscribedMessage != nil:
		v = m.SubscribedMessage
	}

	if v != nil {
		return json.Marshal(v)
	}

	return nil, errors.New("woox websocket: message must have exactly one of the Response or SubscribedMessage fields set")
}

func (m *AnyMessage) UnmarshalJSON(data []byte) error {
	var p fastjson.Parser
	v, err := p.ParseBytes(data)
	if err != nil {
		return err
	}

	if v.Exists("event") {
		var resp Response

		if err := json.Unmarshal(data, &resp); err != nil {
			return err
		}

		m.Response = &resp

		return nil
	}

	if v.Exists("topic") {
		des := make([]byte, len(data))
		copy(des, data)

		var msg = &SubscribedMessage{
			OriginData: des,
			Topic:      string(v.GetStringBytes("topic")),
			Timestamp:  v.GetInt64("ts"),
			Data:       v.GetStringBytes("data"),
		}

		m.SubscribedMessage = msg

		return nil
	}

	return nil
}
