package svc

import (
	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/locks"
	"github.com/wushiling50/aster/pkg/model/developer"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/developer/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	DeveloperModel developer.DeveloperModel
	Locks          *locks.BLock
	AsynqClient    *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DeveloperModel: developer.NewDeveloperModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		Locks: locks.NewBLock(redis.MustNewRedis(c.RedisClient), c.Timeout),
		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     c.AsynqRedisConf.Addr,
			Password: c.AsynqRedisConf.Pass,
			DB:       c.AsynqRedisConf.DB,
		}),
	}
}
