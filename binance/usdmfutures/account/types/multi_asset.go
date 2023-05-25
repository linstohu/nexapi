package types

import (
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type ChangeMultiAssetsModeParam struct {
	MultiAssetsMargin string `url:"multiAssetsMargin" validate:"required,oneof=ture false"`
}

type ChangeMultiAssetsModeParams struct {
	ChangeMultiAssetsModeParam
	bnutils.DefaultParam
}

type GetCurrentMultiAssetsModeResp struct {
	MultiAssetsMargin bool `json:"multiAssetsMargin"`
}
