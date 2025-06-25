package svc

import (
	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/api/internal/config"
	analysis "github.com/wushiling50/aster/rpc/analysis/analysisclient"
	contribution "github.com/wushiling50/aster/rpc/contribution/contributionclient"
	developer "github.com/wushiling50/aster/rpc/developer/developerclient"
	"github.com/wushiling50/aster/rpc/id_generator/idgenerator"
	relation "github.com/wushiling50/aster/rpc/relation/relationclient"
	repo "github.com/wushiling50/aster/rpc/repo/repoclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	AsynqClient    *asynq.Client
	AsynqInspector *asynq.Inspector
	RedisClient    *redis.Redis

	DeveloperRpcClient    developer.DeveloperZrpcClient
	RepoRpcClient         repo.RepoZrpcClient
	ContributionRpcClient contribution.ContributionZrpcClient
	RelationRpcClient     relation.Relation
	AnalysisRpcClient     analysis.Analysis
	IdGeneratorRpcClient  idgenerator.IdGenerator
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		AsynqClient: asynq.NewClient(&asynq.RedisClientOpt{
			Addr:     c.AsynqRedisConf.Addr,
			Password: c.AsynqRedisConf.Pass,
			DB:       c.AsynqRedisConf.DB,
		}),
		AsynqInspector: asynq.NewInspector(&asynq.RedisClientOpt{
			Addr:     c.AsynqRedisConf.Addr,
			Password: c.AsynqRedisConf.Pass,
			DB:       c.AsynqRedisConf.DB,
		}),
		RedisClient: redis.MustNewRedis(c.Redis),

		DeveloperRpcClient:    developer.NewDeveloperZrpcClient(zrpc.MustNewClient(c.Services.Developer)),
		RepoRpcClient:         repo.NewRepoZrpcClient(zrpc.MustNewClient(c.Services.Repo)),
		ContributionRpcClient: contribution.NewContributionZrpcClient(zrpc.MustNewClient(c.Services.Contribution)),
		RelationRpcClient:     relation.NewRelation(zrpc.MustNewClient(c.Services.Relation)),
		AnalysisRpcClient:     analysis.NewAnalysis(zrpc.MustNewClient(c.Services.Analysis)),
		IdGeneratorRpcClient:  idgenerator.NewIdGenerator(zrpc.MustNewClient(c.Services.IdGenerator)),
	}
}
