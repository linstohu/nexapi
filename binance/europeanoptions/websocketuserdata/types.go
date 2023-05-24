package websocketuserdata

type AccountData struct {
	EventType string `json:"e"`
	EventTime int64  `json:"E"`
	Balances  []struct {
		AccountBalance    string  `json:"b"`
		PositionValue     string  `json:"m"`
		UnPNL             string  `json:"u"`
		U                 float64 `json:"U"`
		MaintenanceMargin string  `json:"M"`
		InitialMargin     string  `json:"i"`
		MarginAsset       string  `json:"a"`
	} `json:"B"`
	Greek []struct {
		Underlying string  `json:"ui"`
		Delta      float64 `json:"d"`
		Theta      float64 `json:"t"`
		Gamma      float64 `json:"g"`
		VegaV      float64 `json:"v"`
	} `json:"G"`
	Position []struct {
		Symbol                  string `json:"s"`
		PositionNum             string `json:"c"`
		PositionNumCanBeReduced string `json:"r"`
		PositionValue           string `json:"p"`
		EntryPrice              string `json:"a"`
	} `json:"P"`
	UID int64 `json:"uid"`
}

type OrderUpdate struct {
	EventType string `json:"e"`
	EventTime int64  `json:"E"`
	Orders    []struct {
		CreateTime           int64  `json:"T"`
		UpdateTime           int64  `json:"t"`
		Symbol               string `json:"s"`
		ClientOrderID        string `json:"c"`
		OrderID              string `json:"oid"`
		OrderPrice           string `json:"p"`
		OrderQuantity        string `json:"q"`
		ReduceOnly           bool   `json:"r"`
		PostOnly             bool   `json:"po"`
		Status               string `json:"S"`
		CompletedTradeVolume string `json:"e"`
		CompletedTradeAmount string `json:"ec"`
		Fee                  string `json:"f"`
		TimeInForce          string `json:"tif"`
		OrderType            string `json:"oty"`
		Fill                 []struct {
			TradeID            string `json:"t"`
			TradePrice         string `json:"p"`
			TradeQuantity      string `json:"q"`
			TradeTime          int64  `json:"T"`
			TakerOrMaker       string `json:"m"`
			CommissionOrRebate string `json:"f"`
		} `json:"fi"`
	} `json:"o"`
}
