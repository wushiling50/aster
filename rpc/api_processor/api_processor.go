package main

import (
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

	c.MustSetUp()

	svcCtx := svc.NewServiceContext(c)
	taskConsumer := consumer.NewAPITaskConsumer(svcCtx)
	mux := taskConsumer.Register()
	if err := svcCtx.AsynqServer.Run(mux); err != nil {
		logx.Error(err)
		return
	}
}
