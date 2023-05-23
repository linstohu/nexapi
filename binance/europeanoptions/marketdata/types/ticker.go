package types

type GetTickerPriceParam struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
}

type TickerPrice struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	Open               string `json:"open"`
	High               string `json:"high"`
	Low                string `json:"low"`
	Volume             string `json:"volume"`
	Amount             string `json:"amount"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstTradeID       int64  `json:"firstTradeId"`
	TradeCount         int64  `json:"tradeCount"`
	StrikePrice        string `json:"strikePrice"`
	ExercisePrice      string `json:"exercisePrice"`
}
