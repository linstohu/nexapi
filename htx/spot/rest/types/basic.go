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
	htxutils "github.com/linstohu/nexapi/htx/utils"
)

type GetSymbolsParam struct {
	Ts int64 `url:"ts,omitempty" validate:"omitempty"`
}

type GetSymbolsResp struct {
	htxutils.V1Response
	Ts   string   `json:"ts"`
	Data []Symbol `json:"data,omitempty"`
}

type Symbol struct {
	Sc          string `json:"sc,omitempty"`
	Dn          string `json:"dn,omitempty"`
	Bc          string `json:"bc,omitempty"`
	Bcdn        string `json:"bcdn,omitempty"`
	Qc          string `json:"qc,omitempty"`
	Qcdn        string `json:"qcdn,omitempty"`
	State       string `json:"state,omitempty"`
	Whe         bool   `json:"whe,omitempty"`
	Cd          bool   `json:"cd,omitempty"`
	Te          bool   `json:"te,omitempty"`
	Toa         int64  `json:"toa,omitempty"`
	Sp          string `json:"sp,omitempty"`
	W           int    `json:"w,omitempty"`
	Tpp         int    `json:"tpp,omitempty"`
	Tap         int    `json:"tap,omitempty"`
	Ttp         int    `json:"ttp,omitempty"`
	Fp          int    `json:"fp,omitempty"`
	SuspendDesc string `json:"suspend_desc,omitempty"`
	Tags        string `json:"tags,omitempty"`
	Lr          any    `json:"lr,omitempty"`
	Smlr        any    `json:"smlr,omitempty"`
	Flr         any    `json:"flr,omitempty"`
	Wr          string `json:"wr,omitempty"`
	D           any    `json:"d,omitempty"`
	Elr         any    `json:"elr,omitempty"`
	P           any    `json:"p,omitempty"`
}

type GetSymbolInfoResp struct {
	htxutils.V1Response
	Ts   string       `json:"ts"`
	Data []SymbolInfo `json:"data,omitempty"`
}

type SymbolInfo struct {
	Symbol  string  `json:"symbol"`
	State   string  `json:"state"`
	Bc      string  `json:"bc"`
	Qc      string  `json:"qc"`
	Pp      int     `json:"pp"`
	Ap      int     `json:"ap"`
	Sp      string  `json:"sp"`
	Vp      int     `json:"vp"`
	Minoa   float64 `json:"minoa"`
	Maxoa   float64 `json:"maxoa"`
	Minov   float64 `json:"minov"`
	Lominoa float64 `json:"lominoa"`
	Lomaxoa float64 `json:"lomaxoa"`
	Lomaxba float64 `json:"lomaxba"`
	Lomaxsa float64 `json:"lomaxsa"`
	Smminoa float64 `json:"smminoa"`
	Blmlt   float64 `json:"blmlt"`
	Slmgt   float64 `json:"slmgt"`
	Smmaxoa float64 `json:"smmaxoa"`
	Bmmaxov float64 `json:"bmmaxov"`
	Msormlt float64 `json:"msormlt"`
	Mbormlt float64 `json:"mbormlt"`
	Maxov   float64 `json:"maxov"`
	U       string  `json:"u"`
	Mfr     float64 `json:"mfr"`
	Ct      string  `json:"ct"`
	Rt      string  `json:"rt"`
	Rthr    float64 `json:"rthr"`
	In      float64 `json:"in"`
	At      string  `json:"at"`
	Tags    string  `json:"tags"`
}
