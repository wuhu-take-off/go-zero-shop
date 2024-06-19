package svc

import (
	"TongChi_shop/api/internal/config"
	"TongChi_shop/rpc/client/userserve"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userserve.UserServe
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userserve.NewUserServe(zrpc.MustNewClient(c.ShopRpc)),
	}
}
