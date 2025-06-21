package contribution

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContributionModel = (*customContributionModel)(nil)

type (
	// ContributionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContributionModel.
	ContributionModel interface {
		contributionModel
	}

	customContributionModel struct {
		*defaultContributionModel
	}
)

// NewContributionModel returns a model for the database table.
func NewContributionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ContributionModel {
	return &customContributionModel{
		defaultContributionModel: newContributionModel(conn, c, opts...),
	}
}
