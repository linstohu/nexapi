package bitfinexrestpub

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

type BitfinexClient struct {
	// debug mode
	debug bool
	// logger
	logger *log.Logger

	baseURL string
}

type BitfinexClientCfg struct {
	Debug bool
	// Logger
	Logger *log.Logger

	BaseURL    string `validate:"required"`
	Key        string
	Secret     string
	RecvWindow int
}

func NewBitfinexClient(cfg *BitfinexClientCfg) (*BitfinexClient, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := BitfinexClient{
		debug:   cfg.Debug,
		logger:  cfg.Logger,
		baseURL: cfg.BaseURL,
	}

	if cli.logger == nil {
		cli.logger = log.Default()
	}

	return &cli, nil
}

type HTTPRequest struct {
	BaseURL string
	Path    string
	Method  string
	Headers map[string]string
	Query   any
	Body    any
}

func (b *BitfinexClient) GetDebug() bool {
	return b.debug
}

func (b *BitfinexClient) GetBaseURL() string {
	return b.baseURL
}

func (b *BitfinexClient) SendHTTPRequest(ctx context.Context, req HTTPRequest) ([]byte, error) {
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

	if b.GetDebug() {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		b.logger.Printf("\n%s\n", string(dump))
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if b.GetDebug() {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		b.logger.Printf("\n%s\n", string(dump))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned a non-200 status code: [%d] - [%s]", resp.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}
