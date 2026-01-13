package trades

import (
	"encoding/xml"
	"net/url"

	"github.com/go-marvis/taobao-sdk-go/core"
	"github.com/go-marvis/taobao-sdk-go/service/taobao/trades/model"
)

type SoldGetReqBuilder struct {
	apiReq *core.ApiReq
}

func NewSoldGetReqBuilder() *SoldGetReqBuilder {
	return &SoldGetReqBuilder{
		apiReq: &core.ApiReq{
			QueryParams: url.Values{},
		},
	}
}

func (builder *SoldGetReqBuilder) Fields(fields string) *SoldGetReqBuilder {
	builder.apiReq.QueryParams.Set("fields", fields)
	return builder
}

func (builder *SoldGetReqBuilder) StartCreated(startCreated string) *SoldGetReqBuilder {
	builder.apiReq.QueryParams.Set("start_created", startCreated)
	return builder
}

func (builder *SoldGetReqBuilder) Type(_type string) *SoldGetReqBuilder {
	builder.apiReq.QueryParams.Set("type", _type)
	return builder

}

func (builder *SoldGetReqBuilder) Build() *SoldGetReq {
	return &SoldGetReq{
		apiReq: builder.apiReq,
	}
}

type SoldGetReq struct {
	apiReq *core.ApiReq
}

type SoldGetResp struct {
	*core.ApiResp
	core.CommonResponse
	SoldGetModel
}

type SoldGetModel struct {
	XMLName xml.Name `xml:"trades_sold_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
	Trades []*model.Trade `json:"trades,omitempty" xml:"trades>trade,omitempty"`
	// 搜索到的交易信息总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否存在下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
