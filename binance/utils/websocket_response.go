package utils


import (
	"encoding/json"
	"errors"

	"github.com/valyala/fastjson"
)

type Request struct {
	ID     uint32   `json:"id,omitempty"`
	Method string   `json:"method,omitempty"`
	Params []string `json:"params,omitempty"`
}

// AnyMessage represents either a JSON Response or SubscribedMessage.
type AnyMessage struct {
	Response          *Response
	SubscribedMessage *SubscribedMessage
}

type Response struct {
	ID     uint `json:"id"`
	Result any  `json:"result"`
}

type SubscribedMessage struct {
	Stream string          `json:"stream"`
	Data   json.RawMessage `json:"data"`
}

func (m AnyMessage) MarshalJSON() ([]byte, error) {
	var v any

	switch {
	case m.Response != nil && m.SubscribedMessage == nil:
		v = m.Response
	case m.Response == nil && m.SubscribedMessage != nil:
		v = m.SubscribedMessage
	}

	if v != nil {
		return json.Marshal(v)
	}

	return nil, errors.New("message must have exactly one of the Response or SubscribedMessage fields set")
}

func (m *AnyMessage) UnmarshalJSON(data []byte) error {
	var p fastjson.Parser
	v, err := p.ParseBytes(data)
	if err != nil {
		return err
	}

	if v.Exists("id") {
		var resp Response

		if err := json.Unmarshal(data, &resp); err != nil {
			return err
		}

		m.Response = &resp

		return nil
	}

	if v.Exists("stream") {
		var msg = &SubscribedMessage{
			Stream: string(v.GetStringBytes("stream")),
			Data:   v.Get("data").MarshalTo(nil),
		}

		m.SubscribedMessage = msg

		return nil
	}

	return nil
}
