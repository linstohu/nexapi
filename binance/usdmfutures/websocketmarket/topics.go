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

package websocketmarket

import (
	"fmt"

	"github.com/go-playground/validator"
	umutils "github.com/linstohu/nexapi/binance/spot/utils"
)

func (u *USDMarginedMarketStreamClient) GetAggTradeTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@aggTrade", symbol), nil
}

type MarkPriceTopicParam struct {
	Symbol      string `validate:"required"`
	UpdateSpeed string `validate:"required,oneof=1s 3s"`
}

func (u *USDMarginedMarketStreamClient) GetMarkPriceTopic(param *MarkPriceTopicParam) (string, error) {
	err := validator.New().Struct(param)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s@markPrice@%s", param.Symbol, param.UpdateSpeed), nil
}

type AllMarkPriceTopicParam struct {
	UpdateSpeed string `validate:"required,oneof=1s 3s"`
}

func (u *USDMarginedMarketStreamClient) GetAllMarketPriceTopic(param *AllMarkPriceTopicParam) (string, error) {
	err := validator.New().Struct(param)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("!markPrice@arr@%s", param.UpdateSpeed), nil
}

type KlineTopicParam struct {
	Symbol   string                `validate:"required"`
	Interval umutils.KlineInterval `validate:"required,oneof=1m 3m 5m 15m 30m 1h 2h 4h 6h 8h 12h 1d 3d 1w 1M"`
}

func (u *USDMarginedMarketStreamClient) GetKlineTopic(params *KlineTopicParam) (string, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s@kline_%s", params.Symbol, params.Interval), nil
}

func (u *USDMarginedMarketStreamClient) GetMiniTickerTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@miniTicker", symbol), nil
}

func (u *USDMarginedMarketStreamClient) GetAllMarketMiniTickersTopic() (string, error) {
	return "!miniTicker@arr", nil
}

func (u *USDMarginedMarketStreamClient) GetTickerTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@ticker", symbol), nil
}

func (u *USDMarginedMarketStreamClient) GetAllMarketTickersTopic() (string, error) {
	return "!ticker@arr", nil
}

func (u *USDMarginedMarketStreamClient) GetBookTickerTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@bookTicker", symbol), nil
}

func (u *USDMarginedMarketStreamClient) GetAllBookTickersTopic() (string, error) {
	return "!bookTicker", nil
}

func (u *USDMarginedMarketStreamClient) GetLiquidationOrderTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@forceOrder", symbol), nil
}

func (u *USDMarginedMarketStreamClient) GetAllLiquidationOrdersTopic() (string, error) {
	return "!forceOrder@arr", nil
}

type BookDepthTopicParam struct {
	Symbol      string `validate:"required"`
	Level       int    `validate:"required,oneof=5 10 20"`
	UpdateSpeed string `validate:"required,oneof=100ms 250ms 500ms"`
}

func (u *USDMarginedMarketStreamClient) GetBookDepthTopic(params *BookDepthTopicParam) (string, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s@depth%d@%s", params.Symbol, params.Level, params.UpdateSpeed), nil
}

type BookDiffDepthTopicParam struct {
	Symbol      string `validate:"required"`
	UpdateSpeed string `validate:"required,oneof=100ms 250ms 500ms"`
}

func (u *USDMarginedMarketStreamClient) GetBookDiffDepthTopic(params *BookDiffDepthTopicParam) (string, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s@depth@%s", params.Symbol, params.UpdateSpeed), nil
}
