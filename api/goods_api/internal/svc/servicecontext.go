package svc

import (
	"TongChi_shop/api/internal/config"
	"TongChi_shop/rpc/client/goodsserve"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	GoodsRpc goodsserve.GoodsServe
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		GoodsRpc: goodsserve.NewGoodsServe(zrpc.MustNewClient(c.ShopRpc)),
	}
}
