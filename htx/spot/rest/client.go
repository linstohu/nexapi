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
	"log/slog"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/htx/utils"
)

type SpotClient struct {
	cli *utils.HTXClient

	// validate struct fields
	validate *validator.Validate
}

type SpotClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL string `validate:"required"`
	Key     string
	Secret  string
}

func NewSpotClient(cfg *SpotClientCfg) (*SpotClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := utils.NewHTXRestClient(&utils.HTXClientCfg{
		Debug:       cfg.Debug,
		Logger:      cfg.Logger,
		BaseURL:     cfg.BaseURL,
		Key:         cfg.Key,
		Secret:      cfg.Secret,
		SignVersion: utils.ApiKeyVersionV2,
	})
	if err != nil {
		return nil, err
	}

	return &SpotClient{
		cli:      cli,
		validate: validator,
	}, nil
}
