package types

import (
	umutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type ChangeMarginTypeParam struct {
	Symbol     string             `url:"symbol" validate:"required"`
	MarginType umutils.MarginType `url:"marginType,omitempty" validate:"omitempty"`
}

type ChangeMarginTypeParams struct {
	ChangeMarginTypeParam
	bnutils.DefaultParam
}

type ModifyIsolatedPositionMarginParam struct {
	Symbol       string               `url:"symbol" validate:"required"`
	PositionSide umutils.PositionSide `url:"positionSide,omitempty" validate:"omitempty"`
	Amount       float64              `url:"amount" validate:"required"`
	Type         int                  `url:"type,omitempty" validate:"required"`
}

type ModifyIsolatedPositionMarginParams struct {
	ModifyIsolatedPositionMarginParam
	bnutils.DefaultParam
}

type ModifyIsolatedPositionMarginResp struct {
	Amount float64 `json:"amount"`
	Code   int     `json:"code"`
	Msg    string  `json:"msg"`
	Type   int     `json:"type"`
}
