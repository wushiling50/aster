package relation

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FollowModel = (*customFollowModel)(nil)

type (
	// FollowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowModel.
	FollowModel interface {
		followModel
		CreateDataId() (int64, error)
	}

	customFollowModel struct {
		*defaultFollowModel
		sf *utils.Snowflake
	}
)

// NewFollowModel returns a model for the database table.
func NewFollowModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) FollowModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customFollowModel{
		defaultFollowModel: newFollowModel(conn, c, opts...),
		sf:                 sf,
	}
}

func (m *customFollowModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
