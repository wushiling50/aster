package svc

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/api/internal/config"
	"github.com/wushiling50/aster/pkg/locks"
	model_developer "github.com/wushiling50/aster/pkg/model/developer"
	"github.com/wushiling50/aster/pkg/model/rank"
	"github.com/wushiling50/aster/pkg/utils"
	analysis "github.com/wushiling50/aster/rpc/analysis/analysisclient"
	developer "github.com/wushiling50/aster/rpc/developer/developerclient"
	"github.com/wushiling50/aster/rpc/id_generator/idgenerator"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	AsynqClient    *asynq.Client
	AsynqInspector *asynq.Inspector

	RankModel      *rank.RankModel
	DeveloperModel model_developer.DeveloperModel

	Locks *locks.BLock

	DeveloperRpcClient   developer.DeveloperZrpcClient
	AnalysisRpcClient    analysis.Analysis
	IdGeneratorRpcClient idgenerator.IdGenerator
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
		RankModel: rank.NewRankModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			redis.MustNewRedis(c.Redis), c.Snowflake.DatancenterId, c.Snowflake.WorkerId),

		DeveloperModel: model_developer.NewDeveloperModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),

		Locks: locks.NewBLock(redis.MustNewRedis(c.Redis), c.Timeout),

		DeveloperRpcClient:   developer.NewDeveloperZrpcClient(zrpc.MustNewClient(c.Services.Developer)),
		AnalysisRpcClient:    analysis.NewAnalysis(zrpc.MustNewClient(c.Services.Analysis)),
		IdGeneratorRpcClient: idgenerator.NewIdGenerator(zrpc.MustNewClient(c.Services.IdGenerator)),
	}
}
