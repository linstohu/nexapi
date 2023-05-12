package types

type TransferAssetParam struct {
	Token     string  `url:"token" json:"token" validate:"required"`
	Amount    float64 `url:"amount" json:"amount" validate:"required"`
	FromAppID string  `url:"from_application_id" json:"from_application_id" validate:"required"`
	ToAppID   string  `url:"to_application_id" json:"to_application_id" validate:"required"`
}

type TransferAssetResp struct {
	Response
	ID int64 `json:"id"`
}
