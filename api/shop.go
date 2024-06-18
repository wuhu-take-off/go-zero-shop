package main

import (
	"TongChi_shop/api/function_api"
	"TongChi_shop/api/goods_api"
	"TongChi_shop/api/internal/config"
	"TongChi_shop/api/merchant_api"
	"TongChi_shop/api/order_api"
	"TongChi_shop/api/task_api"
	"TongChi_shop/api/user_api"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/shop-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	function_api.Run(server, c)
	goods_api.Run(server, c)
	merchant_api.Run(server, c)
	order_api.Run(server, c)
	task_api.Run(server, c)
	user_api.Run(server, c)
	//ctx := svc.NewServiceContext(c)
	//handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
