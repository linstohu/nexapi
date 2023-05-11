package types

type GetAccountInfo struct {
	Response
	Data struct {
		ApplicationID          string  `json:"applicationId"`
		Account                string  `json:"account"`
		Alias                  string  `json:"alias"`
		AccountMode            string  `json:"accountMode"`
		Leverage               int     `json:"leverage"`
		TakerFeeRate           float64 `json:"takerFeeRate"`
		MakerFeeRate           float64 `json:"makerFeeRate"`
		InterestRate           float64 `json:"interestRate"`
		FuturesTakerFeeRate    float64 `json:"futuresTakerFeeRate"`
		FuturesMakerFeeRate    float64 `json:"futuresMakerFeeRate"`
		Otpauth                bool    `json:"otpauth"`
		MarginRatio            float64 `json:"marginRatio"`
		OpenMarginRatio        float64 `json:"openMarginRatio"`
		InitialMarginRatio     float64 `json:"initialMarginRatio"`
		MaintenanceMarginRatio float64 `json:"maintenanceMarginRatio"`
		TotalCollateral        float64 `json:"totalCollateral"`
		FreeCollateral         float64 `json:"freeCollateral"`
		TotalAccountValue      float64 `json:"totalAccountValue"`
		TotalVaultValue        float64 `json:"totalVaultValue"`
		TotalStakingValue      float64 `json:"totalStakingValue"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}
