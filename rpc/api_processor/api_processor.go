package main

import (
	"context"
	"flag"

	"github.com/wushiling50/aster/rpc/api_processor/internal/config"
	"github.com/wushiling50/aster/rpc/api_processor/internal/consumer"
	"github.com/wushiling50/aster/rpc/api_processor/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "rpc/api_processor/etc/api_processor.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := context.Background()
	c.MustSetUp()

	svcContext := svc.NewServiceContext(c)
	taskConsumer := consumer.NewAPITaskConsumer(ctx, svcContext)
	mux := taskConsumer.Register()
	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.Error(err)
		return
	}
}
