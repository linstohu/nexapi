package types

import "github.com/linstohu/nexapi/binance/usdmfutures/utils"

type AggregateTrade struct {
	EventType        string `json:"e"`
	EventTime        int64  `json:"E"`
	Symbol           string `json:"s"`
	AggregateTradeID int64  `json:"a"`
	Price            string `json:"p"`
	Quantity         string `json:"q"`
	FirstTradeID     int64  `json:"f"`
	LastTradeID      int64  `json:"l"`
	TradeTime        int64  `json:"T"`
	IsBuyerMaker     bool   `json:"m"`
}

type MarkPrice struct {
	EventType            string `json:"e"`
	EventTime            int64  `json:"E"`
	Symbol               string `json:"s"`
	MarkPrice            string `json:"p"`
	IndexPrice           string `json:"i"`
	EstimatedSettlePrice string `json:"P"`
	FundingRate          string `json:"r"`
	NextFundingTime      int64  `json:"T"`
}

type Ticker struct {
	EventType                   string `json:"e"`
	EventTime                   int64  `json:"E"`
	Symbol                      string `json:"s"`
	PriceChange                 string `json:"p"`
	PriceChangePercent          string `json:"P"`
	WeightedAveragePrice        string `json:"w"`
	LastPrice                   string `json:"c"`
	LastQuantity                string `json:"Q"`
	OpenPrice                   string `json:"o"`
	HighPrice                   string `json:"h"`
	LowPrice                    string `json:"l"`
	TotalTradedBaseAssetVolume  string `json:"v"`
	TotalTradedQuoteAssetVolume string `json:"q"`
	StatisticsOpenTime          int64  `json:"O"`
	StatisticsCloseTime         int64  `json:"C"`
	FirstTradeID                int64  `json:"F"`
	LastTradeID                 int64  `json:"L"`
	TradesNum                   int64  `json:"n"`
}

type BookTicker struct {
	EventType       string `json:"e"`
	UpdateID        int64  `json:"u"`
	EventTime       int64  `json:"E"`
	TransactionTime int64  `json:"T"`
	Symbol          string `json:"s"`
	BestBidPrice    string `json:"b"`
	BestBidQty      string `json:"B"`
	BestAskPrice    string `json:"a"`
	BestAskQty      string `json:"A"`
}

type LiquidationOrder struct {
	EventType string `json:"e"`
	EventTime int64  `json:"E"`
	Order     struct {
		Symbol                    string            `json:"s"`
		Side                      string            `json:"S"`
		OrderType                 utils.OrderType   `json:"o"`
		TimeInForce               utils.TimeInForce `json:"f"`
		OriginalQuantity          string            `json:"q"`
		Price                     string            `json:"p"`
		AveragePrice              string            `json:"ap"`
		OrderStatus               utils.OrderStatus `json:"X"`
		LastFilledQuantity        string            `json:"l"`
		FilledAccumulatedQuantity string            `json:"z"`
		TradeTime                 int64             `json:"T"`
	} `json:"o"`
}

type OrderbookDepth struct {
	EventType               string     `json:"e"`
	EventTime               int64      `json:"E"`
	TransactionTime         int64      `json:"T"`
	Symbol                  string     `json:"s"`
	FirstUpdateID           int64      `json:"U"`
	FinalUpdateID           int64      `json:"u"`
	FinalUpdateIDLastStream int64      `json:"pu"`
	Bids                    [][]string `json:"b"`
	Asks                    [][]string `json:"a"`
}
