package taobao

import (
	"github.com/go-marvis/taobao-sdk-go/core"
	"github.com/go-marvis/taobao-sdk-go/service/taobao"
)

type Client struct {
	Taobao *taobao.Service
}

func NewClient(config *core.Config) *Client {
	core.Init(config)

	return &Client{
		Taobao: taobao.NewService(config),
	}
}
