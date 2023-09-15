package spotutils

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
	// debug mode
	debug bool
	// logger
	logger *log.Logger

	baseURL     string
	key, secret string
	recvWindow  int
}

type SpotClientCfg struct {
	Debug bool
	// Logger
	Logger *log.Logger

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
		cli.recvWindow = 10
	}

	if cli.logger == nil {
		cli.logger = log.Default()
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

func (s *SpotClient) GenPubHeaders() (map[string]string, error) {
	return map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}, nil
}

func (s *SpotClient) GenAuthHeaders(req HTTPRequest) (map[string]string, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	headers["X-MEXC-APIKEY"] = s.key

	return headers, nil
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
