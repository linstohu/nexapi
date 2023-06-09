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
	"github.com/linstohu/nexapi/binance/spot/spotaccount/types"
	spotutils "github.com/linstohu/nexapi/binance/spot/utils"
	bnutils "github.com/linstohu/nexapi/binance/utils"
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

func (s *SpotAccountClient) TestNewOrder(ctx context.Context, param types.NewOrderParam) error {
	req := spotutils.HTTPRequest{
		SecurityType: spotutils.TRADE,
		BaseURL:      s.GetBaseURL(),
		Path:         "/api/v3/order/test",
		Method:       http.MethodPost,
	}

	{
		headers, err := s.GenHeaders(req.SecurityType)
		if err != nil {
			return err
		}
		req.Headers = headers
	}

	{
		body := types.NewOrderParams{
			NewOrderParam: param,
			DefaultParam: bnutils.DefaultParam{
				RecvWindow: s.GetRecvWindow(),
				Timestamp:  time.Now().UnixMilli(),
			},
		}

		err := s.validate.Struct(body)
		if err != nil {
			return err
		}

		if need := s.NeedSignature(req.SecurityType); need {
			signString, err := bnutils.NormalizeRequestContent(nil, body)
			if err != nil {
				return err
			}

			h := hmac.New(sha256.New, []byte(s.GetSecret()))
			h.Write([]byte(signString))
			signature := hex.EncodeToString(h.Sum(nil))

			body.Signature = signature
		}

		req.Body = body
	}

	_, err := s.SendHTTPRequest(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *SpotAccountClient) NewOrder(ctx context.Context, param types.NewOrderParam) (*types.NewOrderResp, error) {
	req := spotutils.HTTPRequest{
		SecurityType: spotutils.TRADE,
		BaseURL:      s.GetBaseURL(),
		Path:         "/api/v3/order",
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
		body := types.NewOrderParams{
			NewOrderParam: param,
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

	var ret types.NewOrderResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SpotAccountClient) CancelOrder(ctx context.Context, param types.CancelOrderParam) (*types.OrderInfo, error) {
	req := spotutils.HTTPRequest{
		SecurityType: spotutils.TRADE,
		BaseURL:      s.GetBaseURL(),
		Path:         "/api/v3/order",
		Method:       http.MethodDelete,
	}

	{
		headers, err := s.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.CancelOrderParams{
			CancelOrderParam: param,
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

	var ret types.OrderInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SpotAccountClient) CancelOrdersOnOneSymbol(ctx context.Context, param types.CancelOrdersOnOneSymbolParam) ([]*types.OrderInfo, error) {
	req := spotutils.HTTPRequest{
		SecurityType: spotutils.TRADE,
		BaseURL:      s.GetBaseURL(),
		Path:         "/api/v3/openOrders",
		Method:       http.MethodDelete,
	}

	{
		headers, err := s.GenHeaders(req.SecurityType)
		if err != nil {
			return nil, err
		}
		req.Headers = headers
	}

	{
		body := types.CancelOrdersOnOneSymbolParams{
			CancelOrdersOnOneSymbolParam: param,
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

	var ret []*types.OrderInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *SpotAccountClient) QueryOrder(ctx context.Context, param types.QueryOrderParam) (*types.Order, error) {
	req := spotutils.HTTPRequest{
		SecurityType: spotutils.USER_DATA,
		BaseURL:      s.GetBaseURL(),
		Path:         "/api/v3/order",
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
		query := types.QueryOrderParams{
			QueryOrderParam: param,
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

	var ret types.Order
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SpotAccountClient) GetOpenOrders(ctx context.Context, param types.GetOpenOrdersParam) ([]*types.Order, error) {
	req := spotutils.HTTPRequest{
		SecurityType: spotutils.USER_DATA,
		BaseURL:      s.GetBaseURL(),
		Path:         "/api/v3/openOrders",
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
		query := types.GetOpenOrdersParams{
			GetOpenOrdersParam: param,
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

	var ret []*types.Order
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *SpotAccountClient) GetAllOrders(ctx context.Context, param types.GetAllOrdersParam) ([]*types.Order, error) {
	req := spotutils.HTTPRequest{
		SecurityType: spotutils.USER_DATA,
		BaseURL:      s.GetBaseURL(),
		Path:         "/api/v3/allOrders",
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
		query := types.GetAllOrdersParams{
			GetAllOrdersParam: param,
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

	var ret []*types.Order
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *SpotAccountClient) GetAccountInfo(ctx context.Context) (*types.AccountInfo, error) {
	req := spotutils.HTTPRequest{
		SecurityType: spotutils.USER_DATA,
		BaseURL:      s.GetBaseURL(),
		Path:         "/api/v3/account",
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

	var ret types.AccountInfo
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SpotAccountClient) GetTradeList(ctx context.Context, param types.GetTradesParam) ([]*types.Trade, error) {
	req := spotutils.HTTPRequest{
		SecurityType: spotutils.USER_DATA,
		BaseURL:      s.GetBaseURL(),
		Path:         "/api/v3/myTrades",
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
		query := types.GetTradesParams{
			GetTradesParam: param,
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

	var ret []*types.Trade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
