package svc

import (
	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/locks"
	"github.com/wushiling50/aster/pkg/model/relation"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/relation/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	CreateRepoModel relation.CreateRepoModel
	FollowModel     relation.FollowModel
	ForkModel       relation.ForkModel
	StarModel       relation.StarModel

	CreatedRepoUpdatedAtModel relation.CreatedRepoUpdatedAtModel
	FollowingUpdatedAtModel   relation.FollowingUpdatedAtModel
	FollowerUpdatedAtModel    relation.FollowerUpdatedAtModel
	ForkUpdatedAtModel        relation.ForkUpdatedAtModel
	StarredRepoUpdatedAtModel relation.StarredRepoUpdatedAtModel

	Locks       *locks.BLock
	AsynqClient *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		CreateRepoModel: relation.NewCreateRepoModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		FollowModel: relation.NewFollowModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		ForkModel: relation.NewForkModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		StarModel: relation.NewStarModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),

		CreatedRepoUpdatedAtModel: relation.NewCreatedRepoUpdatedAtModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		FollowingUpdatedAtModel: relation.NewFollowingUpdatedAtModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		FollowerUpdatedAtModel: relation.NewFollowerUpdatedAtModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		ForkUpdatedAtModel: relation.NewForkUpdatedAtModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		StarredRepoUpdatedAtModel: relation.NewStarredRepoUpdatedAtModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),

		Locks: locks.NewBLock(redis.MustNewRedis(c.Redis.RedisConf), c.Timeout),
		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     c.AsynqRedisConf.Addr,
			Password: c.AsynqRedisConf.Pass,
			DB:       c.AsynqRedisConf.DB,
		}),
	}
}
