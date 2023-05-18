package types

import "github.com/linstohu/nexapi/binance/utils"

type GetTradeFeeParam struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
}

type TradeFeeParams struct {
	GetTradeFeeParam
	utils.DefaultParam
}

type TradeFee struct {
	Symbol          string `json:"symbol"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}
