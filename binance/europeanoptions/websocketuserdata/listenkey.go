package websocketuserdata

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	eoutils "github.com/linstohu/nexapi/binance/europeanoptions/utils"
	usdmutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
)

type httpAuthClient struct {
	*eoutils.OptionsClient
}

type httpAuthClientCfg struct {
	Debug bool
	// Logger
	Logger *log.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func newHttpAuthClient(cfg *httpAuthClientCfg) (*httpAuthClient, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := eoutils.NewOptionsClient(&eoutils.OptionsClientCfg{
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

	return &httpAuthClient{
		OptionsClient: cli,
	}, nil
}

type listenKeyResp struct {
	ListenKey string `json:"listenKey,omitempty"`
}

func (h *httpAuthClient) genListenKey(ctx context.Context) (*listenKeyResp, error) {
	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.USER_STREAM,
		BaseURL:      h.GetBaseURL(),
		Path:         "/eapi/v1/listenKey",
		Method:       http.MethodPost,
	}
	{
		headers, err := h.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	resp, err := h.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret listenKeyResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (h *httpAuthClient) updateListenKey(ctx context.Context) error {
	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.USER_STREAM,
		BaseURL:      h.GetBaseURL(),
		Path:         "/eapi/v1/listenKey",
		Method:       http.MethodPut,
	}
	{
		headers, err := h.GenHeaders(req.SecurityType)
		if err != nil {
			return err
		}
		req.Headers = headers
	}

	_, err := h.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
