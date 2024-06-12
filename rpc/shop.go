package main

import (
	"flag"
	"fmt"

	"TongChi_shop/rpc/internal/config"
	functionserveServer "TongChi_shop/rpc/internal/server/functionserve"
	goodsserveServer "TongChi_shop/rpc/internal/server/goodsserve"
	userserveServer "TongChi_shop/rpc/internal/server/userserve"
	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/shop.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		shop.RegisterUserServeServer(grpcServer, userserveServer.NewUserServeServer(ctx))
		shop.RegisterGoodsServeServer(grpcServer, goodsserveServer.NewGoodsServeServer(ctx))
		shop.RegisterFunctionServeServer(grpcServer, functionserveServer.NewFunctionServeServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
