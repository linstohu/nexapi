package types

type AccountInfo struct {
	MakerCommission  int `json:"makerCommission"`
	TakerCommission  int `json:"takerCommission"`
	BuyerCommission  int `json:"buyerCommission"`
	SellerCommission int `json:"sellerCommission"`
	CommissionRates  struct {
		Maker  string `json:"maker"`
		Taker  string `json:"taker"`
		Buyer  string `json:"buyer"`
		Seller string `json:"seller"`
	} `json:"commissionRates"`
	CanTrade                   bool   `json:"canTrade"`
	CanWithdraw                bool   `json:"canWithdraw"`
	CanDeposit                 bool   `json:"canDeposit"`
	Brokered                   bool   `json:"brokered"`
	RequireSelfTradePrevention bool   `json:"requireSelfTradePrevention"`
	UpdateTime                 int    `json:"updateTime"`
	AccountType                string `json:"accountType"`
	Balances                   []struct {
		Asset  string `json:"asset"`
		Free   string `json:"free"`
		Locked string `json:"locked"`
	} `json:"balances"`
	Permissions []string `json:"permissions"`
}