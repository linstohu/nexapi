package types

import "github.com/linstohu/nexapi/binance/utils"

type GetSubAccountTransferHistoryParam struct {
	Asset     string `url:"asset,omitempty" validate:"omitempty"`
	Type      int    `url:"type,omitempty" validate:"omitempty"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty"`
}

type GetSubAccountTransferHistoryParams struct {
	GetSubAccountTransferHistoryParam
	utils.DefaultParam
}

type SubAccountTransferHistory struct {
	CounterParty    string `json:"counterParty"`
	Email           string `json:"email"`
	Type            int    `json:"type"`
	Asset           string `json:"asset"`
	Qty             string `json:"qty"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Status          string `json:"status"`
	TranID          int64  `json:"tranId"`
	Time            int64  `json:"time"`
}
