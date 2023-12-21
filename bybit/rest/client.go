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

package rest

import (
	"github.com/go-playground/validator"
	bybitutils "github.com/linstohu/nexapi/bybit/utils"
)

type BybitClient struct {
	cli *bybitutils.BybitClient

	// validate struct fields
	validate *validator.Validate
}

func NewBybitClient(cfg *bybitutils.BybitClientCfg) (*BybitClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := bybitutils.NewBybitClient(cfg)
	if err != nil {
		return nil, err
	}

	return &BybitClient{
		cli:      cli,
		validate: validator,
	}, nil
}
