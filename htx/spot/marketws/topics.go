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

package marketws

import (
	"fmt"

	"github.com/go-playground/validator"
	spottypes "github.com/linstohu/nexapi/htx/spot/rest/types"
)

type KlineTopicParam struct {
	Symbol   string                  `validate:"required"`
	Interval spottypes.KlineInterval `validate:"required,oneof=1min 5min 15min 30min 60min 4hour 1day 1mon 1week 1year"`
}

func (m *MarketWsClient) GetKlineTopic(params *KlineTopicParam) (string, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("market.%s.kline.%s", params.Symbol, params.Interval), nil
}

func (m *MarketWsClient) GetBBOTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("market.%s.bbo", symbol), nil
}

type DepthTopicParam struct {
	Symbol string `validate:"required"`
	Type   string `validate:"required,oneof=step0 step1 step2 step3 step4 step5"`
}

func (m *MarketWsClient) GetDepthTopic(params *DepthTopicParam) (string, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("market.%s.depth.%s", params.Symbol, params.Type), nil
}

func (m *MarketWsClient) GetTickerTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("market.%s.ticker", symbol), nil
}

func (m *MarketWsClient) GetMarketTradeTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("market.%s.trade.detail", symbol), nil
}

type MBPDepthUpdateTopicParam struct {
	Symbol string `validate:"required"`
	Level  int    `validate:"required,oneof=5 20 150 400"`
}

func (m *MarketWsClient) GetMBPDepthUpdateTopic(params *MBPDepthUpdateTopicParam) (string, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("market.%s.mbp.%v", params.Symbol, params.Level), nil
}

type MBPDepthRefreshTopicParam struct {
	Symbol string `validate:"required"`
	Level  int    `validate:"required,oneof=5 10 20"`
}

func (m *MarketWsClient) GetMBPRefreshDepthTopic(params *MBPDepthRefreshTopicParam) (string, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("market.%s.mbp.refresh.%v", params.Symbol, params.Level), nil
}
