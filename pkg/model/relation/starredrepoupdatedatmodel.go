package relation

import (
	"context"
	"fmt"

	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StarredRepoUpdatedAtModel = (*customStarredRepoUpdatedAtModel)(nil)

type (
	// StarredRepoUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStarredRepoUpdatedAtModel.
	StarredRepoUpdatedAtModel interface {
		starredRepoUpdatedAtModel
		FindOneByDeveloperId(ctx context.Context, developerId int64) (*StarredRepoUpdatedAt, error)
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

func (m *customStarredRepoUpdatedAtModel) FindOneByDeveloperId(ctx context.Context, developerId int64) (*StarredRepoUpdatedAt, error) {
	cacheStarredRepoUpdatedAtDeveloperIdKey := fmt.Sprintf("%s%v", "cache:starredRepoUpdatedAt:developerId:", developerId)
	var resp StarredRepoUpdatedAt
	err := m.QueryRowIndexCtx(ctx, &resp, cacheStarredRepoUpdatedAtDeveloperIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where developer_id = ? limit 1", starredRepoUpdatedAtRows, m.table)
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

func (m *customStarredRepoUpdatedAtModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
