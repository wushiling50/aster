package config

import (
	"github.com/wushiling50/aster/config"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Snowflake      config.Snowflake
	Mysql          config.MysqlConf
	RedisClient    redis.RedisConf
	Cache          cache.CacheConf
	AsynqRedisConf config.AsynqRedisConf
	DeepSeek       config.DeepSeekModel

	Services struct {
		Developer    zrpc.RpcClientConf
		Relation     zrpc.RpcClientConf
		Contribution zrpc.RpcClientConf
		Repo         zrpc.RpcClientConf
	}
}
