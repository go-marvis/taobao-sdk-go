package sold

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

// 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss
func (builder *GetReqBuilder) StartCreated(startCreated string) *GetReqBuilder {
	builder.apiReq.QueryParams.Set("start_created", startCreated)
	return builder
}

// 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss
func (builder *GetReqBuilder) EndCreated(endCreated string) *GetReqBuilder {
	builder.apiReq.QueryParams.Set("end_created", endCreated)
	return builder
}

func (builder *GetReqBuilder) Status(status string) *GetReqBuilder {
	builder.apiReq.QueryParams.Set("status", status)
	return builder
}

func (builder *GetReqBuilder) BuyerNick(buyerNick string) *GetReqBuilder {
	builder.apiReq.QueryParams.Set("buyer_nick", buyerNick)
	return builder
}

func (builder *GetReqBuilder) Type(_type string) *GetReqBuilder {
	builder.apiReq.QueryParams.Set("type", _type)
	return builder
}

func (builder *GetReqBuilder) ExtType(extType string) *GetReqBuilder {
	builder.apiReq.QueryParams.Set("ext_type", extType)
	return builder
}

func (builder *GetReqBuilder) RateStatus(rateStatus string) *GetReqBuilder {
	builder.apiReq.QueryParams.Set("rate_status", rateStatus)
	return builder
}

func (builder *GetReqBuilder) Tag(tag string) *GetReqBuilder {
	builder.apiReq.QueryParams.Set("tag", tag)
	return builder
}

func (builder *GetReqBuilder) PageNo(pageNo int) *GetReqBuilder {
	builder.apiReq.QueryParams.Set("page_no", fmt.Sprint(pageNo))
	return builder
}

func (builder *GetReqBuilder) PageSize(pageSize int) *GetReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}

func (builder *GetReqBuilder) UseHasNext(useHasNext bool) *GetReqBuilder {
	builder.apiReq.QueryParams.Set("use_has_next", fmt.Sprint(useHasNext))
	return builder
}

func (builder *GetReqBuilder) BuyerOpenId(buyerOpenId string) *GetReqBuilder {
	builder.apiReq.QueryParams.Set("buyer_open_id", buyerOpenId)
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

type IncrementGetReqBuilder struct {
	apiReq *core.ApiReq
}

func NewIncrementGetReqBuilder() *IncrementGetReqBuilder {
	return &IncrementGetReqBuilder{core.NewApiReq()}
}

// 需要返回的字段列表
func (builder *IncrementGetReqBuilder) Fields(fields string) *IncrementGetReqBuilder {
	builder.apiReq.QueryParams.Set("fields", fields)
	return builder
}

// 查询修改开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss
func (builder *IncrementGetReqBuilder) StartModified(startModified string) *IncrementGetReqBuilder {
	builder.apiReq.QueryParams.Set("start_modified", startModified)
	return builder
}

// 查询修改结束时间，必须大于修改开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。
func (builder *IncrementGetReqBuilder) EndModified(endModified string) *IncrementGetReqBuilder {
	builder.apiReq.QueryParams.Set("end_modified", endModified)
	return builder
}

func (builder *IncrementGetReqBuilder) Status(status string) *IncrementGetReqBuilder {
	builder.apiReq.QueryParams.Set("status", status)
	return builder
}

func (builder *IncrementGetReqBuilder) Type(_type string) *IncrementGetReqBuilder {
	builder.apiReq.QueryParams.Set("type", _type)
	return builder
}

func (builder *IncrementGetReqBuilder) BuyerNick(buyerNick string) *IncrementGetReqBuilder {
	builder.apiReq.QueryParams.Set("buyer_nick", buyerNick)
	return builder
}

func (builder *IncrementGetReqBuilder) ExtType(extType string) *IncrementGetReqBuilder {
	builder.apiReq.QueryParams.Set("ext_type", extType)
	return builder
}

func (builder *IncrementGetReqBuilder) Tag(tag string) *IncrementGetReqBuilder {
	builder.apiReq.QueryParams.Set("tag", tag)
	return builder
}

func (builder *IncrementGetReqBuilder) PageNo(pageNo int) *IncrementGetReqBuilder {
	builder.apiReq.QueryParams.Set("page_no", fmt.Sprint(pageNo))
	return builder
}

func (builder *IncrementGetReqBuilder) PageSize(pageSize int) *IncrementGetReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}

func (builder *IncrementGetReqBuilder) UseHasNext(useHasNext bool) *IncrementGetReqBuilder {
	builder.apiReq.QueryParams.Set("use_has_next", fmt.Sprint(useHasNext))
	return builder
}

func (builder *IncrementGetReqBuilder) BuyerOpenId(buyerOpenId string) *IncrementGetReqBuilder {
	builder.apiReq.QueryParams.Set("buyer_open_id", buyerOpenId)
	return builder
}

func (builder *IncrementGetReqBuilder) Build() *IncrementGetReq {
	return &IncrementGetReq{
		apiReq: builder.apiReq,
	}
}

type IncrementGetReq struct {
	apiReq *core.ApiReq
}

type IncrementResp struct {
	*core.ApiResp
	core.CommonResponse
	IncrementGetModel
}

type IncrementGetModel struct {
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
