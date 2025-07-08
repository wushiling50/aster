package config

import (
	"github.com/wushiling50/aster/config"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Snowflake      config.Snowflake
	Mysql          config.MysqlConf
	Cache          cache.CacheConf
	AsynqRedisConf config.AsynqRedisConf

	KqContributionConsumerConf kq.KqConf
}
