package svc

import (
	"TongChi_shop/api/internal/config"
	"TongChi_shop/rpc/client/orderserve"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	OrderRpc orderserve.OrderServe
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		OrderRpc: orderserve.NewOrderServe(zrpc.MustNewClient(c.ShopRpc)),
	}
}
