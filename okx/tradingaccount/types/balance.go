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

import (
	okxutils "github.com/linstohu/nexapi/okx/utils"
)

type GetBalanceParam struct {
	Currency string `url:"ccy,omitempty"`
}

type GetBalanceResp struct {
	okxutils.Response
	Data []struct {
		AdjEq       string          `json:"adjEq"`
		BorrowFroz  string          `json:"borrowFroz"`
		Details     []BalanceDetail `json:"details"`
		Imr         string          `json:"imr"`
		IsoEq       string          `json:"isoEq"`
		MgnRatio    string          `json:"mgnRatio"`
		Mmr         string          `json:"mmr"`
		NotionalUsd string          `json:"notionalUsd"`
		OrdFroz     string          `json:"ordFroz"`
		TotalEq     string          `json:"totalEq"`
		UTime       string          `json:"uTime"`
	} `json:"data"`
}

type BalanceDetail struct {
	AvailBal      string `json:"availBal"`
	AvailEq       string `json:"availEq"`
	CashBal       string `json:"cashBal"`
	Ccy           string `json:"ccy"`
	CrossLiab     string `json:"crossLiab"`
	DisEq         string `json:"disEq"`
	Eq            string `json:"eq"`
	EqUsd         string `json:"eqUsd"`
	FixedBal      string `json:"fixedBal"`
	FrozenBal     string `json:"frozenBal"`
	Interest      string `json:"interest"`
	IsoEq         string `json:"isoEq"`
	IsoLiab       string `json:"isoLiab"`
	IsoUpl        string `json:"isoUpl"`
	Liab          string `json:"liab"`
	MaxLoan       string `json:"maxLoan"`
	MgnRatio      string `json:"mgnRatio"`
	NotionalLever string `json:"notionalLever"`
	OrdFrozen     string `json:"ordFrozen"`
	Twap          string `json:"twap"`
	UTime         string `json:"uTime"`
	Upl           string `json:"upl"`
	UplLiab       string `json:"uplLiab"`
	StgyEq        string `json:"stgyEq"`
	SpotInUseAmt  string `json:"spotInUseAmt"`
	BorrowFroz    string `json:"borrowFroz"`
}
