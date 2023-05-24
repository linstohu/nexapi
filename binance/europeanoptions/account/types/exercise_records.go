package types

import (
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type GetExerciseRecordParam struct {
	Symbol    string `url:"symbol,omitempty" validate:"omitempty"`
	FromID    int64  `url:"fromId,omitempty" validate:"omitempty"`
	StartTime int64  `url:"startTime,omitempty" validate:"omitempty"`
	EndTime   int64  `url:"endTime,omitempty" validate:"omitempty"`
	Limit     int    `url:"limit,omitempty" validate:"omitempty,max=1000"`
}

type GetExerciseRecordParams struct {
	GetExerciseRecordParam
	bnutils.DefaultParam
}

type ExerciseRecord struct {
	ID            string `json:"id"`
	Currency      string `json:"currency"`
	Symbol        string `json:"symbol"`
	ExercisePrice string `json:"exercisePrice"`
	MarkPrice     string `json:"markPrice"`
	Quantity      string `json:"quantity"`
	Amount        string `json:"amount"`
	Fee           string `json:"fee"`
	CreateDate    int64  `json:"createDate"`
	PriceScale    int    `json:"priceScale"`
	QuantityScale int    `json:"quantityScale"`
	OptionSide    string `json:"optionSide"`
	PositionSide  string `json:"positionSide"`
	QuoteAsset    string `json:"quoteAsset"`
}
