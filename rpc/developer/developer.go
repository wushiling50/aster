package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/wushiling50/aster/gen/developer"
	"github.com/wushiling50/aster/rpc/developer/internal/config"
	"github.com/wushiling50/aster/rpc/developer/internal/consumer"
	"github.com/wushiling50/aster/rpc/developer/internal/server"
	"github.com/wushiling50/aster/rpc/developer/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/developer.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		developer.RegisterDeveloperServer(grpcServer, server.NewDeveloperServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	serviceGroup.Add(s)

	for _, s := range consumer.Consumers(c, context.Background(), ctx) {
		serviceGroup.Add(s)
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	serviceGroup.Start()
}
