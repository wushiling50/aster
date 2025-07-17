package config

import (
	"github.com/wushiling50/aster/config"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	AsynqRedisConf config.AsynqRedisConf

	KafkaQueue struct {
		KqDeveloperPusherConf    config.KqPusherConf
		KqContributionPusherConf config.KqPusherConf
		KqCreateRepoPusherConf   config.KqPusherConf
		KqStarPusherConf         config.KqPusherConf
		KqFollowPusherConf       config.KqPusherConf
		KqRepoPusherConf         config.KqPusherConf
	}

	Services struct {
		Developer    zrpc.RpcClientConf
		Relation     zrpc.RpcClientConf
		Repo         zrpc.RpcClientConf
		Contribution zrpc.RpcClientConf
	}
}
