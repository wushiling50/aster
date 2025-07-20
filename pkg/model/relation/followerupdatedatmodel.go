package relation

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FollowerUpdatedAtModel = (*customFollowerUpdatedAtModel)(nil)

type (
	// FollowerUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowerUpdatedAtModel.
	FollowerUpdatedAtModel interface {
		followerUpdatedAtModel
		CreateDataId() (int64, error)
	}

	customFollowerUpdatedAtModel struct {
		*defaultFollowerUpdatedAtModel
		sf *utils.Snowflake
	}
)

// NewFollowerUpdatedAtModel returns a model for the database table.
func NewFollowerUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) FollowerUpdatedAtModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customFollowerUpdatedAtModel{
		defaultFollowerUpdatedAtModel: newFollowerUpdatedAtModel(conn, c, opts...),
		sf:                            sf,
	}
}

func (m *customFollowerUpdatedAtModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
