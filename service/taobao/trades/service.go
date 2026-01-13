package trades

import "github.com/go-marvis/taobao-sdk-go/core"

type Service struct {
	Sold *sold
}

func NewService(config *core.Config) *Service {
	return &Service{
		Sold: &sold{config: config},
	}
}
