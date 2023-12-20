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
	usdmtypes "github.com/linstohu/nexapi/htx/usdm/rest/types"
)

type KlineTopicParam struct {
	ContractCode string                  `validate:"required"`
	Interval     usdmtypes.KlineInterval `validate:"required,oneof=1min 5min 15min 30min 1hour 4hour 1day 1mon"`
}

func (m *MarketWsClient) GetKlineTopic(params *KlineTopicParam) (string, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("market.%s.kline.%s", params.ContractCode, params.Interval), nil
}

type DepthTopicParam struct {
	ContractCode string `validate:"required"`
	Type         string `validate:"required,oneof=step0 step1 step2 step3 step4 step5 step6 ste7 step8 step9 step10 step11 step12 step13 step14 step15 step16 step17 step18 step19"`
}

func (m *MarketWsClient) GetDepthTopic(params *DepthTopicParam) (string, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("market.%s.depth.%s", params.ContractCode, params.Type), nil
}

func (m *MarketWsClient) GetBBOTopic(contractCode string) (string, error) {
	if contractCode == "" {
		return "", fmt.Errorf("the contract_code field must be provided")
	}
	return fmt.Sprintf("market.%s.bbo", contractCode), nil
}

func (m *MarketWsClient) GetMarketTradeTopic(contractCode string) (string, error) {
	if contractCode == "" {
		return "", fmt.Errorf("the contract_code field must be provided")
	}
	return fmt.Sprintf("market.%s.trade.detail", contractCode), nil
}
