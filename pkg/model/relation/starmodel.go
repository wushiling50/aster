package relation

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StarModel = (*customStarModel)(nil)

type (
	// StarModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStarModel.
	StarModel interface {
		starModel
		CreateDataId() (int64, error)
	}

	customStarModel struct {
		*defaultStarModel
		sf *utils.Snowflake
	}
)

// NewStarModel returns a model for the database table.
func NewStarModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) StarModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customStarModel{
		defaultStarModel: newStarModel(conn, c, opts...),
		sf:               sf,
	}
}

func (m *customStarModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
