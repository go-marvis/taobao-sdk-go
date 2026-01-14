package trade

import (
	"github.com/go-marvis/taobao-sdk-go/core"
	"github.com/go-marvis/taobao-sdk-go/service/taobao/trade/fullinfo"
)

type Service struct {
	Fullinfo *fullinfo.Service
}

func NewService(config *core.Config) *Service {
	return &Service{
		Fullinfo: fullinfo.NewService(config),
	}
}
