package api

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/go-querystring/query"
	"github.com/linstohu/nexapi/woox/api/types"
)

type WooXClient struct {
	basePath    string
	key, secret string
	// debug mode
	debug bool
	// logger
	logger *log.Logger
	// validate struct fields
	validate *validator.Validate
}

type WooXCfg struct {
	BasePath string `validate:"required"`
	Key      string
	Secret   string
	Debug    bool
	// Logger
	Logger *log.Logger
}

func NewWooXClient(cfg *WooXCfg) (*WooXClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := WooXClient{
		basePath: cfg.BasePath,
		key:      cfg.Key,
		secret:   cfg.Secret,
		debug:    cfg.Debug,
		logger:   cfg.Logger,

		validate: validator,
	}

	if cli.logger == nil {
		cli.logger = log.Default()
	}

	return &cli, nil
}

func (w *WooXClient) SendHTTPRequest(ctx context.Context, req types.HTTPRequest) ([]byte, error) {
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
		w.logger.Printf("\n%s\n", string(dump))
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
		w.logger.Printf("\n%s\n", string(dump))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned a non-200 status code: [%d] - [%s]", resp.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}

func (w *WooXClient) GenV1APIAuthHeaders(req types.HTTPRequest) (map[string]string, error) {
	if w.key == "" || w.secret == "" {
		return nil, fmt.Errorf("key and secret needed when init client")
	}

	headers := DefaultContentType
	signString, err := NormalizeV1RequestContent(req)
	if err != nil {
		return nil, err
	}
	timestamp := time.Now().UnixMilli()
	signString = fmt.Sprintf("%s|%d", signString, timestamp)

	h := hmac.New(sha256.New, []byte(w.secret))
	h.Write([]byte(signString))
	signature := hex.EncodeToString(h.Sum(nil))

	headers["x-api-key"] = w.key
	headers["x-api-signature"] = signature
	headers["x-api-timestamp"] = strconv.FormatInt(timestamp, 10)

	return headers, nil
}

func NormalizeV1RequestContent(req types.HTTPRequest) (string, error) {
	// Get query parameters or body parameters based on HTTP method
	params := make(url.Values)
	switch req.Method {
	case http.MethodGet:
		if req.Query != nil {
			// attention: do not forget url tag after struct's fields
			q, err := query.Values(req.Query)
			if err != nil {
				return "", err
			}
			params = q
		}
	case http.MethodPost, http.MethodDelete:
		if req.Body != nil {
			// attention: do not forget url tag after struct's fields
			q, err := query.Values(req.Body)
			if err != nil {
				return "", err
			}
			params = q
		}
	}

	return params.Encode(), nil
}

func (w *WooXClient) GenV3APIAuthHeaders(req types.HTTPRequest) (map[string]string, error) {
	if w.key == "" || w.secret == "" {
		return nil, fmt.Errorf("key and secret needed when init client")
	}

	headers := DefaultContentType

	var signString string
	timestamp := time.Now().UnixMilli()

	if req.Body != nil {
		jsonBody, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}

		signString = fmt.Sprintf("%d%s%s%s", timestamp, req.Method, req.Path, string(jsonBody))
	}

	h := hmac.New(sha256.New, []byte(w.secret))
	h.Write([]byte(signString))
	signature := hex.EncodeToString(h.Sum(nil))

	headers["x-api-key"] = w.key
	headers["x-api-signature"] = signature
	headers["x-api-timestamp"] = strconv.FormatInt(timestamp, 10)

	return headers, nil
}
