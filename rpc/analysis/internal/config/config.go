package config

import (
	"github.com/wushiling50/aster/config"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Snowflake config.Snowflake
	Mysql     config.MysqlConf
	Cache     cache.CacheConf
	DeepSeek  config.DeepSeekModel

	Services struct {
		Developer    zrpc.RpcClientConf
		Relation     zrpc.RpcClientConf
		Contribution zrpc.RpcClientConf
		Repo         zrpc.RpcClientConf
	}
}
