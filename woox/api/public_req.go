package api

import (
	"encoding/json"
	"net/http"

	"github.com/linstohu/nexapi/woox/api/types"
	"github.com/linstohu/nexapi/woox/api/utils"
)

func (w *WooXClient) GetPublicMarketTrades(params *types.GetMarketTradesParam) (*types.MarketTrade, error) {
	err := w.validate.Struct(params)
	if err != nil {
		return nil, err
	}

	req := utils.HTTPRequest{
		URL:     w.basePath + "/v1/public/market_trades",
		Method:  http.MethodGet,
		Headers: DefaultContentType,
		Query:   params,
	}

	resp, err := utils.SendHTTPRequest(w.debug, req)
	if err != nil {
		return nil, err
	}

	var ret types.MarketTrade
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
