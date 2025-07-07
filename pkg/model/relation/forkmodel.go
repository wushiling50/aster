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

var _ ForkModel = (*customForkModel)(nil)

type (
	// ForkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customForkModel.
	ForkModel interface {
		forkModel
		FindOneByForkRepoId(ctx context.Context, forkRepoId int64) (*Fork, error)
		SearchFork(ctx context.Context, originalRepoId int64, page int64, limit int64) ([]*Fork, error)
		CreateDataId() (int64, error)
	}

	customForkModel struct {
		*defaultForkModel
		sf *utils.Snowflake
	}
)

// NewForkModel returns a model for the database table.
func NewForkModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) ForkModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customForkModel{
		defaultForkModel: newForkModel(conn, c, opts...),
		sf:               sf,
	}
}

func (m *customForkModel) FindOneByForkRepoId(ctx context.Context, forkRepoId int64) (*Fork, error) {
	cacheForkForkRepoIdKey := fmt.Sprintf("%s%v", "cache:fork:forkRepoId:", forkRepoId)
	var resp Fork
	err := m.QueryRowIndexCtx(ctx, &resp, cacheForkForkRepoIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where fork_repo_id = $1 limit 1", forkRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, forkRepoId); err != nil {
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

func (m *customForkModel) SearchFork(ctx context.Context, originalRepoId int64, page int64, limit int64) ([]*Fork, error) {
	var resp []*Fork

	query := fmt.Sprintf("select %s from %s where original_repo_id = %d limit %d offset %d", forkRows, m.table, originalRepoId, limit, (page-1)*limit)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query); err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *customForkModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
