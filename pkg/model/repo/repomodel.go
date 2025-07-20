package repo

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RepoModel = (*customRepoModel)(nil)

type (
	// RepoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRepoModel.
	RepoModel interface {
		repoModel
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

func (m *customRepoModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
