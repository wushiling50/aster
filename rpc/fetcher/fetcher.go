package main

import (
	"flag"

	"github.com/wushiling50/aster/rpc/fetcher/internal/config"
	"github.com/wushiling50/aster/rpc/fetcher/internal/consumer"
	"github.com/wushiling50/aster/rpc/fetcher/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "rpc/fetcher/etc/fetcher.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	c.MustSetUp()

	svcContext := svc.NewServiceContext(c)
	taskConsumer := consumer.NewFetcherTaskConsumer(svcContext)
	mux := taskConsumer.Register()

	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.Error(err)
		return
	}
}
