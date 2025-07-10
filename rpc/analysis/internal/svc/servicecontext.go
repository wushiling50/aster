package svc

import (
	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/locks"
	"github.com/wushiling50/aster/pkg/model/analysis"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/analysis/internal/config"
	contribution "github.com/wushiling50/aster/rpc/contribution/contributionclient"
	developer "github.com/wushiling50/aster/rpc/developer/developerclient"
	relation "github.com/wushiling50/aster/rpc/relation/relationclient"
	repo "github.com/wushiling50/aster/rpc/repo/repoclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	Locks       *locks.BLock
	AsynqClient *asynq.Client

	NationModel    analysis.NationModel
	LanguagesModel analysis.LanguagesModel
	ScoreModel     analysis.ScoreModel
	SummaryModel   analysis.SummaryModel

	DeveloperRpcClient    developer.DeveloperZrpcClient
	RelationRpcClient     relation.Relation
	ContributionRpcClient contribution.ContributionZrpcClient
	RepoRpcClient         repo.RepoZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Locks:  locks.NewBLock(redis.MustNewRedis(c.Redis.RedisConf), c.Timeout),
		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     c.AsynqRedisConf.Addr,
			Password: c.AsynqRedisConf.Pass,
			DB:       c.AsynqRedisConf.DB,
		}),

		NationModel: analysis.NewNationModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		LanguagesModel: analysis.NewLanguagesModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		ScoreModel: analysis.NewScoreModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		SummaryModel: analysis.NewSummaryModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),

		DeveloperRpcClient:    developer.NewDeveloperZrpcClient(zrpc.MustNewClient(c.Services.Developer)),
		RelationRpcClient:     relation.NewRelation(zrpc.MustNewClient(c.Services.Relation)),
		ContributionRpcClient: contribution.NewContributionZrpcClient(zrpc.MustNewClient(c.Services.Contribution)),
		RepoRpcClient:         repo.NewRepoZrpcClient(zrpc.MustNewClient(c.Services.Repo)),
	}
}
