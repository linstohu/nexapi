package types

import "github.com/linstohu/nexapi/mexc/utils"

type TransferParam struct {
	FromAccountType string `url:"fromAccountType" validate:"required,oneof=SPOT FUTURES"`
	ToAccountType   string `url:"toAccountType" validate:"required,oneof=SPOT FUTURES"`
	Asset           string `url:"asset" validate:"required"`
	Amount          string `url:"amount" validate:"required"`
}

type TransferParams struct {
	TransferParam
	utils.DefaultParam
}
