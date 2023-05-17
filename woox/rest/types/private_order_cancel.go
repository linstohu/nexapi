package types

type CancelOrderParam struct {
	Symbol  string `url:"symbol" validate:"required"`
	OrderID int64  `url:"order_id" validate:"required"`
}

type CancelOrderByClientOrderIDParam struct {
	Symbol        string `url:"symbol" validate:"required"`
	ClientOrderID int64  `url:"client_order_id" validate:"required"`
}

type CancelOrdersParam struct {
	Symbol string `url:"symbol" validate:"required"`
}

type CancelOrderResp struct {
	Response
	Status string `json:"status"`
}
