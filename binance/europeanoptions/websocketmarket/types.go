package websocketmarket

type Trade struct {
	EventType          string `json:"e"`
	EventTime          int64  `json:"E"`
	Symbol             string `json:"s"`
	TradeID            string `json:"t"`
	Price              string `json:"p"`
	Quantity           string `json:"q"`
	BuyOrderID         string `json:"b"`
	SellOrderID        string `json:"a"`
	TradeCompletedTime int64  `json:"T"`
	Direction          string `json:"S"`
}

type IndexPrice struct {
	EventType  string `json:"e"`
	EventTime  int64  `json:"E"`
	Symbol     string `json:"s"`
	IndexPrice string `json:"p"`
}

type MarkPrice struct {
	EventType string `json:"e"`
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`
	MarkPrice string `json:"mp"`
}

type Kline struct {
	EventType string `json:"e"`
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`
	Kline     struct {
		StartTime             int64  `json:"t"`
		CloseTime             int64  `json:"T"`
		Symbol                string `json:"s"`
		Interval              string `json:"i"`
		FirstTradeID          int64  `json:"F"`
		LastTradeID           int64  `json:"L"`
		OpenPrice             string `json:"o"`
		ClosePrice            string `json:"c"`
		HighPrice             string `json:"h"`
		LowPrice              string `json:"l"`
		BaseAssetVolume       string `json:"v"`
		TradesNum             int64  `json:"n"`
		IsClosed              bool   `json:"x"`
		QuoteAssetTradeAmount string `json:"q"`
		TakerTradeVolume      string `json:"V"`
		TakerTradeAmount      string `json:"Q"`
	} `json:"k"`
}

type Ticker struct {
	EventType            string `json:"e"`
	EventTime            int64  `json:"E"`
	Symbol               string `json:"s"`
	OpeningPrice         string `json:"o"`
	HighPrice            string `json:"h"`
	LowPrice             string `json:"l"`
	LatestPrice          string `json:"c"`
	TradingVolume        string `json:"V"`
	TradingAmount        string `json:"A"`
	PriceChangePercent   string `json:"P"`
	PriceChange          string `json:"p"`
	LastTradeVolume      string `json:"Q"`
	FirstTradeID         string `json:"F"`
	LastTradeID          string `json:"L"`
	NumberOfTrades       int64  `json:"n"`
	BestBuyPrice         string `json:"bo"`
	BestSellPrice        string `json:"ao"`
	BestBuyQuantity      string `json:"bq"`
	BestSellQuantity     string `json:"aq"`
	BuyImplied           string `json:"b"`
	SellImplied          string `json:"a"`
	Delta                string `json:"d"`
	Theta                string `json:"t"`
	Gamma                string `json:"g"`
	Vega                 string `json:"v"`
	ImpliedVolatility    string `json:"vo"`
	MarkPrice            string `json:"mp"`
	BuyMaximumPrice      string `json:"hl"`
	SellMinimumPrice     string `json:"ll"`
	EstimatedStrikePrice string `json:"eep"`
}

type OpenInterest struct {
	EventType               string `json:"e"`
	EventTime               int64  `json:"E"`
	Symbol                  string `json:"s"`
	OpenInterestInContracts string `json:"o"`
	OpenInterestInUSDT      string `json:"h"`
}

type OrderbookDepth struct {
	EventType            string     `json:"e"`
	TransactionEventTime int64      `json:"T"`
	Symbol               string     `json:"s"`
	UpdateID             int64      `json:"u"`
	PU                   int64      `json:"pu"`
	Bids                 [][]string `json:"b"`
	Asks                 [][]string `json:"a"`
}
