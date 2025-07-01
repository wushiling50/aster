package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/wushiling50/aster/gen/repo"
	"github.com/wushiling50/aster/rpc/repo/internal/config"
	"github.com/wushiling50/aster/rpc/repo/internal/consumer"
	"github.com/wushiling50/aster/rpc/repo/internal/server"
	"github.com/wushiling50/aster/rpc/repo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "rpc/repo/etc/repo.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		repo.RegisterRepoServer(grpcServer, server.NewRepoServer(ctx))

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
