package contribution

import (
	"context"
	"fmt"

	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentOfUserUpdatedAtModel = (*customCommentOfUserUpdatedAtModel)(nil)

type (
	// CommentOfUserUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentOfUserUpdatedAtModel.
	CommentOfUserUpdatedAtModel interface {
		commentOfUserUpdatedAtModel
		FindOneByDeveloperId(ctx context.Context, developerId int64) (*CommentOfUserUpdatedAt, error)
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

func (m *customCommentOfUserUpdatedAtModel) FindOneByDeveloperId(ctx context.Context, developerId int64) (*CommentOfUserUpdatedAt, error) {
	cacheCommentOfUserUpdatedAtDeveloperIdKey := fmt.Sprintf("%s%v", "cache:commentOfUserUpdatedAt:developerId:", developerId)
	var resp CommentOfUserUpdatedAt
	err := m.QueryRowIndexCtx(ctx, &resp, cacheCommentOfUserUpdatedAtDeveloperIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where developer_id = ? limit 1", commentOfUserUpdatedAtRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, developerId); err != nil {
			return nil, err
		}
		return resp.DataId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customCommentOfUserUpdatedAtModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
