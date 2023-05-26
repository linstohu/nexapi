package websocketmarket

import (
	"fmt"

	"github.com/go-playground/validator"
	eoutils "github.com/linstohu/nexapi/binance/europeanoptions/utils"
)

// GetTradeTopic
// The Trade Streams push raw trade information for specific symbol or underlying asset. E.g.ETH@trade
func (o *OptionsMarketStreamClient) GetTradeTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@trade", symbol), nil
}

// GetIndexPriceTopic e.g ETHUSDT
// Underlying index stream.
func (o *OptionsMarketStreamClient) GetIndexPriceTopic(underlyingIndex string) (string, error) {
	if underlyingIndex == "" {
		return "", fmt.Errorf("the UnderlyingIndex field must be provided")
	}
	return fmt.Sprintf("%s@index", underlyingIndex), nil
}

// GetMarkPriceTopic
// The mark price for all option symbols on specific underlying asset. E.g.ETH@markPrice
func (o *OptionsMarketStreamClient) GetMarkPriceTopic(underlying string) (string, error) {
	if underlying == "" {
		return "", fmt.Errorf("the underlying field must be provided")
	}
	return fmt.Sprintf("%s@markPrice", underlying), nil
}

type KlineTopicParam struct {
	Symbol   string                `validate:"required"`
	Interval eoutils.KlineInterval `validate:"required,oneof=1m 3m 5m 15m 30m 1h 2h 4h 6h 12h 1d 3d 1w"`
}

func (o *OptionsMarketStreamClient) GetKlineTopic(params *KlineTopicParam) (string, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s@kline_%s", params.Symbol, params.Interval), nil
}

func (o *OptionsMarketStreamClient) Get24HourTickerTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@ticker", symbol), nil
}

// Get24HourTickerByUnderlyingAndexpirationTopic
// 24hr ticker info by underlying asset and expiration date. E.g.ETH@ticker@220930
func (o *OptionsMarketStreamClient) Get24HourTickerByUnderlyingAndexpirationTopic(underlying, expiration string) (string, error) {
	if underlying == "" || expiration == "" {
		return "", fmt.Errorf("underlying|expiration field must be provided")
	}
	return fmt.Sprintf("%s@ticker@%s", underlying, expiration), nil
}

// GetOpenInterestTopic
// Option open interest for specific underlying asset on specific expiration date. E.g.ETH@openInterest@221125
func (o *OptionsMarketStreamClient) GetOpenInterestTopic(underlying, expiration string) (string, error) {
	if underlying == "" || expiration == "" {
		return "", fmt.Errorf("underlying|expiration field must be provided")
	}
	return fmt.Sprintf("%s@openInterest@%s", underlying, expiration), nil
}

type BookDepthTopicParam struct {
	Symbol      string `validate:"required"`
	Level       int    `validate:"required,oneof=10 20 50 100"`
	UpdateSpeed string `validate:"omitempty,oneof=1000ms 100ms"`
}

func (o *OptionsMarketStreamClient) GetBookDepthTopic(params *BookDepthTopicParam) (string, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return "", err
	}

	if params.UpdateSpeed == "" {
		return fmt.Sprintf("%s@depth%d", params.Symbol, params.Level), nil
	} else {
		return fmt.Sprintf("%s@depth%d@%s", params.Symbol, params.Level, params.UpdateSpeed), nil
	}
}

func (o *OptionsMarketStreamClient) GetBookDiffDepthTopic(symbol string) (string, error) {
	if symbol == "" {
		return "", fmt.Errorf("the symbol field must be provided")
	}
	return fmt.Sprintf("%s@depth1000", symbol), nil
}
