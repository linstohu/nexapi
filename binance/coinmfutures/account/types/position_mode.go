package types

import (
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type ChangePositionModeParam struct {
	DualSidePosition string `url:"dualSidePosition" validate:"required,oneof=true false"`
}

type ChangePositionModeParams struct {
	ChangePositionModeParam
	bnutils.DefaultParam
}

type GetCurrentPositionModeResp struct {
	DualSidePosition bool `json:"dualSidePosition"`
}
