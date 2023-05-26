package types

import (
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type GetPositionParam struct {
	MarginAsset string `url:"marginAsset,omitempty" validate:"omitempty"`
	Pair        string `url:"pair,omitempty" validate:"omitempty"`
}

type GetPositionParams struct {
	GetPositionParam
	bnutils.DefaultParam
}

type Position struct {
	Symbol           string `json:"symbol"`
	PositionAmt      string `json:"positionAmt"`
	EntryPrice       string `json:"entryPrice"`
	MarkPrice        string `json:"markPrice"`
	UnRealizedProfit string `json:"unRealizedProfit"`
	LiquidationPrice string `json:"liquidationPrice"`
	Leverage         string `json:"leverage"`
	MaxQty           string `json:"maxQty"`
	MarginType       string `json:"marginType"`
	IsolatedMargin   string `json:"isolatedMargin"`
	IsAutoAddMargin  string `json:"isAutoAddMargin"`
	PositionSide     string `json:"positionSide"`
	UpdateTime       int64  `json:"updateTime"`
}
