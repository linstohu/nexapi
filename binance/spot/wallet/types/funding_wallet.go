package types

type GetFundingAssetParam struct {
	Asset            string `url:"asset,omitempty" validate:"omitempty"`
	NeedBtcValuation string `url:"needBtcValuation,omitempty" validate:"omitempty"`
}

type GetFundingAssetParams struct {
	GetFundingAssetParam
	DefaultParam
}

type FundingAsset struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	BtcValuation string `json:"btcValuation"`
}
