package types

import (
	okxutils "github.com/linstohu/nexapi/okx/utils"
)

type GetInstrumentsParam struct {
	InstType   okxutils.InstrumentType `url:"instType" validate:"required,oneof=SPOT MARGIN SWAP FUTURES OPTION"`
	Uly        string                  `url:"uly,omitempty"`
	InstFamily string                  `url:"instFamily,omitempty"`
	InstId     string                  `url:"instId,omitempty"`
}

type Instruments struct {
	okxutils.Response
	Data []Instrument `json:"data"`
}

type Instrument struct {
	Alias        string `json:"alias"`
	BaseCcy      string `json:"baseCcy"`
	Category     string `json:"category"`
	CtMult       string `json:"ctMult"`
	CtType       string `json:"ctType"`
	CtVal        string `json:"ctVal"`
	CtValCcy     string `json:"ctValCcy"`
	ExpTime      string `json:"expTime"`
	InstFamily   string `json:"instFamily"`
	InstID       string `json:"instId"`
	InstType     string `json:"instType"`
	Lever        string `json:"lever"`
	ListTime     string `json:"listTime"`
	LotSz        string `json:"lotSz"`
	MaxIcebergSz string `json:"maxIcebergSz"`
	MaxLmtSz     string `json:"maxLmtSz"`
	MaxMktSz     string `json:"maxMktSz"`
	MaxStopSz    string `json:"maxStopSz"`
	MaxTriggerSz string `json:"maxTriggerSz"`
	MaxTwapSz    string `json:"maxTwapSz"`
	MinSz        string `json:"minSz"`
	OptType      string `json:"optType"`
	QuoteCcy     string `json:"quoteCcy"`
	SettleCcy    string `json:"settleCcy"`
	State        string `json:"state"`
	Stk          string `json:"stk"`
	TickSz       string `json:"tickSz"`
	Uly          string `json:"uly"`
}
