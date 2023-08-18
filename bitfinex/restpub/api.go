package bitfinexrestpub

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
)

var (
	BaseURL = "https://api-pub.bitfinex.com"
)

type RestPubClient struct {
	*BitfinexClient

	// validate struct fields
	validate *validator.Validate
}

func NewRestPubClient(cfg *BitfinexClientCfg) (*RestPubClient, error) {
	cli, err := NewBitfinexClient(cfg)
	if err != nil {
		return nil, err
	}

	validator := validator.New()

	return &RestPubClient{
		BitfinexClient: cli,
		validate:       validator,
	}, nil
}

func (r *RestPubClient) PlatformStatus(ctx context.Context) error {
	req := HTTPRequest{
		BaseURL: r.GetBaseURL(),
		Path:    "/v2/platform/status",
		Method:  http.MethodGet,
		Headers: make(map[string]string),
	}

	_, err := r.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
