package types

type GetOpenInterestParam struct {
	UnderlyingAsset string `url:"underlyingAsset" validate:"required"`
	Expiration      string `url:"expiration" validate:"required"`
}

type OpenInterest struct {
	Symbol             string `json:"symbol"`
	SumOpenInterest    string `json:"sumOpenInterest"`
	SumOpenInterestUsd string `json:"sumOpenInterestUsd"`
	Timestamp          string `json:"timestamp"`
}
