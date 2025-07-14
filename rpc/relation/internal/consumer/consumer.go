package consumer

import (
	"context"

	"github.com/wushiling50/aster/rpc/relation/internal/config"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svc *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.KafkaQueue.KqCreateRepoConsumerConf, NewCreateRepoConsumer(ctx, svc)),
		kq.MustNewQueue(c.KafkaQueue.KqFollowConsumerConf, NewFollowConsumer(ctx, svc)),
		kq.MustNewQueue(c.KafkaQueue.KqStarConsumerConf, NewStarConsumer(ctx, svc)),
	}
}
