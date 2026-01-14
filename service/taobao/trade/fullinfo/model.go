package fullinfo

import (
	"encoding/xml"
	"fmt"

	"github.com/go-marvis/taobao-sdk-go/core"
	"github.com/go-marvis/taobao-sdk-go/service/taobao/trade/model"
)

type GetReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetReqBuilder() *GetReqBuilder {
	return &GetReqBuilder{core.NewApiReq()}
}

// 需要返回的字段列表
func (builder *GetReqBuilder) Fields(fields string) *GetReqBuilder {
	builder.apiReq.QueryParams.Set("fields", fields)
	return builder
}

// 交易编号
func (builder *GetReqBuilder) Tid(tid int) *GetReqBuilder {
	builder.apiReq.QueryParams.Set("tid", fmt.Sprint(tid))
	return builder
}

// include_oaid
func (builder *GetReqBuilder) IncludeOaid(includeOaid string) *GetReqBuilder {
	builder.apiReq.QueryParams.Set("include_oaid", includeOaid)
	return builder
}

func (builder *GetReqBuilder) Build() *GetReq {
	return &GetReq{apiReq: builder.apiReq}
}

type GetReq struct {
	apiReq *core.ApiReq
}

type GetResp struct {
	*core.ApiResp
	core.CommonResponse
	GetModel
}

type GetModel struct {
	XMLName xml.Name `xml:"trade_fullinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 交易主订单信息
	Trade *model.Trade `json:"trade,omitempty" xml:"trade,omitempty"`
}
