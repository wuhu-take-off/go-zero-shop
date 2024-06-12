package svc

import (
	"TongChi_shop/rpc/internal/config"
	"TongChi_shop/tool/db_init"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     db_init.ConnMysql(c.Mysql.DataSource),
	}
}
