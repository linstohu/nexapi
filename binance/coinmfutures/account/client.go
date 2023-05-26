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
	"github.com/linstohu/nexapi/binance/coinmfutures/account/types"
	cmutils "github.com/linstohu/nexapi/binance/coinmfutures/utils"
	umutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type CoinMFuturesAccountClient struct {
	*cmutils.CoinMarginedClient

	// validate struct fields
	validate *validator.Validate
}

type CoinMFuturesAccountClientCfg struct {
	Debug bool
	// Logger
	Logger *log.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewCoinMFuturesAccountClient(cfg *cmutils.CoinMarginedClientCfg) (*CoinMFuturesAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := cmutils.NewCoinMarginedClient(&cmutils.CoinMarginedClientCfg{
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

	return &CoinMFuturesAccountClient{
		CoinMarginedClient: cli,
		validate:           validator,
	}, nil
}

func (o *CoinMFuturesAccountClient) ChangePositionMode(ctx context.Context, param types.ChangePositionModeParam) (*types.Response, error) {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/positionSide/dual",
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

func (o *CoinMFuturesAccountClient) GetPositionMode(ctx context.Context) (*types.GetCurrentPositionModeResp, error) {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/positionSide/dual",
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

func (o *CoinMFuturesAccountClient) NewOrder(ctx context.Context, param types.NewOrderParam) (*types.Order, error) {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/order",
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
		body := types.NewOrderParams{
			NewOrderParam: param,
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

	var ret types.Order
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (o *CoinMFuturesAccountClient) QueryOrder(ctx context.Context, param types.GetOrderParam) (*types.Order, error) {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/order",
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
		query := types.GetOrderParams{
			GetOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
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

	var ret types.Order
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (o *CoinMFuturesAccountClient) QueryOpenOrder(ctx context.Context, param types.GetOrderParam) (*types.Order, error) {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/openOrder",
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
		query := types.GetOrderParams{
			GetOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
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

	var ret types.Order
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (o *CoinMFuturesAccountClient) QueryAllOpenOrders(ctx context.Context, param types.GetAllOpenOrdersParam) ([]*types.Order, error) {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/openOrders",
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
		query := types.GetAllOpenOrdersParams{
			GetAllOpenOrdersParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
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

	var ret []*types.Order
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (o *CoinMFuturesAccountClient) CancelOrder(ctx context.Context, param types.GetOrderParam) (*types.Order, error) {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/order",
		Method:       http.MethodDelete,
	}
	{
		headers, err := o.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.GetOrderParams{
			GetOrderParam: param,
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

	var ret types.Order
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (o *CoinMFuturesAccountClient) CancelAllOpenOrders(ctx context.Context, param types.CancelAllOpenOrdersParam) error {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/allOpenOrders",
		Method:       http.MethodDelete,
	}
	{
		headers, err := o.GenHeaders(req.SecurityType)
		if err != nil {
			return err
		}
		req.Headers = headers
	}

	{
		body := types.CancelAllOpenOrdersParams{
			CancelAllOpenOrdersParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(body)
		if err != nil {
			return err
		}

		if need := o.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	_, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (o *CoinMFuturesAccountClient) GetAllOrders(ctx context.Context, param types.GetAllOrdersParam) ([]*types.Order, error) {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/allOrders",
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
		query := types.GetAllOrdersParams{
			GetAllOrdersParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
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

	var ret []*types.Order
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (o *CoinMFuturesAccountClient) GetBalance(ctx context.Context) ([]*types.Balance, error) {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/balance",
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

	var ret []*types.Balance
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (o *CoinMFuturesAccountClient) GetAccountInformation(ctx context.Context) (*types.Account, error) {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/account",
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

	var ret types.Account
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (o *CoinMFuturesAccountClient) ChangeInitialLeverage(ctx context.Context, param types.ChangeLeverageParam) (*types.ChangeLeverageResp, error) {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/leverage",
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
		body := types.ChangeLeverageParams{
			ChangeLeverageParam: param,
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

	var ret types.ChangeLeverageResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (o *CoinMFuturesAccountClient) ChangeMarginType(ctx context.Context, param types.ChangeMarginTypeParam) error {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/marginType",
		Method:       http.MethodPost,
	}
	{
		headers, err := o.GenHeaders(req.SecurityType)
		if err != nil {
			return err
		}
		req.Headers = headers
	}

	{
		body := types.ChangeMarginTypeParams{
			ChangeMarginTypeParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(body)
		if err != nil {
			return err
		}

		if need := o.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	_, err := o.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (o *CoinMFuturesAccountClient) ModifyIsolatedPositionMargin(ctx context.Context, param types.ModifyIsolatedPositionMarginParam) (*types.ModifyIsolatedPositionMarginResp, error) {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/positionMargin",
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
		body := types.ModifyIsolatedPositionMarginParams{
			ModifyIsolatedPositionMarginParam: param,
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

	var ret types.ModifyIsolatedPositionMarginResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (o *CoinMFuturesAccountClient) GetPositionInformation(ctx context.Context, param types.GetPositionParam) ([]*types.Position, error) {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/positionRisk",
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
		query := types.GetPositionParams{
			GetPositionParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
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

	var ret []*types.Position
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (o *CoinMFuturesAccountClient) GetAccountTradeList(ctx context.Context, param types.GetTradeListParam) ([]*types.Trade, error) {
	req := cmutils.HTTPRequest{
		SecurityType: umutils.USER_DATA,
		BaseURL:      o.GetBaseURL(),
		Path:         "/dapi/v1/userTrades",
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
		query := types.GetTradeListParams{
			GetTradeListParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
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

	var ret []*types.Trade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
