package rank

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RankModel struct {
	db    *sqlx.SqlConn
	cache *redis.Redis
}

func NewRankModel(conn sqlx.SqlConn, c *redis.Redis) *RankModel {
	return &RankModel{
		db:    &conn,
		cache: c,
	}
}
