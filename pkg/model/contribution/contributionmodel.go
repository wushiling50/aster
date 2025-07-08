package contribution

import (
	"context"
	"fmt"

	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContributionModel = (*customContributionModel)(nil)

type (
	// ContributionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContributionModel.
	ContributionModel interface {
		contributionModel
		SearchByCategory(ctx context.Context, category string, page int64, limit int64) ([]*Contribution, error)
		SearchByDeveloperId(ctx context.Context, developerId int64, page int64, limit int64) ([]*Contribution, error)
		SearchByRepoId(ctx context.Context, repoId int64, page int64, limit int64) ([]*Contribution, error)
		CreateDataId() (int64, error)
	}

	customContributionModel struct {
		*defaultContributionModel
		sf *utils.Snowflake
	}
)

// NewContributionModel returns a model for the database table.
func NewContributionModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) ContributionModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customContributionModel{
		defaultContributionModel: newContributionModel(conn, c, opts...),
		sf:                       sf,
	}
}

func (m *customContributionModel) SearchByCategory(ctx context.Context, category string, page int64, limit int64) ([]*Contribution, error) {
	var resp []*Contribution

	query := fmt.Sprintf("select %s from %s where category = '%s' limit %d offset %d", contributionRows, m.table, category, limit, (page-1)*limit)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query); err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *customContributionModel) SearchByDeveloperId(ctx context.Context, developerId int64, page int64, limit int64) ([]*Contribution, error) {
	var resp []*Contribution

	query := fmt.Sprintf("select %s from %s where developer_id = %d limit %d offset %d", contributionRows, m.table, developerId, limit, (page-1)*limit)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query); err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *customContributionModel) SearchByRepoId(ctx context.Context, repoId int64, page int64, limit int64) ([]*Contribution, error) {
	var resp []*Contribution

	query := fmt.Sprintf("select %s from %s where repo_id = %d limit %d offset %d", contributionRows, m.table, repoId, limit, (page-1)*limit)
	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query); err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *customContributionModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
