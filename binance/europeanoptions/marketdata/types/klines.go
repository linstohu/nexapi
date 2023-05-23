package types

type Kline struct {
	Open        string `json:"open"`
	High        string `json:"high"`
	Low         string `json:"low"`
	Close       string `json:"close"`
	Volume      string `json:"volume"`
	Interval    string `json:"interval"`
	TradeCount  int    `json:"tradeCount"`
	TakerVolume string `json:"takerVolume"`
	TakerAmount string `json:"takerAmount"`
	Amount      string `json:"amount"`
	OpenTime    int64  `json:"openTime"`
	CloseTime   int64  `json:"closeTime"`
}