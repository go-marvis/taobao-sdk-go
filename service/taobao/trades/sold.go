package trades

import (
	"context"
	"net/http"

	"github.com/go-marvis/taobao-sdk-go/core"
)

type sold struct {
	config *core.Config
}

func (s *sold) Get(ctx context.Context, req *SoldGetReq, options ...core.ReqOptionFunc) (*SoldGetResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "taobao.trades.sold.get"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)

	resp := &SoldGetResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
