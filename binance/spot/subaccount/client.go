package subaccount

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/spot/subaccount/types"
	"github.com/linstohu/nexapi/binance/spot/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type SpotSubAccountClient struct {
	*utils.SpotClient

	// validate struct fields
	validate *validator.Validate
}

func NewSpotSubAccountClient(cfg *utils.SpotClientCfg) (*SpotSubAccountClient, error) {
	cli, err := utils.NewSpotClient(cfg)
	if err != nil {
		return nil, err
	}

	validator := validator.New()

	return &SpotSubAccountClient{
		SpotClient: cli,
		validate:   validator,
	}, nil
}

func (s *SpotSubAccountClient) GetSubAccountTransferHistory(ctx context.Context, param types.GetSubAccountTransferHistoryParam) ([]*types.SubAccountTransferHistory, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.USER_DATA,
		BaseURL:      s.GetBaseURL(),
		Path:         "/sapi/v1/sub-account/transfer/subUserHistory",
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
		query := types.GetSubAccountTransferHistoryParams{
			GetSubAccountTransferHistoryParam: param,
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

	var ret []*types.SubAccountTransferHistory
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
