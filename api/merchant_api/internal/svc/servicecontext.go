package svc

import (
	"TongChi_shop/api/internal/config"
	"TongChi_shop/rpc/client/merchantserve"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	MerchantRpc merchantserve.MerchantServe
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		MerchantRpc: merchantserve.NewMerchantServe(zrpc.MustNewClient(c.ShopRpc)),
	}
}
