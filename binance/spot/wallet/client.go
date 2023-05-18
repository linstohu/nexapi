package wallet

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/spot/utils"
	"github.com/linstohu/nexapi/binance/spot/wallet/types"
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type SpotWalletClient struct {
	*utils.SpotClient

	// validate struct fields
	validate *validator.Validate
}

func NewSpotWalletClient(cfg *utils.SpotClientCfg) (*SpotWalletClient, error) {
	cli, err := utils.NewSpotClient(cfg)
	if err != nil {
		return nil, err
	}

	validator := validator.New()

	return &SpotWalletClient{
		SpotClient: cli,
		validate:   validator,
	}, nil
}

func (s *SpotWalletClient) GetSystemStatus(ctx context.Context) (*types.SystemStatus, error) {
	req := utils.HTTPRequest{
		BaseURL: s.GetBaseURL(),
		Path:    "/sapi/v1/system/status",
		Method:  http.MethodGet,
	}

	headers, err := s.GenHeaders(utils.NONE)
	if err != nil {
		return nil, err
	}
	req.Headers = headers

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.SystemStatus
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SpotWalletClient) GetAllCoinsInfo(ctx context.Context) ([]*types.CoinInfo, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.USER_DATA,
		BaseURL:      s.GetBaseURL(),
		Path:         "/sapi/v1/capital/config/getall",
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
		query := bnutils.DefaultParam{
			RecvWindow: s.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
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

	var ret []*types.CoinInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *SpotWalletClient) GetAssetDetail(ctx context.Context, param types.GetAssetDetailParam) (map[string]*types.AssetDetail, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.USER_DATA,
		BaseURL:      s.GetBaseURL(),
		Path:         "/sapi/v1/asset/assetDetail",
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
		query := types.AssetDetailParams{
			GetAssetDetailParam: param,
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

	ret := make(map[string]*types.AssetDetail)
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *SpotWalletClient) GetTradeFee(ctx context.Context, param types.GetTradeFeeParam) ([]*types.TradeFee, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.USER_DATA,
		BaseURL:      s.GetBaseURL(),
		Path:         "/sapi/v1/asset/tradeFee",
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
		query := types.TradeFeeParams{
			GetTradeFeeParam: param,
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

	var ret []*types.TradeFee
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *SpotWalletClient) UniversalTransfer(ctx context.Context, param types.UniversalTransferParam) (*types.UniversalTransferResp, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.USER_DATA,
		BaseURL:      s.GetBaseURL(),
		Path:         "/sapi/v1/asset/transfer",
		Method:       http.MethodPost,
	}

	{
		headers, err := s.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.UniversalTransferParams{
			UniversalTransferParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: s.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := s.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := s.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(s.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.UniversalTransferResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SpotWalletClient) GetUniversalTransferHistory(ctx context.Context, param types.GetUniversalTransferHistoryParam) (*types.GetUniversalTransferHistory, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.USER_DATA,
		BaseURL:      s.GetBaseURL(),
		Path:         "/sapi/v1/asset/transfer",
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
		query := types.GetUniversalTransferHistoryParams{
			GetUniversalTransferHistoryParam: param,
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

	var ret types.GetUniversalTransferHistory
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SpotWalletClient) GetFundingAsset(ctx context.Context, param types.GetFundingAssetParam) ([]*types.FundingAsset, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.USER_DATA,
		BaseURL:      s.GetBaseURL(),
		Path:         "/sapi/v1/asset/get-funding-asset",
		Method:       http.MethodPost,
	}

	{
		headers, err := s.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.GetFundingAssetParams{
			GetFundingAssetParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: s.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := s.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := s.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(s.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.FundingAsset
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *SpotWalletClient) GetUserAsset(ctx context.Context, param types.GetUserAssetParam) ([]*types.UserAsset, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.USER_DATA,
		BaseURL:      s.GetBaseURL(),
		Path:         "/sapi/v3/asset/getUserAsset",
		Method:       http.MethodPost,
	}

	{
		headers, err := s.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.GetUserAssetParams{
			GetUserAssetParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: s.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := s.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := s.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(s.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret []*types.UserAsset
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *SpotWalletClient) GetApiRestrictions(ctx context.Context) (*types.ApiRestrictions, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.USER_DATA,
		BaseURL:      s.GetBaseURL(),
		Path:         "/sapi/v1/account/apiRestrictions",
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
		query := bnutils.DefaultParam{
			RecvWindow: s.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
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

	var ret *types.ApiRestrictions
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
