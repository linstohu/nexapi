package types

import (
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type GetPositionInfoParam struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
}

type GetPositionInfoParams struct {
	GetPositionInfoParam
	bnutils.DefaultParam
}

type Position struct {
	EntryPrice    string `json:"entryPrice"`
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Quantity      string `json:"quantity"`
	ReducibleQty  string `json:"reducibleQty"`
	MarkValue     string `json:"markValue"`
	Ror           string `json:"ror"`
	UnrealizedPNL string `json:"unrealizedPNL"`
	MarkPrice     string `json:"markPrice"`
	StrikePrice   string `json:"strikePrice"`
	PositionCost  string `json:"positionCost"`
	ExpiryDate    int64  `json:"expiryDate"`
	PriceScale    int    `json:"priceScale"`
	QuantityScale int    `json:"quantityScale"`
	OptionSide    string `json:"optionSide"`
	QuoteAsset    string `json:"quoteAsset"`
}
