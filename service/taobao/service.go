package taobao

import (
	"github.com/go-marvis/taobao-sdk-go/core"
	"github.com/go-marvis/taobao-sdk-go/service/taobao/trades"
)

type Service struct {
	Trades *trades.Service
}

func NewService(config *core.Config) *Service {
	return &Service{
		Trades: trades.NewService(config),
	}
}
