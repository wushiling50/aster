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

var _ CreateRepoModel = (*customCreateRepoModel)(nil)

type (
	// CreateRepoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCreateRepoModel.
	CreateRepoModel interface {
		createRepoModel
		SearchCreatedRepo(ctx context.Context, developerId int64, page int64, limit int64) ([]*CreateRepo, error)
		FindOneByRepoId(ctx context.Context, repoId int64) (*CreateRepo, error)
		CreateDataId() (int64, error)
	}

	customCreateRepoModel struct {
		*defaultCreateRepoModel
		sf *utils.Snowflake
	}
)

// NewCreateRepoModel returns a model for the database table.
func NewCreateRepoModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) CreateRepoModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customCreateRepoModel{
		defaultCreateRepoModel: newCreateRepoModel(conn, c, opts...),
		sf:                     sf,
	}
}

func (m *customCreateRepoModel) SearchCreatedRepo(ctx context.Context, developerId int64, page int64, limit int64) ([]*CreateRepo, error) {
	var resp []*CreateRepo

	query := fmt.Sprintf("select %s from %s where developer_id = %d limit %d offset %d", createRepoRows, m.table, developerId, limit, (page-1)*limit)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query); err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *customCreateRepoModel) FindOneByRepoId(ctx context.Context, repoId int64) (*CreateRepo, error) {
	cacheCreateRepoRepoIdPrefix := fmt.Sprintf("%s%v", "cache:createRepo:repoId:", repoId)
	var resp CreateRepo
	err := m.QueryRowIndexCtx(ctx, &resp, cacheCreateRepoRepoIdPrefix, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where repo_id = $1 limit 1", createRepoRows, m.table)
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

func (m *customCreateRepoModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
