package types

type GetUnderlyingIndexPriceParams struct {
	Underlying string `url:"underlying" validate:"required"`
}

type UnderlyingIndexPrice struct {
	Time       int64  `json:"time"`
	IndexPrice string `json:"indexPrice"`
}
