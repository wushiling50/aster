package svc

import (
	"time"

	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/constants"
	analysis "github.com/wushiling50/aster/rpc/analysis/analysisclient"
	"github.com/wushiling50/aster/rpc/api_processor/internal/config"
	contribution "github.com/wushiling50/aster/rpc/contribution/contributionclient"
	developer "github.com/wushiling50/aster/rpc/developer/developerclient"
	relation "github.com/wushiling50/aster/rpc/relation/relationclient"
	repo "github.com/wushiling50/aster/rpc/repo/repoclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Redis       *redis.Redis
	AsynqServer *asynq.Server

	DeveloperRpcClient    developer.DeveloperZrpcClient
	RelationRpcClient     relation.Relation
	ContributionRpcClient contribution.ContributionZrpcClient
	RepoRpcClient         repo.RepoZrpcClient
	AnalysisRpcClient     analysis.Analysis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis:  redis.MustNewRedis(c.Redis),
		AsynqServer: asynq.NewServer(
			asynq.RedisClientOpt{
				Addr:     c.AsynqRedisConf.Addr,
				Password: c.AsynqRedisConf.Pass,
				DB:       c.AsynqRedisConf.DB,
			}, asynq.Config{
				Concurrency: constants.APIConcurrency,
				RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
					return constants.APIRetryDelay
				},
				Queues: map[string]int{constants.APITaskQueue: 1},
			}),

		DeveloperRpcClient:    developer.NewDeveloperZrpcClient(zrpc.MustNewClient(c.Services.Developer)),
		RepoRpcClient:         repo.NewRepoZrpcClient(zrpc.MustNewClient(c.Services.Repo)),
		ContributionRpcClient: contribution.NewContributionZrpcClient(zrpc.MustNewClient(c.Services.Contribution)),
		RelationRpcClient:     relation.NewRelation(zrpc.MustNewClient(c.Services.Relation)),
		AnalysisRpcClient:     analysis.NewAnalysis(zrpc.MustNewClient(c.Services.Analysis)),
	}
}
