package relation

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StarredRepoUpdatedAtModel = (*customStarredRepoUpdatedAtModel)(nil)

type (
	// StarredRepoUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStarredRepoUpdatedAtModel.
	StarredRepoUpdatedAtModel interface {
		starredRepoUpdatedAtModel
		CreateDataId() (int64, error)
	}

	customStarredRepoUpdatedAtModel struct {
		*defaultStarredRepoUpdatedAtModel
		sf *utils.Snowflake
	}
)

// NewStarredRepoUpdatedAtModel returns a model for the database table.
func NewStarredRepoUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) StarredRepoUpdatedAtModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customStarredRepoUpdatedAtModel{
		defaultStarredRepoUpdatedAtModel: newStarredRepoUpdatedAtModel(conn, c, opts...),
		sf:                               sf,
	}
}

func (m *customStarredRepoUpdatedAtModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
