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

var _ FollowerUpdatedAtModel = (*customFollowerUpdatedAtModel)(nil)

type (
	// FollowerUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowerUpdatedAtModel.
	FollowerUpdatedAtModel interface {
		followerUpdatedAtModel
		FindOneByDeveloperId(ctx context.Context, developerId int64) (*FollowerUpdatedAt, error)
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

func (m *customFollowerUpdatedAtModel) FindOneByDeveloperId(ctx context.Context, developerId int64) (*FollowerUpdatedAt, error) {
	cacheFollowerUpdatedAtDeveloperIdPrefix := fmt.Sprintf("%s%v", "cache:followerUpdatedAt:developerId:", developerId)
	var resp FollowerUpdatedAt
	err := m.QueryRowIndexCtx(ctx, &resp, cacheFollowerUpdatedAtDeveloperIdPrefix, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where developer_id = $1 limit 1", followerUpdatedAtRows, m.table)
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
