package types

type GetMarkPriceParam struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
}

type MarkPrice struct {
	Symbol         string `json:"symbol"`
	MarkPrice      string `json:"markPrice"`
	BidIV          string `json:"bidIV"`
	AskIV          string `json:"askIV"`
	MarkIV         string `json:"markIV"`
	Delta          string `json:"delta"`
	Theta          string `json:"theta"`
	Gamma          string `json:"gamma"`
	Vega           string `json:"vega"`
	HighPriceLimit string `json:"highPriceLimit"`
	LowPriceLimit  string `json:"lowPriceLimit"`
}
