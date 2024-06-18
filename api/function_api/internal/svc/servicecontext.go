package svc

import (
	"TongChi_shop/api/function_api/internal/config"
	"TongChi_shop/rpc/client/functionserve"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	ShopRpc functionserve.FunctionServe
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		ShopRpc: functionserve.NewFunctionServe(zrpc.MustNewClient(c.ShopRpc)),
	}
}
