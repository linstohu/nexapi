/*
 * Copyright (c) 2023, LinstoHu
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package types

import okxutils "github.com/linstohu/nexapi/okx/utils"

type GetPositionsParam struct {
	InstType PosInstType `url:"instType,omitempty" validate:"omitempty,oneof=MARGIN SWAP FUTURES OPTION"`
	InstId   string      `url:"instId,omitempty"`
	PosId    string      `url:"posId,omitempty"`
}

type PosInstType = string

const (
	Margin  = "MARGIN"
	Swap    = "SWAP"
	Futures = "FUTURES"
	Option  = "OPTION"
)

type GetPositionsResp struct {
	okxutils.Response
	Data []*Position `json:"data"`
}

// Position
// doc: https://www.okx.com/docs-v5/en/#trading-account-rest-api-get-positions
type Position struct {
	InstType        string           `json:"instType,omitempty"`
	MgnMode         string           `json:"mgnMode,omitempty"`
	PosId           string           `json:"posId,omitempty"`
	PosSide         string           `json:"posSide,omitempty"`
	Pos             string           `json:"pos,omitempty"`
	BaseBal         string           `json:"baseBal,omitempty"`
	QuoteBal        string           `json:"quoteBal,omitempty"`
	BaseBorrowed    string           `json:"baseBorrowed,omitempty"`
	BaseInterest    string           `json:"baseInterest,omitempty"`
	QuoteBorrowed   string           `json:"quoteBorrowed,omitempty"`
	QuoteInterest   string           `json:"quoteInterest,omitempty"`
	PosCcy          string           `json:"posCcy,omitempty"`
	AvailPos        string           `json:"availPos,omitempty"`
	AvgPx           string           `json:"avgPx,omitempty"`
	UPL             string           `json:"upl,omitempty"`
	UplRatio        string           `json:"uplRatio,omitempty"`
	UplLastPx       string           `json:"uplLastPx,omitempty"`
	UplRatioLastPx  string           `json:"uplRatioLastPx,omitempty"`
	InstId          string           `json:"instId,omitempty"`
	Lever           string           `json:"lever,omitempty"`
	LiqPx           string           `json:"liqPx,omitempty"`
	MarkPx          string           `json:"markPx,omitempty"`
	IMR             string           `json:"imr,omitempty"`
	Margin          string           `json:"margin,omitempty"`
	MgnRatio        string           `json:"mgnRatio,omitempty"`
	MMR             string           `json:"mmr,omitempty"`
	Liab            string           `json:"liab,omitempty"`
	LiabCcy         string           `json:"liabCcy,omitempty"`
	Interest        string           `json:"interest,omitempty"`
	TradeId         string           `json:"tradeId,omitempty"`
	OptVal          string           `json:"optVal,omitempty"`
	NotionalUsd     string           `json:"notionalUsd,omitempty"`
	ADL             string           `json:"adl,omitempty"`
	CCY             string           `json:"ccy,omitempty"`
	Last            string           `json:"last,omitempty"`
	IdxPx           string           `json:"idxPx,omitempty"`
	UsdPx           string           `json:"usdPx,omitempty"`
	BePx            string           `json:"bePx,omitempty"`
	DeltaBS         string           `json:"deltaBS,omitempty"`
	DeltaPA         string           `json:"deltaPA,omitempty"`
	GammaBS         string           `json:"gammaBS,omitempty"`
	GammaPA         string           `json:"gammaPA,omitempty"`
	ThetaBS         string           `json:"thetaBS,omitempty"`
	ThetaPA         string           `json:"thetaPA,omitempty"`
	VegaBS          string           `json:"vegaBS,omitempty"`
	VegaPA          string           `json:"vegaPA,omitempty"`
	CTime           string           `json:"cTime,omitempty"`
	UTime           string           `json:"uTime,omitempty"`
	SpotInUseAmt    string           `json:"spotInUseAmt,omitempty"`
	SpotInUseCcy    string           `json:"spotInUseCcy,omitempty"`
	RealizedPnl     string           `json:"realizedPnl,omitempty"`
	PNL             string           `json:"pnl,omitempty"`
	Fee             string           `json:"fee,omitempty"`
	FundingFee      string           `json:"fundingFee,omitempty"`
	LiqPenalty      string           `json:"liqPenalty,omitempty"`
	CloseOrderAlgos []CloseOrderAlgo `json:"closeOrderAlgo,omitempty"`
	BizRefId        string           `json:"bizRefId,omitempty"`
	BizRefType      string           `json:"bizRefType,omitempty"`
}

type CloseOrderAlgo struct {
	AlgoId          string `json:"algoId,omitempty"`
	SlTriggerPx     string `json:"slTriggerPx,omitempty"`
	SlTriggerPxType string `json:"slTriggerPxType,omitempty"`
	TpTriggerPx     string `json:"tpTriggerPx,omitempty"`
	TpTriggerPxType string `json:"tpTriggerPxType,omitempty"`
	CloseFraction   string `json:"closeFraction,omitempty"`
}
