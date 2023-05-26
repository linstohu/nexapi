package types

type Account struct {
	Assets []struct {
		Asset                  string `json:"asset"`
		WalletBalance          string `json:"walletBalance"`
		UnrealizedProfit       string `json:"unrealizedProfit"`
		MarginBalance          string `json:"marginBalance"`
		MaintMargin            string `json:"maintMargin"`
		InitialMargin          string `json:"initialMargin"`
		PositionInitialMargin  string `json:"positionInitialMargin"`
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
		MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
		CrossWalletBalance     string `json:"crossWalletBalance"`
		CrossUnPnl             string `json:"crossUnPnl"`
		AvailableBalance       string `json:"availableBalance"`
	} `json:"assets"`
	Positions []struct {
		Symbol                 string `json:"symbol"`
		PositionAmt            string `json:"positionAmt"`
		InitialMargin          string `json:"initialMargin"`
		MaintMargin            string `json:"maintMargin"`
		UnrealizedProfit       string `json:"unrealizedProfit"`
		PositionInitialMargin  string `json:"positionInitialMargin"`
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
		Leverage               string `json:"leverage"`
		Isolated               bool   `json:"isolated"`
		PositionSide           string `json:"positionSide"`
		EntryPrice             string `json:"entryPrice"`
		MaxQty                 string `json:"maxQty"`
		UpdateTime             int64  `json:"updateTime"`
	} `json:"positions"`
	CanDeposit  bool  `json:"canDeposit"`
	CanTrade    bool  `json:"canTrade"`
	CanWithdraw bool  `json:"canWithdraw"`
	FeeTier     int   `json:"feeTier"`
	UpdateTime  int64 `json:"updateTime"`
}
