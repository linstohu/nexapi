package publicdata

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/okx/tradingaccount/types"
	okxutils "github.com/linstohu/nexapi/okx/utils"
)

type TradingAccountClient struct {
	*okxutils.OKXRestClient

	// validate struct fields
	validate *validator.Validate
}

type TradingAccountClientCfg struct {
	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	Passphrase string `validate:"required"`
	Debug      bool
	// Logger
	Logger *log.Logger
}

func NewTradingAccountClient(cfg *TradingAccountClientCfg) (*TradingAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := okxutils.NewOKXRestClient(&okxutils.OKXRestClientCfg{
		Debug:      cfg.Debug,
		Logger:     cfg.Logger,
		BaseURL:    cfg.BaseURL,
		Key:        cfg.Key,
		Secret:     cfg.Secret,
		Passphrase: cfg.Passphrase,
	})
	if err != nil {
		return nil, err
	}

	return &TradingAccountClient{
		OKXRestClient: cli,
		validate:      validator,
	}, nil
}

func (t *TradingAccountClient) GetBalance(ctx context.Context, param types.GetBalanceParam) (*types.GetBalanceResp, error) {
	err := t.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := okxutils.HTTPRequest{
		BaseURL: t.GetBaseURL(),
		Path:    "/api/v5/account/balance",
		Method:  http.MethodGet,
		Query:   param,
	}

	headers, err := t.GenAuthHeaders(req)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := t.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetBalanceResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
