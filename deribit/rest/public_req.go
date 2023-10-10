package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/linstohu/nexapi/deribit/rest/types"
	"github.com/linstohu/nexapi/deribit/rest/types/supporting"
)

func (w *DeribitRestClient) Test(ctx context.Context) (*supporting.TestResponse, error) {
	req := types.HTTPRequest{
		URL:     w.baseURL + "/api/v2/public/test",
		Method:  http.MethodPost,
		Headers: DefaultContentType,
		Body: &types.Body{
			Jsonrpc: JsonRPCVersion,
			Method:  "public/test",
		},
		Debug: w.debug,
	}

	resp, err := w.SendHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	var jsonMsg types.JsonrpcMessage
	if err := json.Unmarshal(resp, &jsonMsg); err != nil {
		return nil, err
	}

	var ret supporting.TestResponse
	if err := json.Unmarshal(jsonMsg.Result, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
