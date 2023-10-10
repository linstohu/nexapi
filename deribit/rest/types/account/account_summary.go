package account

type AccountSummary struct {
	AvailableFunds            float64 `json:"available_funds"`
	AvailableWithdrawalFunds  float64 `json:"available_withdrawal_funds"`
	Balance                   float64 `json:"balance"`
	Currency                  string  `json:"currency"`
	DeltaTotal                float64 `json:"delta_total"`
	DepositAddress            string  `json:"deposit_address"`
	Email                     string  `json:"email"`
	Equity                    float64 `json:"equity"`
	FuturesPl                 float64 `json:"futures_pl"`
	FuturesSessionRpl         float64 `json:"futures_session_rpl"`
	FuturesSessionUpl         float64 `json:"futures_session_upl"`
	ID                        int     `json:"id"`
	InitialMargin             float64 `json:"initial_margin"`
	MaintenanceMargin         float64 `json:"maintenance_margin"`
	MarginBalance             float64 `json:"margin_balance"`
	OptionsDelta              float64 `json:"options_delta"`
	OptionsGamma              float64 `json:"options_gamma"`
	OptionsPl                 float64 `json:"options_pl"`
	OptionsSessionRpl         float64 `json:"options_session_rpl"`
	OptionsSessionUpl         float64 `json:"options_session_upl"`
	OptionsTheta              float64 `json:"options_theta"`
	OptionsVega               float64 `json:"options_vega"`
	PortfolioMarginingEnabled bool    `json:"portfolio_margining_enabled"`
	SessionFunding            float64 `json:"session_funding"`
	SessionRpl                float64 `json:"session_rpl"`
	SessionUpl                float64 `json:"session_upl"`
	SystemName                string  `json:"system_name"`
	TfaEnabled                bool    `json:"tfa_enabled"`
	TotalPl                   float64 `json:"total_pl"`
	Type                      string  `json:"type"`
	Username                  string  `json:"username"`
	Fees                      []Fee   `json:"fees"`
}

type Fee struct {
	Currency       string  `json:"currency"`
	FeeType        string  `json:"fee_type"`
	InstrumentType string  `json:"instrument_type"`
	MakerFee       float64 `json:"maker_fee"`
	TakerFee       float64 `json:"taker_fee"`
}

type GetAccountSummaryParams struct {
	Currency string `json:"currency"`
	Extended bool   `json:"extended,omitempty"`
}
