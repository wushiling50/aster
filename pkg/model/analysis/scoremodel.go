package analysis

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ScoreModel = (*customScoreModel)(nil)

type (
	// ScoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customScoreModel.
	ScoreModel interface {
		scoreModel
	}

	customScoreModel struct {
		*defaultScoreModel
	}
)

// NewScoreModel returns a model for the database table.
func NewScoreModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ScoreModel {
	return &customScoreModel{
		defaultScoreModel: newScoreModel(conn, c, opts...),
	}
}
