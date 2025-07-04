package repo

import (
	"context"
	"fmt"

	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
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
		CreateDataId() (int64, error)
	}

	customRepoModel struct {
		*defaultRepoModel
		sf *utils.Snowflake
	}
)

// NewRepoModel returns a model for the database table.
func NewRepoModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) RepoModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customRepoModel{
		defaultRepoModel: newRepoModel(conn, c, opts...),
		sf:               sf,
	}
}

func (m *customRepoModel) FindOneById(ctx context.Context, id int64) (*Repo, error) {
	repoRepoIdKey := fmt.Sprintf("%s%v", "cache:repo:id:", id)
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

func (m *customRepoModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
