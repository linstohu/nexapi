package types

type Tokens struct {
	Response
	Rows []struct {
		Token         string `json:"token"`
		Fullname      string `json:"fullname"`
		Decimals      int    `json:"decimals"`
		BalanceToken  string `json:"balance_token"`
		CreatedTime   string `json:"created_time"`
		UpdatedTime   string `json:"updated_time"`
		CanCollateral bool   `json:"can_collateral"`
		CanShort      bool   `json:"can_short"`
	} `json:"rows"`
}
