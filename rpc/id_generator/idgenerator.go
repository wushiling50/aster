package main

import (
	"flag"
	"fmt"

	"github.com/wushiling50/aster/gen/id_generator"
	"github.com/wushiling50/aster/rpc/id_generator/internal/config"
	"github.com/wushiling50/aster/rpc/id_generator/internal/server"
	"github.com/wushiling50/aster/rpc/id_generator/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/idgenerator.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		id_generator.RegisterIdGeneratorServer(grpcServer, server.NewIdGeneratorServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
