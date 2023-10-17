package margin

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/spot/margin/types"
	spotutils "github.com/linstohu/nexapi/binance/spot/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type SpotMarginClient struct {
	*spotutils.SpotClient

	// validate struct fields
	validate *validator.Validate
}

type SpotMarginClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewSpotMarginClient(cfg *SpotMarginClientCfg) (*SpotMarginClient, error) {
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

	return &SpotMarginClient{
		SpotClient: cli,
		validate:   validator,
	}, nil
}

func (s *SpotMarginClient) GetInterestHistory(ctx context.Context, param types.GetInterestHistoryParam) (*types.InterestHistory, error) {
	req := spotutils.HTTPRequest{
		SecurityType: spotutils.USER_DATA,
		BaseURL:      s.GetBaseURL(),
		Path:         "/sapi/v1/margin/interestHistory",
		Method:       http.MethodGet,
	}

	{
		headers, err := s.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := types.GetInterestHistoryParams{
			GetInterestHistoryParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: s.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := s.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := s.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(s.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.InterestHistory
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
