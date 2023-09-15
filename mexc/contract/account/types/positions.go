package types

type GetOpenPositionsParams struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
}

type GetOpenPositions struct {
	Response
	Data []*OpenPosition `json:"data"`
}

type OpenPosition struct {
	PositionID   int64  `json:"positionId"`
	Symbol       string `json:"symbol"`
	PositionType int    `json:"positionType"`
	OpenType     int    `json:"openType"`
	State        int    `json:"state"`

	FrozenVol      float64 `json:"frozenVol"`
	CloseVol       float64 `json:"closeVol"`
	HoldAvgPrice   float64 `json:"holdAvgPrice"`
	CloseAvgPrice  float64 `json:"closeAvgPrice"`
	OpenAvgPrice   float64 `json:"openAvgPrice"`
	LiquidatePrice float64 `json:"liquidatePrice"`
	Oim            float64 `json:"oim"`
	Im             float64 `json:"im"`
	HoldFee        float64 `json:"holdFee"`
	Realised       float64 `json:"realised"`

	HoldVol    float64 `json:"holdVol"`
	Leverage   int     `json:"leverage"`
	CreateTime int64   `json:"createTime"`
	UpdateTime int64   `json:"updateTime"`
	AutoAddIm  bool    `json:"autoAddIm"`
}
