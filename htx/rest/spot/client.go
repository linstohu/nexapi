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

package spot

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/htx/rest/spot/types"
	"github.com/linstohu/nexapi/htx/rest/utils"
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

	BaseURL     string `validate:"required"`
	Key         string `validate:"required"`
	Secret      string `validate:"required"`
	SignVersion string `validate:"required"`
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
		SignVersion: cfg.SignVersion,
	})
	if err != nil {
		return nil, err
	}

	return &SpotClient{
		cli:      cli,
		validate: validator,
	}, nil
}

func (scli *SpotClient) GetAccountInfo(ctx context.Context) (*types.GetAccountInfoResponse, error) {
	req := utils.HTTPRequest{
		BaseURL: scli.cli.GetBaseURL(),
		Path:    "/v1/account/accounts",
		Method:  http.MethodGet,
	}

	{
		headers, err := scli.cli.GetHeaders()
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		values, err := scli.cli.GenSignatureValues(req)
		if err != nil {
			return nil, err
		}

		signStr, err := scli.cli.NormalizeRequestContent(req, values)
		if err != nil {
			return nil, err
		}

		h := scli.cli.Sign([]byte(signStr))
		if err != nil {
			return nil, err
		}

		values.Add("Signature", h)
		req.Query = values
	}

	resp, err := scli.cli.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetAccountInfoResponse
	if err := resp.ReadJsonBody(&ret); err != nil {
		return nil, errors.New(resp.Error())
	}

	return &ret, nil
}
