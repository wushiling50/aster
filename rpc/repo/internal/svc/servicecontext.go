package svc

import (
	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/model/repo"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/repo/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RepoModel   repo.RepoModel
	Redis       *redis.Redis
	AsynqClient *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RepoModel: repo.NewRepoModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		Redis: redis.MustNewRedis(c.Redis.RedisConf),
		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     c.AsynqRedisConf.Addr,
			Password: c.AsynqRedisConf.Pass,
			DB:       c.AsynqRedisConf.DB,
		}),
	}
}
