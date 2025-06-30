package svc

import (
	"time"

	contribution "github.com/wushiling50/aster/rpc/contribution/contributionclient"
	developer "github.com/wushiling50/aster/rpc/developer/developerclient"
	relation "github.com/wushiling50/aster/rpc/relation/relationclient"
	repo "github.com/wushiling50/aster/rpc/repo/repoclient"

	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/rpc/fetcher/internal/config"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Redis       *redis.Redis
	AsynqServer *asynq.Server

	KqDeveloperPusher    *kq.Pusher
	KqContributionPusher *kq.Pusher
	KqCreateRepoPusher   *kq.Pusher
	KqForkPusher         *kq.Pusher
	KqStarPusher         *kq.Pusher
	KqFollowPusher       *kq.Pusher
	KqRepoPusher         *kq.Pusher

	DeveloperRpcClient    developer.DeveloperZrpcClient
	RelationRpcClient     relation.Relation
	ContributionRpcClient contribution.ContributionZrpcClient
	RepoRpcClient         repo.RepoZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		Redis:                 redis.MustNewRedis(c.Redis),
		DeveloperRpcClient:    developer.NewDeveloperZrpcClient(zrpc.MustNewClient(c.Services.Developer)),
		RelationRpcClient:     relation.NewRelation(zrpc.MustNewClient(c.Services.Relation)),
		ContributionRpcClient: contribution.NewContributionZrpcClient(zrpc.MustNewClient(c.Services.Contribution)),
		RepoRpcClient:         repo.NewRepoZrpcClient(zrpc.MustNewClient(c.Services.Repo)),

		KqDeveloperPusher: kq.NewPusher(
			c.KafkaQueue.KqDeveloperPusherConf.Brokers,
			c.KafkaQueue.KqDeveloperPusherConf.Topic,
			kq.WithAllowAutoTopicCreation(),
			kq.WithSyncPush()),
		KqContributionPusher: kq.NewPusher(
			c.KafkaQueue.KqContributionPusherConf.Brokers,
			c.KafkaQueue.KqContributionPusherConf.Topic,
			kq.WithAllowAutoTopicCreation(),
			kq.WithSyncPush()),
		KqCreateRepoPusher: kq.NewPusher(
			c.KafkaQueue.KqCreateRepoPusherConf.Brokers,
			c.KafkaQueue.KqCreateRepoPusherConf.Topic,
			kq.WithAllowAutoTopicCreation(),
			kq.WithSyncPush()),
		KqForkPusher: kq.NewPusher(
			c.KafkaQueue.KqForkPusherConf.Brokers,
			c.KafkaQueue.KqForkPusherConf.Topic,
			kq.WithAllowAutoTopicCreation(),
			kq.WithSyncPush()),
		KqStarPusher: kq.NewPusher(
			c.KafkaQueue.KqStarPusherConf.Brokers,
			c.KafkaQueue.KqStarPusherConf.Topic,
			kq.WithAllowAutoTopicCreation(),
			kq.WithSyncPush()),
		KqFollowPusher: kq.NewPusher(
			c.KafkaQueue.KqFollowPusherConf.Brokers,
			c.KafkaQueue.KqFollowPusherConf.Topic,
			kq.WithAllowAutoTopicCreation(),
			kq.WithSyncPush()),
		KqRepoPusher: kq.NewPusher(
			c.KafkaQueue.KqRepoPusherConf.Brokers,
			c.KafkaQueue.KqRepoPusherConf.Topic,
			kq.WithAllowAutoTopicCreation(),
			kq.WithSyncPush()),

		AsynqServer: asynq.NewServer(
			asynq.RedisClientOpt{
				Addr:     c.AsynqRedisConf.Addr,
				Password: c.AsynqRedisConf.Pass,
				DB:       c.AsynqRedisConf.DB,
			}, asynq.Config{
				Concurrency: constants.FetchConcurrency,
				RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
					return constants.FetchRetryDelay
				},
				Queues: map[string]int{constants.FetcherTaskQueue: 1},
			}),
	}
}
