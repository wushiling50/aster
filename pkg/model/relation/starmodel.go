package relation

import (
	"context"
	"fmt"

	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StarModel = (*customStarModel)(nil)

type (
	// StarModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStarModel.
	StarModel interface {
		starModel
		SearchStarredRepo(ctx context.Context, developerId int64, page int64, limit int64) ([]*Star, error)
		SearchStaringDeveloper(ctx context.Context, repoId int64, page int64, limit int64) ([]*Star, error)
		CreateDataId() (int64, error)
	}

	customStarModel struct {
		*defaultStarModel
		sf *utils.Snowflake
	}
)

// NewStarModel returns a model for the database table.
func NewStarModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) StarModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customStarModel{
		defaultStarModel: newStarModel(conn, c, opts...),
		sf:               sf,
	}
}

func (m *customStarModel) SearchStarredRepo(ctx context.Context, developerId int64, page int64, limit int64) ([]*Star, error) {
	var resp []*Star

	query := fmt.Sprintf("select %s from %s where developer_id = %d limit %d offset %d", starRows, m.table, developerId, limit, (page-1)*limit)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query); err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *customStarModel) SearchStaringDeveloper(ctx context.Context, repoId int64, page int64, limit int64) ([]*Star, error) {
	var resp []*Star

	query := fmt.Sprintf("select %s from %s where repo_id = %d limit %d offset %d", starRows, m.table, repoId, limit, (page-1)*limit)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query); err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *customStarModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
