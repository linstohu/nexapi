package types

import (
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type GetFundingFlowParam struct {
	Currency  string `url:"currency" validate:"required"`
	RecordID  int64  `url:"recordId,omitempty" validate:"omitempty"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type GetFundingFlowParams struct {
	GetFundingFlowParam
	bnutils.DefaultParam
}

type FundingFlow struct {
	ID         string `json:"id"`
	Asset      string `json:"asset"`
	Amount     string `json:"amount"`
	Type       string `json:"type"`
	CreateDate int64  `json:"createDate"`
}
