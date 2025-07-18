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

var _ CreatedRepoUpdatedAtModel = (*customCreatedRepoUpdatedAtModel)(nil)

type (
	// CreatedRepoUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCreatedRepoUpdatedAtModel.
	CreatedRepoUpdatedAtModel interface {
		createdRepoUpdatedAtModel
		FindOneByDeveloperId(ctx context.Context, developerId int64) (*CreatedRepoUpdatedAt, error)
		CreateDataId() (int64, error)
	}

	customCreatedRepoUpdatedAtModel struct {
		*defaultCreatedRepoUpdatedAtModel
		sf *utils.Snowflake
	}
)

// NewCreatedRepoUpdatedAtModel returns a model for the database table.
func NewCreatedRepoUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) CreatedRepoUpdatedAtModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customCreatedRepoUpdatedAtModel{
		defaultCreatedRepoUpdatedAtModel: newCreatedRepoUpdatedAtModel(conn, c, opts...),
		sf:                               sf,
	}
}

func (m *customCreatedRepoUpdatedAtModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}

func (m *customCreatedRepoUpdatedAtModel) FindOneByDeveloperId(ctx context.Context, developerId int64) (*CreatedRepoUpdatedAt, error) {
	cacheCreatedRepoUpdatedAtDeveloperIdKey := fmt.Sprintf("%s%v", "cache:createdRepoUpdatedAt:developerId:", developerId)
	var resp CreatedRepoUpdatedAt
	err := m.QueryRowIndexCtx(ctx, &resp, cacheCreatedRepoUpdatedAtDeveloperIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where developer_id = ? limit 1", createdRepoUpdatedAtRows, m.table)
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
