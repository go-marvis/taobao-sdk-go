package taobao

import (
	"context"
	"os"
	"testing"

	"github.com/go-marvis/taobao-sdk-go/core"
	"github.com/go-marvis/taobao-sdk-go/service/taobao/trade/fullinfo"
	"github.com/go-marvis/taobao-sdk-go/service/taobao/trades/sold"
	"github.com/stretchr/testify/assert"
)

var (
	api = NewClient(&core.Config{
		AppKey:     os.Getenv("APP_KEY"),
		Secret:     os.Getenv("SECRET"),
		BaseUrl:    "https://eco.taobao.com/router/rest",
		SignMethod: core.MD5,
		Format:     core.JSON,
		Version:    "2.0",
		Session:    os.Getenv("SESSION"),
		Retry:      2,
	})
)

func Test_TaobaoTradeFullInfoGet(t *testing.T) {
	ctx := context.Background()
	resp, err := api.Taobao.Trade.Fullinfo.Get(ctx, fullinfo.NewGetReqBuilder().
		Fields("tid,type,status,payment,orders,rx_audit_status").
		Tid(2590044351089921152).
		Build())
	assert.Nil(t, err)
	assert.NotNil(t, resp.Trade)
}

func Test_TaobaoTradesSoldGet(t *testing.T) {
	ctx := context.Background()
	resp, err := api.Taobao.Trades.Sold.Get(ctx, sold.NewGetReqBuilder().
		PageNo(10).PageSize(100).
		Fields("tid").
		StartCreated("2025-11-10 00:00:00").
		EndCreated("2025-11-11 00:00:00").
		Build())
	assert.Nil(t, err)
	t.Log(len(resp.Trades), resp.TotalResults, resp.HasNext)
	assert.True(t, len(resp.Trades) > 0)
}

func Test_TaobaoTradesSoldGetIncrement(t *testing.T) {
	ctx := context.Background()
	resp, err := api.Taobao.Trades.Sold.IncrementGet(ctx, (*sold.IncrementGetReq)(sold.NewIncrementGetReqBuilder().
		PageNo(0).PageSize(1).
		Fields("tid,type,status").
		StartModified("2025-08-01 10:00:00").
		EndModified("2025-08-01 14:00:00").
		Build()))
	assert.Nil(t, err)
	t.Log(len(resp.Trades))
	assert.True(t, len(resp.Trades) > 0)
}
