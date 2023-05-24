package websocketuserdata

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

func (o *OptionsUserDataStreamClient) handle(origind []byte) error {
	var p fastjson.Parser
	pb, err := p.ParseBytes(origind)
	if err != nil {
		return err
	}

	eventType := string(pb.GetStringBytes("e"))

	if o.debug {
		o.logger.Printf("subscribed message, event-type: %s", eventType)
	}

	switch eventType {
	case o.GenAccountDataTopic():
		var data AccountData
		err := json.Unmarshal(origind, &data)
		if err != nil {
			return err
		}
		o.GetListeners(o.GenAccountDataTopic(), &data)
	case o.GenOrderUpdateTopic():
		var data OrderUpdate
		err := json.Unmarshal(origind, &data)
		if err != nil {
			return err
		}
		o.GetListeners(o.GenOrderUpdateTopic(), &data)
	default:
		return fmt.Errorf("unknown message, event-type: %s", eventType)
	}

	return nil
}

func (o *OptionsUserDataStreamClient) GenAccountDataTopic() string {
	return "ACCOUNT_UPDATE"
}

func (o *OptionsUserDataStreamClient) GenOrderUpdateTopic() string {
	return "ORDER_TRADE_UPDATE"
}
