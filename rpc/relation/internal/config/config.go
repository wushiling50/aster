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
	Redis          cache.CacheConf
	AsynqRedisConf config.AsynqRedisConf

	KafkaQueue struct {
		KqCreateRepoConsumerConf kq.KqConf
		KqFollowConsumerConf     kq.KqConf
		KqStarConsumerConf       kq.KqConf
		KqForkConsumerConf       kq.KqConf
	}
}
