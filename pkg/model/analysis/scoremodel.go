package analysis

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ScoreModel = (*customScoreModel)(nil)

type (
	// ScoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customScoreModel.
	ScoreModel interface {
		scoreModel
		CreateDataId() (int64, error)
	}

	customScoreModel struct {
		*defaultScoreModel
		sf *utils.Snowflake
	}
)

// NewScoreModel returns a model for the database table.
func NewScoreModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) ScoreModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customScoreModel{
		defaultScoreModel: newScoreModel(conn, c, opts...),
		sf:                sf,
	}
}

func (m *customScoreModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
