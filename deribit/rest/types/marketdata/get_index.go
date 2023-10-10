package marketdata

type GetIndexPriceParams struct {
	IndexName string `json:"index_name"`
}

type GetIndexPriceResponse struct {
	EstimatedDeliveryPrice float64 `json:"estimated_delivery_price"`
	IndexPrice             float64 `json:"index_price"`
}
