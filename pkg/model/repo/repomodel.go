package repo

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RepoModel = (*customRepoModel)(nil)

type (
	// RepoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRepoModel.
	RepoModel interface {
		repoModel
		FindOneById(ctx context.Context, id int64) (*Repo, error)
	}

	customRepoModel struct {
		*defaultRepoModel
	}
)

// NewRepoModel returns a model for the database table.
func NewRepoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RepoModel {
	return &customRepoModel{
		defaultRepoModel: newRepoModel(conn, c, opts...),
	}
}

func (m *customRepoModel) FindOneById(ctx context.Context, id int64) (*Repo, error) {
	repoRepoIdKey := fmt.Sprintf("%s%v", "cache:repo:dataId:", id)
	var resp Repo
	err := m.QueryRowIndexCtx(ctx, &resp, repoRepoIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", repoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, id); err != nil {
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
