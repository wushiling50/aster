package config

import (
	"github.com/wushiling50/aster/config"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Snowflake      config.Snowflake
	Mysql          config.MysqlConf
	Redis          redis.RedisConf
	Cache          cache.CacheConf
	AsynqRedisConf config.AsynqRedisConf

	Services struct {
		Developer   zrpc.RpcClientConf
		Analysis    zrpc.RpcClientConf
		IdGenerator zrpc.RpcClientConf
	}
}
