package types

import (
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type GetTradeListParam struct {
	Symbol    string `url:"symbol" validate:"required"`
	OrderID   int64  `url:"orderId,omitempty" validate:"omitempty"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	FromID    int64  `url:"fromId,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type GetTradeListParams struct {
	GetTradeListParam
	bnutils.DefaultParam
}

type Trade struct {
	Symbol          string `json:"symbol"`
	ID              int64  `json:"id"`
	OrderID         int64  `json:"orderId"`
	Pair            string `json:"pair"`
	Side            string `json:"side"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	RealizedPnl     string `json:"realizedPnl"`
	MarginAsset     string `json:"marginAsset"`
	BaseQty         string `json:"baseQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	PositionSide    string `json:"positionSide"`
	Buyer           bool   `json:"buyer"`
	Maker           bool   `json:"maker"`
}
