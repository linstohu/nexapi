package types

import (
	"github.com/linstohu/nexapi/binance/utils"
)

type GetTradesParam struct {
	Symbol    string `url:"symbol" validate:"required"`
	OrderID   int64  `url:"orderId,omitempty" validate:"omitempty"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	FromID    int64  `url:"fromId,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type GetTradesParams struct {
	GetTradesParam
	utils.DefaultParam
}

type Trade struct {
	Symbol          string `json:"symbol"`
	ID              int64  `json:"id"`
	OrderID         int64  `json:"orderId"`
	OrderListID     int    `json:"orderListId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	QuoteQty        string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
}
