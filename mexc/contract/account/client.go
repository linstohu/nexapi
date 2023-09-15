package account

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/mexc/contract/account/types"
	ctutils "github.com/linstohu/nexapi/mexc/contract/utils"
)

type ContractAccountClient struct {
	*ctutils.ContractClient

	// validate struct fields
	validate *validator.Validate
}

type ContractAccountClientCfg struct {
	Debug bool
	// Logger
	Logger *log.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewContractAccountClient(cfg *ctutils.ContractClientCfg) (*ContractAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := ctutils.NewContractClient(&ctutils.ContractClientCfg{
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

	return &ContractAccountClient{
		ContractClient: cli,
		validate:       validator,
	}, nil
}

func (c *ContractAccountClient) GetAccountAsset(ctx context.Context) (*types.GetAccountAsset, error) {
	req := ctutils.HTTPRequest{
		BaseURL: c.GetBaseURL(),
		Path:    "/api/v1/private/account/assets",
		Method:  http.MethodGet,
	}
	{
		headers, err := c.GenAuthHeaders(req)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := c.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetAccountAsset
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
