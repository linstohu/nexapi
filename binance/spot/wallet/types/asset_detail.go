package types

import "github.com/linstohu/nexapi/binance/utils"

type GetAssetDetailParam struct {
	Asset string `url:"asset,omitempty" validate:"omitempty"`
}

type AssetDetailParams struct {
	GetAssetDetailParam
	utils.DefaultParam
}

type AssetDetail struct {
	MinWithdrawAmount string `json:"minWithdrawAmount"`
	DepositStatus     bool   `json:"depositStatus"`
	WithdrawFee       string `json:"withdrawFee"`
	WithdrawStatus    bool   `json:"withdrawStatus"`
	DepositTip        string `json:"depositTip"`
}
