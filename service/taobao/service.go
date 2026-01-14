package taobao

import (
	"github.com/go-marvis/taobao-sdk-go/core"
	"github.com/go-marvis/taobao-sdk-go/service/taobao/trade"
	"github.com/go-marvis/taobao-sdk-go/service/taobao/trades"
)

type Service struct {
	Trade  *trade.Service
	Trades *trades.Service
}

func NewService(config *core.Config) *Service {
	return &Service{
		Trade:  trade.NewService(config),
		Trades: trades.NewService(config),
	}
}
