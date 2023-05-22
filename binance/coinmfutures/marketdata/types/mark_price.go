package types

type GetMarkPriceParam struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
	Pair   string `url:"pair,omitempty" validate:"omitempty"`
}

type MarkPrice struct {
	Symbol               string `json:"symbol"`
	Pair                 string `json:"pair"`
	MarkPrice            string `json:"markPrice"`
	IndexPrice           string `json:"indexPrice"`
	EstimatedSettlePrice string `json:"estimatedSettlePrice"`
	LastFundingRate      string `json:"lastFundingRate"`
	InterestRate         string `json:"interestRate"`
	NextFundingTime      int64  `json:"nextFundingTime"`
	Time                 int64  `json:"time"`
}
