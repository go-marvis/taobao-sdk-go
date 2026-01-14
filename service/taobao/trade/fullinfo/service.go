package fullinfo

import (
	"context"
	"net/http"

	"github.com/go-marvis/taobao-sdk-go/core"
)

type Service struct {
	config *core.Config
}

func NewService(config *core.Config) *Service {
	return &Service{config}
}

// Get 获取单笔交易的详细信息
// https://open.taobao.com/api.htm?docId=54&docType=2&scopeId=12146
func (s *Service) Get(ctx context.Context, req *GetReq, options ...core.ReqOptionFunc) (*GetResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "taobao.trade.fullinfo.get"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)

	resp := &GetResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
