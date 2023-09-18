package okxutils

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

	"github.com/go-playground/validator"
	"github.com/google/go-querystring/query"
)

type OKXRestClient struct {
	baseURL                 string
	key, secret, passphrase string
	// debug mode
	debug bool
	// logger
	logger *log.Logger
	// validate struct fields
	validate *validator.Validate
}

type OKXRestClientCfg struct {
	BaseURL    string `validate:"required"`
	Key        string
	Secret     string
	Passphrase string
	Debug      bool
	// Logger
	Logger *log.Logger
}

func NewOKXRestClient(cfg *OKXRestClientCfg) (*OKXRestClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli := OKXRestClient{
		baseURL:    cfg.BaseURL,
		key:        cfg.Key,
		secret:     cfg.Secret,
		passphrase: cfg.Passphrase,
		debug:      cfg.Debug,
		logger:     cfg.Logger,

		validate: validator,
	}

	if cli.logger == nil {
		cli.logger = log.Default()
		cli.logger.SetPrefix("okx-rest-api")
	}

	return &cli, nil
}

func (o *OKXRestClient) GetDebug() bool {
	return o.debug
}

func (o *OKXRestClient) GetBaseURL() string {
	return o.baseURL
}

func (o *OKXRestClient) GetKey() string {
	return o.key
}

func (o *OKXRestClient) GetSecret() string {
	return o.secret
}

func (o *OKXRestClient) GetPassphrase() string {
	return o.passphrase
}

func (o *OKXRestClient) SendHTTPRequest(ctx context.Context, req HTTPRequest) ([]byte, error) {
	client := http.Client{}

	var body io.Reader
	if req.Body != nil {
		jsonBody, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(jsonBody)
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

	if o.debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		o.logger.Printf("\n%s\n", string(dump))
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if o.debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		o.logger.Printf("\n%s\n", string(dump))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned a non-200 status code: [%d] - [%s]", resp.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}

func (o *OKXRestClient) GenPubHeaders() (map[string]string, error) {
	return map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}, nil
}

func (o *OKXRestClient) GenAuthHeaders(req HTTPRequest) (map[string]string, error) {
	if o.key == "" || o.secret == "" {
		return nil, fmt.Errorf("key and secret needed when init client")
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	strBody := ""
	if req.Body != nil {
		jsonBody, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		strBody = string(jsonBody)
	}

	path := req.Path
	if req.Query != nil {
		// attention: do not forget url tag after struct's fields
		q, err := query.Values(req.Query)
		if err != nil {
			return nil, err
		}

		encode := q.Encode()
		if encode != "" {
			path = fmt.Sprintf("%s?%s", req.Path, encode)
		}
	}

	timestamp := time.Now().UnixMilli()
	signString := fmt.Sprintf("%d%s%s%s", timestamp, req.Method, path, strBody)

	h := hmac.New(sha256.New, []byte(o.secret))
	h.Write([]byte(signString))
	signature := hex.EncodeToString(h.Sum(nil))

	headers["OK-ACCESS-KEY"] = o.key
	headers["OK-ACCESS-PASSPHRASE"] = o.passphrase
	headers["OK-ACCESS-TIMESTAMP"] = strconv.FormatInt(timestamp, 10)
	headers["OK-ACCESS-SIGN"] = signature

	return headers, nil
}
