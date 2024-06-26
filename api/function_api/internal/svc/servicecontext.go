package svc

import (
	"TongChi_shop/api/internal/config"
	"TongChi_shop/rpc/client/functionserve"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	FunctionRpc functionserve.FunctionServe
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		FunctionRpc: functionserve.NewFunctionServe(zrpc.MustNewClient(c.ShopRpc)),
	}
}
