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
		Developer    zrpc.RpcClientConf
		Relation     zrpc.RpcClientConf
		Repo         zrpc.RpcClientConf
		Contribution zrpc.RpcClientConf
	}

	KqDeveloperPusherConf    config.KqPusherConf
	KqContributionPusherConf config.KqPusherConf
	KqCreateRepoPusherConf   config.KqPusherConf
	KqForkPusherConf         config.KqPusherConf
	KqStarPusherConf         config.KqPusherConf
	KqFollowPusherConf       config.KqPusherConf
	KqRepoPusherConf         config.KqPusherConf

	KqDeveloperUpdateCompletePusherConf    config.KqPusherConf
	KqRepoUpdateCompletePusherConf         config.KqPusherConf
	KqContributionUpdateCompletePusherConf config.KqPusherConf
	KqRelationUpdateCompletePusherConf     config.KqPusherConf
}
