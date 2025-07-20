package contribution

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentOfUserUpdatedAtModel = (*customCommentOfUserUpdatedAtModel)(nil)

type (
	// CommentOfUserUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentOfUserUpdatedAtModel.
	CommentOfUserUpdatedAtModel interface {
		commentOfUserUpdatedAtModel
		CreateDataId() (int64, error)
	}

	customCommentOfUserUpdatedAtModel struct {
		*defaultCommentOfUserUpdatedAtModel
		sf *utils.Snowflake
	}
)

// NewCommentOfUserUpdatedAtModel returns a model for the database table.
func NewCommentOfUserUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) CommentOfUserUpdatedAtModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customCommentOfUserUpdatedAtModel{
		defaultCommentOfUserUpdatedAtModel: newCommentOfUserUpdatedAtModel(conn, c, opts...),
		sf:                                 sf,
	}
}

func (m *customCommentOfUserUpdatedAtModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
