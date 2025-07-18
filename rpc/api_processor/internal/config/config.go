package config

import (
	"github.com/wushiling50/aster/config"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf

	AsynqRedisConf config.AsynqRedisConf

	Services struct {
		Analysis zrpc.RpcClientConf
	}
}
