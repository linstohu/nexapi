package types

type GetLeverageParams struct {
	Symbol string `url:"symbol,omitempty" validate:"required"`
}

type GetLeverageResp struct {
	Response
	Data struct {
		PositionType int     `json:"positionType"`
		Level        int     `json:"level"`
		Imr          float64 `json:"imr"`
		Mmr          float64 `json:"mmr"`
		Leverage     int     `json:"leverage"`
	} `json:"data"`
}

type SetLeverageParams struct {
	PositionId   int64  `url:"positionId,omitempty" validate:"omitempty"`
	Leverage     int    `url:"positionId,omitempty" validate:"required"`
	OpenType     int    `url:"openType,omitempty" validate:"omitempty"`
	Symbol       string `url:"symbol,omitempty" validate:"omitempty"`
	PositionType int    `url:"positionType,omitempty" validate:"omitempty"`
}

type SetLeverageResp struct {
	Response
	Data struct {
		PositionId   int64  `json:"positionId"`
		Leverage     int    `json:"leverage"`
		Symbol       string `json:"symbol"`
		PositionType int    `json:"positionType"`
	} `json:"data"`
}
