package types

type GetAssetHisotryParam struct {
	Token        string `url:"token,omitempty"`
	BalanceToken string `url:"balance_token,omitempty"`
	Type         string `url:"type,omitempty"`
	TokenSide    string `url:"token_side,omitempty"`
	Status       string `url:"status,omitempty"`
	StartTime    int64  `url:"start_t,omitempty" validate:"omitempty,gt=999999999999"`
	EndTime      int64  `url:"end_t,omitempty" validate:"omitempty,gt=999999999999"`
	Page         int64  `url:"page,omitempty"`
	Size         int64  `url:"size,omitempty"`
}

type AssetHisotryResp struct {
	Response
	Meta struct {
		RecordsPerPage int `json:"records_per_page"`
		CurrentPage    int `json:"current_page"`
	} `json:"meta"`
	Rows []Hisotry `json:"rows"`
}

type Hisotry struct {
	CreatedTime         string  `json:"created_time"`
	UpdatedTime         string  `json:"updated_time"`
	ID                  string  `json:"id"`
	ExternalID          string  `json:"external_id"`
	ApplicationID       string  `json:"application_id"`
	Token               string  `json:"token"`
	TargetAddress       string  `json:"target_address"`
	SourceAddress       string  `json:"source_address"`
	ConfirmingThreshold int     `json:"confirming_threshold"`
	ConfirmedNumber     int     `json:"confirmed_number"`
	Extra               string  `json:"extra"`
	Type                string  `json:"type"`
	TokenSide           string  `json:"token_side"`
	Amount              float64 `json:"amount"`
	TxID                string  `json:"tx_id"`
	FeeToken            string  `json:"fee_token"`
	FeeAmount           float64 `json:"fee_amount"`
	Status              string  `json:"status"`
}
