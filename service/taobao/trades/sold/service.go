package sold

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

// Get 查询卖家已卖出的交易数据（根据创建时间）
// https://open.taobao.com/api.htm?docId=46&docType=2&scopeId=12146
func (s *Service) Get(ctx context.Context, req *GetReq, options ...core.ReqOptionFunc) (*GetResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "taobao.trades.sold.get"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)

	resp := &GetResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// IncrementGet 查询卖家已卖出的增量交易数据（根据修改时间）
// https://open.taobao.com/api.htm?docId=128&docType=2&scopeId=12146
func (s *Service) IncrementGet(ctx context.Context, req *IncrementGetReq, options ...core.ReqOptionFunc) (*IncrementResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "taobao.trades.sold.increment.get"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)

	resp := &IncrementResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
