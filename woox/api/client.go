package api

import (
	"encoding/json"
	"log"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/linstohu/nexapi/woox/api/types"
	"github.com/linstohu/nexapi/woox/api/utils"
)

type WooXClient struct {
	basePath    string
	key, secret string
	// debug mode
	debug bool
	// logger
	logger *log.Logger
	// validate struct fields
	validate *validator.Validate
}

type WooXCfg struct {
	BasePath string `validate:"required"`
	Key      string
	Secret   string
	Debug    bool
	// Logger
	Logger *log.Logger
}

func NewWooXClient(cfg *WooXCfg) (*WooXClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := WooXClient{
		basePath: cfg.BasePath,
		key:      cfg.Key,
		secret:   cfg.Secret,
		debug:    cfg.Debug,
		logger:   cfg.Logger,

		validate: validator,
	}

	if cli.logger == nil {
		cli.logger = log.Default()
	}

	return &cli, nil
}

var DefaultContentType = map[string]string{
	"Content-Type": "application/json",
}

func (w *WooXClient) GetPublicMarketTrades(params *types.GetMarketTradesParam) (*types.MarketTrade, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		URL:     w.basePath + "/v1/public/market_trades",
		Method:  http.MethodGet,
		Headers: DefaultContentType,
		Query:   params,
	}

	resp, err := utils.SendHTTPRequest(w.debug, req)
	if err != nil {
		return nil, err
	}

	var ret types.MarketTrade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
