package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	CacheRedis struct {
		Addr     string
		Password string
	}
	Mysql struct {
		DataSource string
	}
	AES struct {
		Key string
	}
}
