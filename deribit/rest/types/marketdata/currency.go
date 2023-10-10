package marketdata

type Currency struct {
	CoinType                       string               `json:"coin_type"`
	Currency                       string               `json:"currency"`
	CurrencyLong                   string               `json:"currency_long"`
	DisabledDepositAddressCreation bool                 `json:"disabled_deposit_address_creation"`
	FeePrecision                   int                  `json:"fee_precision"`
	MinConfirmations               int                  `json:"min_confirmations"`
	MinWithdrawalFee               float64              `json:"min_withdrawal_fee"`
	WithdrawalFee                  float64              `json:"withdrawal_fee"`
	WithdrawalPriorities           []WithdrawalPriority `json:"withdrawal_priorities"`
}

type WithdrawalPriority struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}
