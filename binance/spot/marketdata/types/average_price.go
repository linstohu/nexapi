package types

type GetAvgPriceParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type AvgPrice struct {
	Mins  int64  `json:"mins"`
	Price string `json:"price"`
}
