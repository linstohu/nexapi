package types

type UpdateAccountModeParam struct {
	AccountMode string `url:"account_mode" json:"account_mode" validate:"required,oneof=PURE_SPOT MARGIN FUTURES"`
}
