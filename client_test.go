package taobao

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/go-marvis/taobao-sdk-go/core"
	"github.com/go-marvis/taobao-sdk-go/service/taobao/trades"
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

func Test_TaobaoTradesSoldGet(t *testing.T) {
	ctx := context.Background()
	resp, err := api.Taobao.Trades.Sold.Get(ctx, trades.NewSoldGetReqBuilder().
		Fields("tid,type,status,payment,orders,rx_audit_status").
		Build())
	assert.Nil(t, err)
	for _, i := range resp.Trades {
		data, _ := json.MarshalIndent(i, "", "  ")
		fmt.Println(string(data))
	}
}
