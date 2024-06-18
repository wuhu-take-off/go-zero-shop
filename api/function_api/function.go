package function_api

import (
	"TongChi_shop/api/function_api/internal/handler"
	"TongChi_shop/api/function_api/internal/svc"
	"TongChi_shop/api/internal/config"
	"github.com/zeromicro/go-zero/rest"
)

//var configFile = flag.String("f", "etc/function-api.yaml", "the config file")

func Run(server *rest.Server, c config.Config) {
	//flag.Parse()

	//var c config.Config
	//conf.MustLoad(*configFile, &c)

	//server := rest.MustNewServer(c.RestConf)
	//defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	//server.Start()
}
