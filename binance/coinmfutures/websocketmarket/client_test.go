package websocketmarket

import (
	"context"
	"fmt"
	"testing"

	spottypes "github.com/linstohu/nexapi/binance/spot/websocketmarket/types"
	usdmtypes "github.com/linstohu/nexapi/binance/usdmfutures/websocketmarket/types"
	"github.com/stretchr/testify/assert"
)

func testNewMarketStreamClient(ctx context.Context, t *testing.T) *CoinMarginedMarketStreamClient {
	cli, err := NewMarketStreamClient(ctx, &CoinMarginedMarketStreamCfg{
		BaseURL: CoinMarginedMarketStreamBaseURL,
		Debug:   true,
	})

	if err != nil {
		t.Fatalf("Could not create websocket client for woox, %s", err)
	}

	return cli
}

func TestSubscribeAggTrade(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetAggTradeTopic("btcusd_perp")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		trade, ok := e.(*usdmtypes.AggregateTrade)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, Price: %v, Quantity: %v, Time: %v\n",
			topic, trade.Symbol, trade.Price, trade.Quantity, trade.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeIndexPrice(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetIndexPriceTopic(&IndexPriceTopicParam{
		Pair:        "btcusd",
		UpdateSpeed: "1s",
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		indexPrice, ok := e.(*IndexPrice)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Pair: %v, IndexPrice: %v, Time: %v\n",
			topic, indexPrice.Pair, indexPrice.IndexPrice, indexPrice.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeMarkPrice(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetMarketPriceTopic(&MarkPriceTopicParam{
		Symbol:      "btcusd_perp",
		UpdateSpeed: "1s",
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		markprice, ok := e.(*MarkPrice)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, MarkPrice: %v, SettlePrice: %v, Time: %v\n",
			topic, markprice.Symbol, markprice.MarkPrice, markprice.EstimatedSettlePrice, markprice.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribePairMarkPrice(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetPairMarketPriceTopic(&PairMarkPriceTopicParam{
		Pair:        "btcusd",
		UpdateSpeed: "1s",
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		markprices, ok := e.([]*MarkPrice)
		if !ok {
			return
		}

		for _, markprice := range markprices {
			fmt.Printf("Topic: %s, Symbol: %v, MarkPrice: %v, SettlePrice: %v, Time: %v\n",
				topic, markprice.Symbol, markprice.MarkPrice, markprice.EstimatedSettlePrice, markprice.EventTime)
		}

	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeKline(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetKlineTopic(&KlineTopicParam{
		Symbol:   "btcusd_perp",
		Interval: "1m",
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		kline, ok := e.(*spottypes.Kline)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, OpenPrice: %v, ClosePrice: %v, Time: %v\n",
			topic, kline.Symbol, kline.Kline.OpenPrice, kline.Kline.ClosePrice, kline.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeMiniTicker(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetMiniTickerTopic("btcusd_perp")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		ticker, ok := e.(*MiniTicker)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, OpenPrice: %v, ClosePrice: %v, Time: %v\n",
			topic, ticker.Symbol, ticker.OpenPrice, ticker.ClosePrice, ticker.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeAllMiniTicker(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetAllMarketMiniTickersTopic()
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		tickers, ok := e.([]*MiniTicker)
		if !ok {
			return
		}

		for _, ticker := range tickers {
			fmt.Printf("Topic: %s, Symbol: %v, OpenPrice: %v, ClosePrice: %v, Time: %v\n",
				topic, ticker.Symbol, ticker.OpenPrice, ticker.ClosePrice, ticker.EventTime)
		}
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeTicker(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetTickerTopic("btcusd_perp")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		ticker, ok := e.(*Ticker)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, OpenPrice: %v, LastPrice: %v, Time: %v\n",
			topic, ticker.Symbol, ticker.OpenPrice, ticker.LastPrice, ticker.EventTime)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeAllTicker(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetAllMarketTickersTopic()
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		tickers, ok := e.([]*Ticker)
		if !ok {
			return
		}

		for _, ticker := range tickers {
			fmt.Printf("Topic: %s, Symbol: %v, OpenPrice: %v, LastPrice: %v, Time: %v\n",
				topic, ticker.Symbol, ticker.OpenPrice, ticker.LastPrice, ticker.EventTime)
		}
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeBookTicker(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetBookTickerTopic("btcusd_perp")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		book, ok := e.(*BookTicker)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, BestBidPrice: %v, BestBidQty: %v, BestAskPrice: %v, BestAskQty: %v\n",
			topic, book.Symbol, book.BestBidPrice, book.BestBidQty, book.BestAskPrice, book.BestAskQty)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeAllBookTickers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetAllBookTickersTopic()
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		book, ok := e.(*BookTicker)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, BestBidPrice: %v, BestBidQty: %v, BestAskPrice: %v, BestAskQty: %v\n",
			topic, book.Symbol, book.BestBidPrice, book.BestBidQty, book.BestAskPrice, book.BestAskQty)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeLiquidationOrder(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetLiquidationOrderTopic("btcusd_perp")
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		order, ok := e.(*LiquidationOrder)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, Side: %v, Price: %v, Quanty: %v\n",
			topic, order.Order.Symbol, order.Order.Side, order.Order.Price, order.Order.FilledAccumulatedQuantity)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeAllLiquidationOrders(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetAllLiquidationOrdersTopic()
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		order, ok := e.(*LiquidationOrder)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, Side: %v, Price: %v, Quanty: %v\n",
			topic, order.Order.Symbol, order.Order.Side, order.Order.Price, order.Order.FilledAccumulatedQuantity)
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeBookDepth(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetBookDepthTopic(&BookDepthTopicParam{
		Symbol:      "btcusd_perp",
		Level:       5,
		UpdateSpeed: "500ms",
	})
	assert.Nil(t, err)

	cli.AddListener(topic, func(e any) {
		book, ok := e.(*OrderbookDepth)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, FirstID: %v, FinalID: %v, Bids-Num: %v, Asks-Num: %v\n",
			topic, book.Symbol, book.FirstUpdateID, book.FinalUpdateID, len(book.Bids), len(book.Asks))
	})

	cli.Subscribe([]string{topic})

	select {}
}

func TestSubscribeBookDiffDepth(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cli := testNewMarketStreamClient(ctx, t)

	topic, err := cli.GetBookDiffDepthTopic(&BookDiffDepthTopicParam{
		Symbol:      "btcusd_perp",
		UpdateSpeed: "500ms",
	})
	assert.Nil(t, err)

	fmt.Println(topic)

	cli.AddListener(topic, func(e any) {
		book, ok := e.(*OrderbookDepth)
		if !ok {
			return
		}

		fmt.Printf("Topic: %s, Symbol: %v, FirstID: %v, FinalID: %v, Bids-Num: %v, Asks-Num: %v\n",
			topic, book.Symbol, book.FirstUpdateID, book.FinalUpdateID, len(book.Bids), len(book.Asks))
	})

	cli.Subscribe([]string{topic})

	select {}
}
