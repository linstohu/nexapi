package spotutils

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
)

type SpotClient struct {
	// debug mode
	debug bool
	// logger
	logger *slog.Logger

	baseURL     string
	key, secret string
	recvWindow  int
}

type SpotClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string
	Secret     string
	RecvWindow int
}

func NewSpotClient(cfg *SpotClientCfg) (*SpotClient, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := SpotClient{
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

func (s *SpotClient) GetDebug() bool {
	return s.debug
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

func (s *SpotClient) GetRecvWindow() int {
	return s.recvWindow
}

func (s *SpotClient) GenHeaders(t SecurityType) (map[string]string, error) {
	headers := DefaultContentType

	// SecurityType each endpoint has a security type that determines how you will interact with it
	// docs: https://binance-docs.github.io/apidocs/spot/en/#endpoint-security-type
	switch t {
	case TRADE, MARGIN, USER_DATA, USER_STREAM, MARKET_DATA:
		key := s.GetKey()
		if key == "" {
			return nil, fmt.Errorf("a valid API-Key required")
		}

		headers["X-MBX-APIKEY"] = s.GetKey()
	}

	return headers, nil
}

func (s *SpotClient) NeedSignature(t SecurityType) bool {
	switch t {
	case TRADE, MARGIN, USER_DATA:
		return true
	default:
		return false
	}
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

		s.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
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
		s.logger.Info(fmt.Sprintf("\n%s\n", string(dump)))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned a non-200 status code: [%d] - [%s]", resp.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}
