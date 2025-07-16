package svc

import (
	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/locks"
	"github.com/wushiling50/aster/pkg/model/contribution"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/contribution/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	ContributionModel contribution.ContributionModel

	IssuePrOfUserUpdatedAtModel contribution.IssuePrOfUserUpdatedAtModel
	ReviewOfUserUpdatedAtModel  contribution.ReviewOfUserUpdatedAtModel
	CommentOfUserUpdatedAtModel contribution.CommentOfUserUpdatedAtModel

	Locks       *locks.BLock
	AsynqClient *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		ContributionModel: contribution.NewContributionModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		IssuePrOfUserUpdatedAtModel: contribution.NewIssuePrOfUserUpdatedAtModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		ReviewOfUserUpdatedAtModel: contribution.NewReviewOfUserUpdatedAtModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		CommentOfUserUpdatedAtModel: contribution.NewCommentOfUserUpdatedAtModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),

		Locks: locks.NewBLock(redis.MustNewRedis(c.RedisClient), c.Timeout),
		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     c.AsynqRedisConf.Addr,
			Password: c.AsynqRedisConf.Pass,
			DB:       c.AsynqRedisConf.DB,
		}),
	}
}
