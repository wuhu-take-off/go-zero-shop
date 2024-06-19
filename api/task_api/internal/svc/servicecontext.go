package svc

import (
	"TongChi_shop/api/internal/config"
	"TongChi_shop/rpc/client/taskserve"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	TaskRpc taskserve.TaskServe
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		TaskRpc: taskserve.NewTaskServe(zrpc.MustNewClient(c.ShopRpc)),
	}
}
