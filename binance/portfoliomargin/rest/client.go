package pmrest

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/portfoliomargin/rest/types"
	pmutils "github.com/linstohu/nexapi/binance/portfoliomargin/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type PortfolioMarginAccountClient struct {
	*pmutils.PortfolioMarginClient

	// validate struct fields
	validate *validator.Validate
}

type PortfolioMarginAccountClientCfg struct {
	Debug bool
	// Logger
	Logger *log.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewPortfolioMarginAccountClient(cfg *pmutils.PortfolioMarginClientCfg) (*PortfolioMarginAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := pmutils.NewPortfolioMarginClient(&pmutils.PortfolioMarginClientCfg{
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

	return &PortfolioMarginAccountClient{
		PortfolioMarginClient: cli,
		validate:              validator,
	}, nil
}

func (p *PortfolioMarginAccountClient) GetAssetBalance(ctx context.Context, param types.GetBalanceParam) (*types.Balance, error) {
	err := p.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := pmutils.HTTPRequest{
		SecurityType: pmutils.USER_DATA,
		BaseURL:      p.GetBaseURL(),
		Path:         "/papi/v1/balance",
		Method:       http.MethodGet,
	}
	{
		headers, err := p.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetBalanceParams{
			GetBalanceParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: p.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := p.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := p.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(p.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := p.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Balance
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (p *PortfolioMarginAccountClient) GetBalance(ctx context.Context) ([]*types.Balance, error) {
	req := pmutils.HTTPRequest{
		SecurityType: pmutils.USER_DATA,
		BaseURL:      p.GetBaseURL(),
		Path:         "/papi/v1/balance",
		Method:       http.MethodGet,
	}
	{
		headers, err := p.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetBalanceParams{
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: p.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := p.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := p.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(p.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := p.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.Balance
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (p *PortfolioMarginAccountClient) GetAccountInformation(ctx context.Context) (*types.Account, error) {
	req := pmutils.HTTPRequest{
		SecurityType: pmutils.USER_DATA,
		BaseURL:      p.GetBaseURL(),
		Path:         "/papi/v1/account",
		Method:       http.MethodGet,
	}
	{
		headers, err := p.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := bnutils.DefaultParam{
			RecvWindow: p.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
		}

		err := p.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := p.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(p.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := p.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Account
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
