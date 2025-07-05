package relation

import (
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

func (m *customCreateRepoModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
