package cmutils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/go-playground/validator"
	"github.com/google/go-querystring/query"
)

type PortfolioMarginClient struct {
	// debug mode
	debug bool
	// logger
	logger *log.Logger

	baseURL     string
	key, secret string
	recvWindow  int
}

type PortfolioMarginClientCfg struct {
	Debug bool
	// Logger
	Logger *log.Logger

	BaseURL    string `validate:"required"`
	Key        string
	Secret     string
	RecvWindow int
}

func NewPortfolioMarginClient(cfg *PortfolioMarginClientCfg) (*PortfolioMarginClient, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := PortfolioMarginClient{
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
		cli.logger = log.Default()
		cli.logger.SetPrefix("binance_portfolio_margin_rest_api")
	}

	return &cli, nil
}

func (p *PortfolioMarginClient) GetDebug() bool {
	return p.debug
}

func (p *PortfolioMarginClient) GetBaseURL() string {
	return p.baseURL
}

func (p *PortfolioMarginClient) GetKey() string {
	return p.key
}

func (p *PortfolioMarginClient) GetSecret() string {
	return p.secret
}

func (p *PortfolioMarginClient) GetRecvWindow() int {
	return p.recvWindow
}

func (p *PortfolioMarginClient) GenHeaders(t SecurityType) (map[string]string, error) {
	headers := DefaultContentType

	// SecurityType each endpoint has a security type that determines how you will interact with it
	// docs: https://binance-docs.github.io/apidocs/pm/en/#endpoint-security-type
	switch t {
	case TRADE, USER_DATA, USER_STREAM:
		key := p.GetKey()
		if key == "" {
			return nil, fmt.Errorf("a valid API-Key required")
		}

		headers["X-MBX-APIKEY"] = p.GetKey()
	}

	return headers, nil
}

func (p *PortfolioMarginClient) NeedSignature(t SecurityType) bool {
	switch t {
	case TRADE, USER_DATA:
		// TRADE and USER_DATA endpoints are SIGNED endpoints.
		return true
	default:
		return false
	}
}

func (p *PortfolioMarginClient) SendHTTPRequest(ctx context.Context, req HTTPRequest) ([]byte, error) {
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

	if p.GetDebug() {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		p.logger.Printf("\n%s\n", string(dump))
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if p.GetDebug() {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		p.logger.Printf("\n%s\n", string(dump))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned a non-200 status code: [%d] - [%s]", resp.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}
