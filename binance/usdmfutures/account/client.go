package account

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
	"github.com/linstohu/nexapi/binance/usdmfutures/account/types"
	"github.com/linstohu/nexapi/binance/usdmfutures/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type UsdMFuturesAccountClient struct {
	*utils.USDMarginedClient

	// validate struct fields
	validate *validator.Validate
}

type UsdMFuturesAccountClientCfg struct {
	Debug bool
	// Logger
	Logger *log.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewUsdMFuturesAccountClient(cfg *utils.USDMarginedClientCfg) (*UsdMFuturesAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := utils.NewUSDMarginedClient(&utils.USDMarginedClientCfg{
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

	return &UsdMFuturesAccountClient{
		USDMarginedClient: cli,
		validate:          validator,
	}, nil
}

func (o *UsdMFuturesAccountClient) ChangePositionMode(ctx context.Context, param types.ChangePositionModeParam) (*types.Response, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/fapi/v1/positionSide/dual",
		Method:       http.MethodPost,
	}
	{
		headers, err := o.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.ChangePositionModeParams{
			ChangePositionModeParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := o.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Response
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (o *UsdMFuturesAccountClient) GetPositionMode(ctx context.Context) (*types.GetCurrentPositionModeResp, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/fapi/v1/positionSide/dual",
		Method:       http.MethodGet,
	}
	{
		headers, err := o.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := bnutils.DefaultParam{
			RecvWindow: o.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
		}

		err := o.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := o.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetCurrentPositionModeResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (o *UsdMFuturesAccountClient) ChangeMultiAssetsMode(ctx context.Context, param types.ChangeMultiAssetsModeParam) (*types.Response, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/fapi/v1/multiAssetsMargin",
		Method:       http.MethodPost,
	}
	{
		headers, err := o.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.ChangeMultiAssetsModeParams{
			ChangeMultiAssetsModeParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(body)
		if err != nil {
			return nil, err
		}

		if need := o.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.Response
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (o *UsdMFuturesAccountClient) GetMultiAssetsMode(ctx context.Context) (*types.GetCurrentMultiAssetsModeResp, error) {
	req := utils.HTTPRequest{
		SecurityType: utils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/fapi/v1/multiAssetsMargin",
		Method:       http.MethodGet,
	}
	{
		headers, err := o.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		query := bnutils.DefaultParam{
			RecvWindow: o.GetRecvWindow(),
			Timestamp:  time.Now().UnixMilli(),
		}

		err := o.validate.Struct(query)
		if err != nil {
			return nil, err
		}

		if need := o.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return nil, err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	resp, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var ret types.GetCurrentMultiAssetsModeResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
