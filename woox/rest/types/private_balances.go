package types

type Balance struct {
	Response
	Data struct {
		Holding []struct {
			Token            string  `json:"token"`
			Holding          float64 `json:"holding"`
			Frozen           float64 `json:"frozen"`
			Staked           float64 `json:"staked"`
			Unbonding        float64 `json:"unbonding"`
			Vault            float64 `json:"vault"`
			Interest         float64 `json:"interest"`
			PendingShortQty  float64 `json:"pendingShortQty"`
			PendingLongQty   float64 `json:"pendingLongQty"`
			AvailableBalance float64 `json:"availableBalance"`
			UpdatedTime      float64 `json:"updatedTime"`
		} `json:"holding"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}
