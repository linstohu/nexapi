package types

import (
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type ChangeLeverageParam struct {
	Symbol   string `url:"symbol" validate:"required"`
	Leverage int64  `url:"leverage,omitempty" validate:"omitempty,min=1,max=125"`
}

type ChangeLeverageParams struct {
	ChangeLeverageParam
	bnutils.DefaultParam
}

type ChangeLeverageResp struct {
	Leverage         int    `json:"leverage"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	Symbol           string `json:"symbol"`
}
