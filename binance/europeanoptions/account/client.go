package account

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/linstohu/nexapi/binance/europeanoptions/account/types"
	eoutils "github.com/linstohu/nexapi/binance/europeanoptions/utils"
	usdmutils "github.com/linstohu/nexapi/binance/usdmfutures/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
)

type OptionsAccountClient struct {
	*eoutils.OptionsClient

	// validate struct fields
	validate *validator.Validate
}

type OptionsAccountClientCfg struct {
	Debug bool
	// Logger
	Logger *slog.Logger

	BaseURL    string `validate:"required"`
	Key        string `validate:"required"`
	Secret     string `validate:"required"`
	RecvWindow int
}

func NewOptionsAccountClient(cfg *eoutils.OptionsClientCfg) (*OptionsAccountClient, error) {
	validator := validator.New()

	err := validator.Struct(cfg)
	if err != nil {
		return nil, err
	}

	cli, err := eoutils.NewOptionsClient(&eoutils.OptionsClientCfg{
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

	return &OptionsAccountClient{
		OptionsClient: cli,
		validate:      validator,
	}, nil
}

func (o *OptionsAccountClient) GetAccountInfo(ctx context.Context) (*types.AccountInfo, error) {
	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/account",
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

	var ret types.AccountInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (o *OptionsAccountClient) NewOrder(ctx context.Context, param types.NewOrderParam) (*types.Order, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/order",
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
		query := types.NewOrderParams{
			NewOrderParam: param,
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

func (o *OptionsAccountClient) GetSingleOrder(ctx context.Context, param types.GetSingleOrderParam) (*types.Order, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	if param.OrderID == 0 && param.ClientOrderID == "" {
		return nil, fmt.Errorf("either orderId or clientOrderId must be sent")
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/order",
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
		query := types.GetSingleOrderParams{
			GetSingleOrderParam: param,
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

func (o *OptionsAccountClient) CancelOrder(ctx context.Context, param types.CancelOrderParam) (*types.Order, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	if param.OrderID == 0 && param.ClientOrderID == "" {
		return nil, fmt.Errorf("either orderId or clientOrderId must be sent")
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/order",
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
		query := types.CancelOrderParams{
			CancelOrderParam: param,
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

func (o *OptionsAccountClient) CancelAllOrdersBySymbol(ctx context.Context, param types.CancelAllOrdersParam) error {
	err := o.validate.Struct(param)
	if err != nil {
		return err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/allOpenOrders",
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
		query := types.CancelAllOrdersParams{
			CancelAllOrdersParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(query)
		if err != nil {
			return err
		}

		if need := o.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	_, err = o.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (o *OptionsAccountClient) CancelAllOrdersByUnderlying(ctx context.Context, param types.CancelAllOrdersByUnderlyingParam) error {
	err := o.validate.Struct(param)
	if err != nil {
		return err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/allOpenOrdersByUnderlying",
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
		query := types.CancelAllOrdersByUnderlyingParams{
			CancelAllOrdersByUnderlyingParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: o.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := o.validate.Struct(query)
		if err != nil {
			return err
		}

		if need := o.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(query, nil)
			if err != nil {
				return err
			}

			h := hmac.New(sha256.New, []byte(o.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			query.Signature = signature
		}

		req.Query = query
	}

	_, err = o.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (o *OptionsAccountClient) GetOpenOrders(ctx context.Context, param types.GetCurrentOpenOrdersParam) ([]*types.Order, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/openOrders",
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
		query := types.GetCurrentOpenOrdersParams{
			GetCurrentOpenOrdersParam: param,
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

func (o *OptionsAccountClient) GetOrderHistory(ctx context.Context, param types.GetOrderHistoryParam) ([]*types.Order, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/historyOrders",
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
		query := types.GetOrderHistoryParams{
			GetOrderHistoryParam: param,
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

func (o *OptionsAccountClient) GetPositionInfo(ctx context.Context, param types.GetPositionInfoParam) ([]*types.Position, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/position",
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
		query := types.GetPositionInfoParams{
			GetPositionInfoParam: param,
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

func (o *OptionsAccountClient) GetTradeList(ctx context.Context, param types.GetTradeListParam) ([]*types.UserTrade, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/userTrades",
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

	var ret []*types.UserTrade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (o *OptionsAccountClient) GetExerciseRecord(ctx context.Context, param types.GetExerciseRecordParam) ([]*types.UserTrade, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/exerciseRecord",
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
		query := types.GetExerciseRecordParams{
			GetExerciseRecordParam: param,
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

	var ret []*types.UserTrade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (o *OptionsAccountClient) GetFundingFlow(ctx context.Context, param types.GetFundingFlowParam) ([]*types.FundingFlow, error) {
	err := o.validate.Struct(param)
	if err != nil {
		return nil, err
	}

	req := eoutils.HTTPRequest{
		SecurityType: usdmutils.TRADE,
		BaseURL:      o.GetBaseURL(),
		Path:         "/eapi/v1/bill",
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
		query := types.GetFundingFlowParams{
			GetFundingFlowParam: param,
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

	var ret []*types.FundingFlow
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
