package config

import (
	"github.com/wushiling50/aster/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql          config.MysqlConf
	Redis          redis.RedisConf
	AsynqRedisConf config.AsynqRedisConf

	Services struct {
		Developer    zrpc.RpcClientConf
		Analysis     zrpc.RpcClientConf
		Relation     zrpc.RpcClientConf
		Repo         zrpc.RpcClientConf
		Contribution zrpc.RpcClientConf
		IdGenerator  zrpc.RpcClientConf
	}
}
