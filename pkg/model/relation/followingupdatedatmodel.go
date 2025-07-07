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

var _ FollowingUpdatedAtModel = (*customFollowingUpdatedAtModel)(nil)

type (
	// FollowingUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowingUpdatedAtModel.
	FollowingUpdatedAtModel interface {
		followingUpdatedAtModel
		FindOneByDeveloperId(ctx context.Context, developerId int64) (*FollowingUpdatedAt, error)
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

func (m *customFollowingUpdatedAtModel) FindOneByDeveloperId(ctx context.Context, developerId int64) (*FollowingUpdatedAt, error) {
	cacheFollowingUpdatedAtDeveloperIdKey := fmt.Sprintf("%s%v", "cache:followingUpdatedAt:developerId:", developerId)
	var resp FollowingUpdatedAt
	err := m.QueryRowIndexCtx(ctx, &resp, cacheFollowingUpdatedAtDeveloperIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where developer_id = $1 limit 1", followingUpdatedAtRows, m.table)
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

func (m *customFollowingUpdatedAtModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
