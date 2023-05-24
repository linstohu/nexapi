package types

import (
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type GetTradeListParam struct {
	Symbol    string `url:"symbol,omitempty" validate:"omitempty"`
	FromID    int64  `url:"fromId,omitempty" validate:"omitempty"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type GetTradeListParams struct {
	GetTradeListParam
	bnutils.DefaultParam
}

type UserTrade struct {
	ID             int64  `json:"id"`
	TradeID        int64  `json:"tradeId"`
	OrderID        int64  `json:"orderId"`
	Symbol         string `json:"symbol"`
	Price          string `json:"price"`
	Quantity       string `json:"quantity"`
	Fee            string `json:"fee"`
	RealizedProfit string `json:"realizedProfit"`
	Side           string `json:"side"`
	Type           string `json:"type"`
	Volatility     string `json:"volatility"`
	Liquidity      string `json:"liquidity"`
	QuoteAsset     string `json:"quoteAsset"`
	Time           int64  `json:"time"`
	PriceScale     int    `json:"priceScale"`
	QuantityScale  int    `json:"quantityScale"`
	OptionSide     string `json:"optionSide"`
}
