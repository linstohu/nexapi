package eoutils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/go-playground/validator"
	"github.com/google/go-querystring/query"
	usdmutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
)

type OptionsClient struct {
	// debug mode
	debug bool
	// logger
	logger *slog.Logger

	baseURL     string
	key, secret string
	recvWindow  int
}

type OptionsClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string
	Secret     string
	RecvWindow int
}

func NewOptionsClient(cfg *OptionsClientCfg) (*OptionsClient, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := OptionsClient{
		debug:      cfg.Debug,
		logger:     cfg.Logger,
		baseURL:    cfg.BaseURL,
		key:        cfg.Key,
		secret:     cfg.Secret,
		recvWindow: cfg.RecvWindow,
	}

	if cfg.RecvWindow == 0 {
		cli.recvWindow = 5000
	}

	if cli.logger == nil {
		cli.logger = slog.Default()
	}

	return &cli, nil
}

func (o *OptionsClient) GetDebug() bool {
	return o.debug
}

func (o *OptionsClient) GetBaseURL() string {
	return o.baseURL
}

func (o *OptionsClient) GetKey() string {
	return o.key
}

func (o *OptionsClient) GetSecret() string {
	return o.secret
}

func (o *OptionsClient) GetRecvWindow() int {
	return o.recvWindow
}

func (o *OptionsClient) GenHeaders(t usdmutils.SecurityType) (map[string]string, error) {
	headers := usdmutils.DefaultContentType

	// SecurityType each endpoint has a security type that determines how you will interact with it
	// docs: https://binance-docs.github.io/apidocs/voptions/en/#endpoint-security-type
	switch t {
	case usdmutils.TRADE, usdmutils.USER_DATA, usdmutils.USER_STREAM, usdmutils.MARKET_DATA:
		key := o.GetKey()
		if key == "" {
			return nil, fmt.Errorf("a valid API-Key required")
		}

		headers["X-MBX-APIKEY"] = o.GetKey()
	}

	return headers, nil
}

func (o *OptionsClient) NeedSignature(t usdmutils.SecurityType) bool {
	switch t {
	case usdmutils.TRADE, usdmutils.USER_DATA:
		// TRADE and USER_DATA endpoints are SIGNED endpoints.
		return true
	default:
		return false
	}
}

func (o *OptionsClient) SendHTTPRequest(ctx context.Context, req HTTPRequest) ([]byte, error) {
	client := http.Client{}

	var body io.Reader
	if req.Body != nil {
		formData, err := query.Values(req.Body)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(formData.Encode())
	}

	url, err := url.Parse(req.BaseURL + req.Path)
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

	if o.GetDebug() {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		o.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if o.GetDebug() {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		o.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned a non-200 status code: [%d] - [%s]", resp.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}
