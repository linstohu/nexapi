package types

import "github.com/linstohu/nexapi/binance/utils"

type GetInterestHistoryParam struct {
	Asset          string `url:"asset,omitempty" validate:"omitempty"`
	IsolatedSymbol string `url:"isolatedSymbol,omitempty" validate:"omitempty"`
	StartTime      int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime        int64  `url:"endTime,omitempty" validate:"omitempty"`
	Current        int64  `url:"current,omitempty" validate:"omitempty"`
	Size           int64  `url:"size,omitempty" validate:"omitempty"`
	Archived       string `url:"archived,omitempty" validate:"omitempty"`
}

type GetInterestHistoryParams struct {
	GetInterestHistoryParam
	utils.DefaultParam
}

type InterestHistory struct {
	Rows  []Interest `json:"rows"`
	Total int64      `json:"total"`
}

type Interest struct {
	TxId                int64  `json:"txId,omitempty"`
	InterestAccuredTime int64  `json:"interestAccuredTime,omitempty"`
	Asset               string `json:"asset,omitempty"`
	RawAsset            string `json:"rawAsset,omitempty"`
	Principal           string `json:"principal,omitempty"`
	Interest            string `json:"interest,omitempty"`
	InterestRate        string `json:"interestRate,omitempty"`
	Type                string `json:"type,omitempty"`
	IsolatedSymbol      string `json:"isolatedSymbol,omitempty"`
}
