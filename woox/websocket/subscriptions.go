package websocket

func (w *WooXStreamClient) Subscribe(channels []string) error {
	w.subscriptions = append(w.subscriptions, channels...)
	return w.subscribe(channels)
}

func (w *WooXStreamClient) UnSubscribe(channels []string) error {
	channelMap := make(map[string]struct{})
	for _, v := range channels {
		channelMap[v] = struct{}{}
	}

	subscriptions := make([]string, 0)

	for _, v := range w.subscriptions {
		if _, ok := channelMap[v]; !ok {
			subscriptions = append(subscriptions, v)
		}
	}

	w.subscriptions = subscriptions

	return w.unsubscribe(channels)
}

func (w *WooXStreamClient) subscriptionsProcess(event *SubscribedMessage) {
	if w.debug {
		w.logger.Printf("woox subscribed message, topic: %s, timestamp: %v, data: %s",
			event.Topic, event.Timestamp, event.Data)
	}
}
