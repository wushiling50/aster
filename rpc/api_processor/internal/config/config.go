package config

import (
	"github.com/wushiling50/aster/config"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	Redis          redis.RedisConf
	AsynqRedisConf config.AsynqRedisConf

	Services struct {
		Developer zrpc.RpcClientConf
		Analysis  zrpc.RpcClientConf
	}
}
