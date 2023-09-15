package spotaccount

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/mexc/spot/spotaccount/types"
	spotutils "github.com/linstohu/nexapi/mexc/spot/utils"
	mexcutils "github.com/linstohu/nexapi/mexc/utils"
)

type SpotAccountClient struct {
	*spotutils.SpotClient

	// validate struct fields
	validate *validator.Validate
}

type SpotAccountClientCfg struct {
	Debug bool
	// Logger
	Logger *log.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewSpotAccountClient(cfg *SpotAccountClientCfg) (*SpotAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := spotutils.NewSpotClient(&spotutils.SpotClientCfg{
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

	return &SpotAccountClient{
		SpotClient: cli,
		validate:   validator,
	}, nil
}

func (s *SpotAccountClient) GetAccountInfo(ctx context.Context) (*types.AccountInfo, error) {
	req := spotutils.HTTPRequest{
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/account",
		Method:  http.MethodGet,
	}

	{
		headers, err := s.GenAuthHeaders(req)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := mexcutils.DefaultParam{
			RecvWindow: s.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
		}

		err := s.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		signString, err := mexcutils.NormalizeRequestContent(query, nil)
		if err != nil {
			return nil, err
		}

		h := hmac.New(sha256.New, []byte(s.GetSecret()))
		h.Write([]byte(signString))
		signature := hex.EncodeToString(h.Sum(nil))
		query.Signature = signature

		req.Query = query
	}

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.AccountInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SpotAccountClient) Transfer(ctx context.Context, param types.TransferParam) error {
	req := spotutils.HTTPRequest{
		BaseURL: s.GetBaseURL(),
		Path:    "/api/v3/capital/transfer",
		Method:  http.MethodPost,
	}

	{
		headers, err := s.GenAuthHeaders(req)
		if err != nil {
			return err
		}
		req.Headers = headers
	}

	{
		query := types.TransferParams{
			TransferParam: param,
			DefaultParam: mexcutils.DefaultParam{
				RecvWindow: s.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := s.validate.Struct(query)
		if err != nil {
			return err
		}

		signString, err := mexcutils.NormalizeRequestContent(query, nil)
		if err != nil {
			return err
		}

		h := hmac.New(sha256.New, []byte(s.GetSecret()))
		h.Write([]byte(signString))
		signature := hex.EncodeToString(h.Sum(nil))
		query.Signature = signature

		req.Query = query
	}

	_, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
