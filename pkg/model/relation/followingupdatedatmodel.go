package relation

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FollowingUpdatedAtModel = (*customFollowingUpdatedAtModel)(nil)

type (
	// FollowingUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowingUpdatedAtModel.
	FollowingUpdatedAtModel interface {
		followingUpdatedAtModel
		CreateDataId() (int64, error)
	}

	customFollowingUpdatedAtModel struct {
		*defaultFollowingUpdatedAtModel
		sf *utils.Snowflake
	}
)

// NewFollowingUpdatedAtModel returns a model for the database table.
func NewFollowingUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) FollowingUpdatedAtModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customFollowingUpdatedAtModel{
		defaultFollowingUpdatedAtModel: newFollowingUpdatedAtModel(conn, c, opts...),
		sf:                             sf,
	}
}

func (m *customFollowingUpdatedAtModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
