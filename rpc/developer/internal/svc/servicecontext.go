package svc

import (
	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/model/developer"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/developer/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	DeveloperModel developer.DeveloperModel
	AsynqClient    *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DeveloperModel: developer.NewDeveloperModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Redis, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     c.AsynqRedisConf.Addr,
			Password: c.AsynqRedisConf.Pass,
			DB:       c.AsynqRedisConf.DB,
		}),
	}
}
