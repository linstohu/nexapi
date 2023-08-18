package bitfinexrestauth

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
)

var (
	BaseURL = "https://api.bitfinex.com"
)

type RestAuthClient struct {
	*BitfinexAuthClient

	// validate struct fields
	validate *validator.Validate
}

func NewRestAuthClient(cfg *BitfinexClientCfg) (*RestAuthClient, error) {
	cli, err := NewBitfinexClient(cfg)
	if err != nil {
		return nil, err
	}

	validator := validator.New()

	return &RestAuthClient{
		BitfinexAuthClient: cli,
		validate:           validator,
	}, nil
}

func (r *RestAuthClient) GetWallets(ctx context.Context) error {
	req := HTTPRequest{
		BaseURL: r.GetBaseURL(),
		Path:    "/v2/auth/r/wallets",
		Method:  http.MethodPost,
	}

	{
		headers, err := r.GenHeaders(req)
		if err != nil {
			return err
		}
		req.Headers = headers
	}

	_, err := r.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
