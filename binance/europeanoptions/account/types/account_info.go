package types

type AccountInfo struct {
	Asset []struct {
		Asset         string `json:"asset"`
		MarginBalance string `json:"marginBalance"`
		Equity        string `json:"equity"`
		Available     string `json:"available"`
		Locked        string `json:"locked"`
		UnrealizedPNL string `json:"unrealizedPNL"`
	} `json:"asset"`
	Greek []struct {
		Underlying string `json:"underlying"`
		Delta      string `json:"delta"`
		Gamma      string `json:"gamma"`
		Theta      string `json:"theta"`
		Vega       string `json:"vega"`
	} `json:"greek"`
	Time int64 `json:"time"`
}
