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

var _ ForkUpdatedAtModel = (*customForkUpdatedAtModel)(nil)

type (
	// ForkUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customForkUpdatedAtModel.
	ForkUpdatedAtModel interface {
		forkUpdatedAtModel
		FindOneByRepoId(ctx context.Context, repoId int64) (*ForkUpdatedAt, error)
		CreateDataId() (int64, error)
	}

	customForkUpdatedAtModel struct {
		*defaultForkUpdatedAtModel
		sf *utils.Snowflake
	}
)

// NewForkUpdatedAtModel returns a model for the database table.
func NewForkUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) ForkUpdatedAtModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customForkUpdatedAtModel{
		defaultForkUpdatedAtModel: newForkUpdatedAtModel(conn, c, opts...),
		sf:                        sf,
	}
}

func (m *customForkUpdatedAtModel) FindOneByRepoId(ctx context.Context, repoId int64) (*ForkUpdatedAt, error) {
	cacheForkUpdatedAtRepoIdPrefix := fmt.Sprintf("%s%v", "cache:forkUpdatedAt:repoId:", repoId)
	var resp ForkUpdatedAt
	err := m.QueryRowIndexCtx(ctx, &resp, cacheForkUpdatedAtRepoIdPrefix, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where repo_id = $1 limit 1", forkUpdatedAtRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, repoId); err != nil {
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

func (m *customForkUpdatedAtModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
