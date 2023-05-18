package types

type GetUserAssetParam struct {
	Asset            string `url:"asset,omitempty" validate:"omitempty"`
	NeedBtcValuation string `url:"needBtcValuation,omitempty" validate:"omitempty"`
}

type GetUserAssetParams struct {
	GetUserAssetParam
	DefaultParam
}

type UserAsset struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	Ipoable      string `json:"ipoable"`
	BtcValuation string `json:"btcValuation"`
}
