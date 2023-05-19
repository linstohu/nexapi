package websocketmarket

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/spot/utils"
)

func (m *MarketStreamClient) GetAggTradeTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@aggTrade", symbol), nil
}

func (m *MarketStreamClient) GetTradeTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@trade", symbol), nil
}

type KlineTopicParam struct {
	Symbol   string              `validate:"required"`
	Interval utils.KlineInterval `validate:"required,oneof=1s 1m 3m 5m 15m 30m 1h 2h 4h 6h 8h 12h 1d 3d 1w 1M"`
}

func (m *MarketStreamClient) GetKlineTopic(params *KlineTopicParam) (string, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s@kline_%s", params.Symbol, params.Interval), nil
}

func (m *MarketStreamClient) GetMiniTickerTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@miniTicker", symbol), nil
}

func (m *MarketStreamClient) GetAllMarketMiniTickersTopic() (string, error) {
	return "!miniTicker@arr", nil
}

func (m *MarketStreamClient) GetTickerTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@ticker", symbol), nil
}

func (m *MarketStreamClient) GetAllMarketTickersTopic() (string, error) {
	return "!ticker@arr", nil
}

func (m *MarketStreamClient) GetBookTickerTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@bookTicker", symbol), nil
}

type BookDepthTopicParam struct {
	Symbol      string `validate:"required"`
	Level       int    `validate:"required,oneof=5 10 20"`
	UpdateSpeed string `validate:"required,oneof=1000ms 100ms"`
}

func (m *MarketStreamClient) GetBookDepthTopic(params *BookDepthTopicParam) (string, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s@depth%d@%s", params.Symbol, params.Level, params.UpdateSpeed), nil
}

type BookDiffDepthTopicParam struct {
	Symbol      string `validate:"required"`
	UpdateSpeed string `validate:"required,oneof=1000ms 100ms"`
}

func (m *MarketStreamClient) GetBookDiffDepthTopic(params *BookDiffDepthTopicParam) (string, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s@depth@%s", params.Symbol, params.UpdateSpeed), nil
}
