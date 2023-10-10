package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/go-querystring/query"
	"github.com/linstohu/nexapi/deribit/rest/types"
	"github.com/linstohu/nexapi/deribit/rest/types/auth"
)

type DeribitRestClient struct {
	baseURL     string
	key, secret string
	// debug mode
	debug bool
	// logger
	logger *log.Logger
	// validate struct fields
	validate *validator.Validate

	auth struct {
		token string
	}
}

type DeribitRestClientCfg struct {
	BaseURL string `validate:"required"`
	Key     string
	Secret  string
	Debug   bool
	// Logger
	Logger *log.Logger
}

func NewDeribitRestClient(cfg *DeribitRestClientCfg) (*DeribitRestClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := DeribitRestClient{
		baseURL: cfg.BaseURL,
		key:     cfg.Key,
		secret:  cfg.Secret,
		debug:   cfg.Debug,
		logger:  cfg.Logger,

		validate: validator,
	}

	if cli.logger == nil {
		cli.logger = log.Default()
		cli.logger.SetPrefix("deribit-rest-api")
	}

	if cfg.Key != "" && cfg.Secret != "" {
		token, err := cli.Auth(context.TODO(), auth.AuthParams{
			GrantType:    "client_credentials",
			ClientID:     cfg.Key,
			ClientSecret: cfg.Secret,
		})
		if err != nil {
			return nil, fmt.Errorf("init private rest client failed, error: %v", err)
		}

		cli.auth.token = token.AccessToken
	}

	return &cli, nil
}

func (d *DeribitRestClient) SendHTTPRequest(ctx context.Context, req types.HTTPRequest) ([]byte, error) {
	client := http.Client{}

	var body io.Reader
	if req.Body != nil {
		jsonBody, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(jsonBody)
	}

	url, err := url.Parse(req.URL)
	if err != nil {
		return nil, err
	}
	if req.Query != nil {
		q, err := query.Values(req.Query)
		if err != nil {
			return nil, err
		}
		url.RawQuery = q.Encode()
	}

	request, err := http.NewRequestWithContext(ctx, req.Method, url.String(), body)
	if err != nil {
		return nil, err
	}

	for k, v := range req.Headers {
		request.Header.Set(k, v)
	}

	if req.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		d.logger.Printf("\n%s\n", string(dump))
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if req.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		d.logger.Printf("\n%s\n", string(dump))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned a non-200 status code: [%d] - [%s]", resp.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}
