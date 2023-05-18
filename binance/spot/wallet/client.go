package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/linstohu/nexapi/binance/spot/utils"
	"github.com/linstohu/nexapi/binance/spot/wallet/types"
)

type SpotWalletClient struct {
	*utils.SpotClient
}

func NewSpotWalletClient(cfg *utils.SpotClientCfg) (*SpotWalletClient, error) {
	cli, err := utils.NewSpotClient(cfg)
	if err != nil {
		return nil, err
	}

	return &SpotWalletClient{cli}, nil
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
