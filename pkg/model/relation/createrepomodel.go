package relation

import (
	"context"
	"fmt"

	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CreateRepoModel = (*customCreateRepoModel)(nil)

type (
	// CreateRepoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCreateRepoModel.
	CreateRepoModel interface {
		createRepoModel
		SearchCreatedRepo(ctx context.Context, developerId int64, page int64, limit int64) ([]*CreateRepo, error)
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

func (m *customCreateRepoModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
