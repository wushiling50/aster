package rank

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RankModel struct {
	sf    *utils.Snowflake
	db    *sqlx.SqlConn
	cache *redis.Redis
}

func NewRankModel(conn sqlx.SqlConn, c *redis.Redis, DatancenterId, WorkerId int64) *RankModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}
	return &RankModel{
		sf:    sf,
		db:    &conn,
		cache: c,
	}
}
