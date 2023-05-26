package websocketmarket

import (
	cmutils "github.com/linstohu/nexapi/binance/coinmfutures/utils"
	usdmutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
)

type IndexPrice struct {
	EventType  string `json:"e"`
	EventTime  int64  `json:"E"`
	Pair       string `json:"i"`
	IndexPrice string `json:"p"`
}

type MarkPrice struct {
	EventType            string `json:"e"`
	EventTime            int64  `json:"E"`
	Symbol               string `json:"s"`
	MarkPrice            string `json:"p"`
	EstimatedSettlePrice string `json:"P"`
	FundingRate          string `json:"r"`
	NextFundingTime      int64  `json:"T"`
}

type MiniTicker struct {
	EventType                  string `json:"e"`
	EventTime                  int64  `json:"E"`
	Symbol                     string `json:"s"`
	Pair                       string `json:"ps"`
	ClosePrice                 string `json:"c"`
	OpenPrice                  string `json:"o"`
	HighPrice                  string `json:"h"`
	LowPrice                   string `json:"l"`
	TotalTradedVolume          string `json:"v"`
	TotalTradedBaseAssetVolume string `json:"q"`
}

type Ticker struct {
	EventType                  string `json:"e"`
	EventTime                  int64  `json:"E"`
	Symbol                     string `json:"s"`
	Pair                       string `json:"ps"`
	PriceChange                string `json:"p"`
	PriceChangePercent         string `json:"P"`
	WeightedAveragePrice       string `json:"w"`
	LastPrice                  string `json:"c"`
	LastQuantity               string `json:"Q"`
	OpenPrice                  string `json:"o"`
	HighPrice                  string `json:"h"`
	LowPrice                   string `json:"l"`
	TotalTradedVolume          string `json:"v"`
	TotalTradedBaseAssetVolume string `json:"q"`
	StatisticsOpenTime         int64  `json:"O"`
	StatisticsCloseTime        int64  `json:"C"`
	FirstTradeID               int64  `json:"F"`
	LastTradeID                int64  `json:"L"`
	TradesNum                  int64  `json:"n"`
}

type BookTicker struct {
	EventType       string `json:"e"`
	UpdateID        int64  `json:"u"`
	Symbol          string `json:"s"`
	Pair            string `json:"ps"`
	BestBidPrice    string `json:"b"`
	BestBidQty      string `json:"B"`
	BestAskPrice    string `json:"a"`
	BestAskQty      string `json:"A"`
	TransactionTime int64  `json:"T"`
	EventTime       int64  `json:"E"`
}

type LiquidationOrder struct {
	EventType string `json:"e"`
	EventTime int64  `json:"E"`
	Order     struct {
		Symbol                    string                `json:"s"`
		Pair                      string                `json:"ps"`
		Side                      string                `json:"S"`
		OrderType                 usdmutils.OrderType   `json:"o"`
		TimeInForce               usdmutils.TimeInForce `json:"f"`
		OriginalQuantity          string                `json:"q"`
		Price                     string                `json:"p"`
		AveragePrice              string                `json:"ap"`
		OrderStatus               cmutils.OrderStatus     `json:"X"`
		LastFilledQuantity        string                `json:"l"`
		FilledAccumulatedQuantity string                `json:"z"`
		TradeTime                 int64                 `json:"T"`
	} `json:"o"`
}

type OrderbookDepth struct {
	EventType               string     `json:"e"`
	EventTime               int64      `json:"E"`
	TransactionTime         int64      `json:"T"`
	Symbol                  string     `json:"s"`
	Pair                    string     `json:"ps"`
	FirstUpdateID           int64      `json:"U"`
	FinalUpdateID           int64      `json:"u"`
	FinalUpdateIDLastStream int64      `json:"pu"`
	Bids                    [][]string `json:"b"`
	Asks                    [][]string `json:"a"`
}
