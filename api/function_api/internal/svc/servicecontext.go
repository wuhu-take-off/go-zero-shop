package svc

import (
	"TongChi_shop/api/function_api/internal/config"
	"TongChi_shop/api/function_api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config            config.Config
	JwtAuthentication rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		JwtAuthentication: middleware.NewJwtAuthenticationMiddleware().Handle,
	}
}
