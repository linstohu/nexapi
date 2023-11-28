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

package subaccount

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/spot/subaccount/types"
	spotutils "github.com/linstohu/nexapi/binance/spot/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
	"github.com/linstohu/nexapi/utils"
)

type SpotSubAccountClient struct {
	*spotutils.SpotClient

	// validate struct fields
	validate *validator.Validate
}

type SpotSubAccountClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewSpotSubAccountClient(cfg *SpotSubAccountClientCfg) (*SpotSubAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := spotutils.NewSpotClient(&spotutils.SpotClientCfg{
		Debug:      cfg.Debug,
		Logger:     cfg.Logger,
		BaseURL:    cfg.BaseURL,
		Key:        cfg.Key,
		Secret:     cfg.Secret,
		RecvWindow: cfg.RecvWindow,
	})
	if err != nil {
		return nil, err
	}

	return &SpotSubAccountClient{
		SpotClient: cli,
		validate:   validator,
	}, nil
}

func (s *SpotSubAccountClient) GetSubAccountTransferHistory(ctx context.Context, param types.GetSubAccountTransferHistoryParam) (*types.GetSubAccountTransferHistoryResp, error) {
	req := utils.HTTPRequest{
		Debug:   s.GetDebug(),
		BaseURL: s.GetBaseURL(),
		Path:    "/sapi/v1/sub-account/transfer/subUserHistory",
		Method:  http.MethodGet,
	}

	st := spotutils.USER_DATA

	{
		headers, err := s.GenHeaders(st)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetSubAccountTransferHistoryParams{
			GetSubAccountTransferHistoryParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: s.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := s.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := s.NeedSignature(st); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(s.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var body []*types.SubAccountTransferHistory
	if err := resp.ReadJsonBody(&body); err != nil {
		return nil, err
	}

	data := &types.GetSubAccountTransferHistoryResp{
		Http: resp,
		Body: body,
	}

	return data, nil
}
