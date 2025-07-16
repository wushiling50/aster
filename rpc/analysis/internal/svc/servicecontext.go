package svc

import (
	"github.com/hibiken/asynq"
	"github.com/sashabaranov/go-openai"
	"github.com/wushiling50/aster/pkg/llm"
	"github.com/wushiling50/aster/pkg/locks"
	"github.com/wushiling50/aster/pkg/model/analysis"
	model_contribution "github.com/wushiling50/aster/pkg/model/contribution"
	model_developer "github.com/wushiling50/aster/pkg/model/developer"
	model_relation "github.com/wushiling50/aster/pkg/model/relation"
	model_repo "github.com/wushiling50/aster/pkg/model/repo"
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

	Locks          *locks.BLock
	RedisClient    *redis.Redis
	AsynqClient    *asynq.Client
	DeepSeekClient *openai.Client

	NationModel    analysis.NationModel
	LanguagesModel analysis.LanguagesModel
	ScoreModel     analysis.ScoreModel
	SummaryModel   analysis.SummaryModel

	CreatedRepoUpdatedAtModel   model_relation.CreatedRepoUpdatedAtModel
	FollowerUpdatedAtModel      model_relation.FollowerUpdatedAtModel
	StarredRepoUpdatedAtModel   model_relation.StarredRepoUpdatedAtModel
	DeveloperModel              model_developer.DeveloperModel
	CommentOfUserUpdatedAtModel model_contribution.CommentOfUserUpdatedAtModel
	IssuePROfUserUpdatedAtModel model_contribution.IssuePrOfUserUpdatedAtModel
	ReviewOfUserUpdatedAtModel  model_contribution.ReviewOfUserUpdatedAtModel
	RepoModel                   model_repo.RepoModel

	DeveloperRpcClient    developer.DeveloperZrpcClient
	RelationRpcClient     relation.Relation
	ContributionRpcClient contribution.ContributionZrpcClient
	RepoRpcClient         repo.RepoZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Locks:       locks.NewBLock(redis.MustNewRedis(c.RedisClient), c.Timeout),
		RedisClient: redis.MustNewRedis(c.RedisClient),
		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     c.AsynqRedisConf.Addr,
			Password: c.AsynqRedisConf.Pass,
			DB:       c.AsynqRedisConf.DB,
		}),
		DeepSeekClient: llm.NewDeepSeekClient(c.DeepSeek),

		NationModel: analysis.NewNationModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		LanguagesModel: analysis.NewLanguagesModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		ScoreModel: analysis.NewScoreModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		SummaryModel: analysis.NewSummaryModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),

		CreatedRepoUpdatedAtModel: model_relation.NewCreatedRepoUpdatedAtModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		FollowerUpdatedAtModel: model_relation.NewFollowerUpdatedAtModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		StarredRepoUpdatedAtModel: model_relation.NewStarredRepoUpdatedAtModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		DeveloperModel: model_developer.NewDeveloperModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		CommentOfUserUpdatedAtModel: model_contribution.NewCommentOfUserUpdatedAtModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		IssuePROfUserUpdatedAtModel: model_contribution.NewIssuePrOfUserUpdatedAtModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		ReviewOfUserUpdatedAtModel: model_contribution.NewReviewOfUserUpdatedAtModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),
		RepoModel: model_repo.NewRepoModel(sqlx.NewMysql(utils.GetMysqlDSN(c.Mysql)),
			c.Cache, c.Snowflake.DatancenterId, c.Snowflake.WorkerId),

		DeveloperRpcClient:    developer.NewDeveloperZrpcClient(zrpc.MustNewClient(c.Services.Developer)),
		RelationRpcClient:     relation.NewRelation(zrpc.MustNewClient(c.Services.Relation)),
		ContributionRpcClient: contribution.NewContributionZrpcClient(zrpc.MustNewClient(c.Services.Contribution)),
		RepoRpcClient:         repo.NewRepoZrpcClient(zrpc.MustNewClient(c.Services.Repo)),
	}
}
