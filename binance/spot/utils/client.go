package utils

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

type SpotClient struct {
	baseURL     string
	key, secret string
	// debug mode
	debug bool
	// logger
	logger *log.Logger
}

type SpotClientCfg struct {
	BaseURL string `validate:"required"`
	Key     string
	Secret  string
	Debug   bool
	// Logger
	Logger *log.Logger
}

func NewSpotClient(cfg *SpotClientCfg) (*SpotClient, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := SpotClient{
		baseURL: cfg.BaseURL,
		key:     cfg.Key,
		secret:  cfg.Secret,
		debug:   cfg.Debug,
		logger:  cfg.Logger,
	}

	if cli.logger == nil {
		cli.logger = log.Default()
	}

	return &cli, nil
}

func (s *SpotClient) GetBaseURL() string {
	return s.baseURL
}

func (s *SpotClient) GetKey() string {
	return s.key
}

func (s *SpotClient) GetSecret() string {
	return s.secret
}

func (s *SpotClient) GetDebug() bool {
	return s.debug
}

func (s *SpotClient) GenHeaders(t SecurityType) (map[string]string, error) {
	headers := DefaultContentType

	// SecurityType each endpoint has a security type that determines how you will interact with it
	// docs: https://binance-docs.github.io/apidocs/spot/en/#endpoint-security-type
	switch t {
	case TRADE, MARGIN, USER_DATA, USER_STREAM, MARKET_DATA:
		headers["X-MBX-APIKEY"] = s.GetKey()
	}

	return headers, nil
}

func NormalizeRequestContent(req HTTPRequest) (string, error) {
	var ret string

	if req.Query != nil {
		// attention: do not forget url tag after struct's fields
		q, err := query.Values(req.Query)
		if err != nil {
			return "", err
		}
		ret += q.Encode()
	}

	if req.Body != nil {
		// attention: do not forget url tag after struct's fields
		q, err := query.Values(req.Body)
		if err != nil {
			return "", err
		}
		ret += q.Encode()
	}

	return ret, nil
}

func (s *SpotClient) SendHTTPRequest(ctx context.Context, req HTTPRequest) ([]byte, error) {
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

	if s.GetDebug() {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		s.logger.Printf("\n%s\n", string(dump))
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if s.GetDebug() {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		s.logger.Printf("\n%s\n", string(dump))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned a non-200 status code: [%d] - [%s]", resp.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}
