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

package utils

import (
	"strconv"
	"sync/atomic"
	"time"
)

// v2 types

type NonceGenerator interface {
	GetNonce() string
}

type EpochNonceGenerator struct {
	nonce uint64
}

// GetNonce is a naive nonce producer that takes the current Unix nano epoch
// and counts upwards.
// This is a naive approach because the nonce bound to the currently used API
// key and as such needs to be synchronised with other instances using the same
// key in order to avoid race conditions.
func (u *EpochNonceGenerator) GetNonce() string {
	return strconv.FormatUint(atomic.AddUint64(&u.nonce, 1), 10)
}

func NewEpochNonceGenerator() *EpochNonceGenerator {
	return &EpochNonceGenerator{
		nonce: uint64(time.Now().Unix()) * 1000000,
	}
}

// v1 support

var nonce uint64

func init() {
	nonce = uint64(time.Now().UnixNano()) * 1000000
}

// GetNonce is a naive nonce producer that takes the current Unix nano epoch
// and counts upwards.
// This is a naive approach because the nonce bound to the currently used API
// key and as such needs to be synchronised with other instances using the same
// key in order to avoid race conditions.
func GetNonce() string {
	return strconv.FormatUint(atomic.AddUint64(&nonce, 1), 10)
}
