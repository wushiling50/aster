package relation

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CreatedRepoUpdatedAtModel = (*customCreatedRepoUpdatedAtModel)(nil)

type (
	// CreatedRepoUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCreatedRepoUpdatedAtModel.
	CreatedRepoUpdatedAtModel interface {
		createdRepoUpdatedAtModel
		CreateDataId() (int64, error)
	}

	customCreatedRepoUpdatedAtModel struct {
		*defaultCreatedRepoUpdatedAtModel
		sf *utils.Snowflake
	}
)

// NewCreatedRepoUpdatedAtModel returns a model for the database table.
func NewCreatedRepoUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) CreatedRepoUpdatedAtModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customCreatedRepoUpdatedAtModel{
		defaultCreatedRepoUpdatedAtModel: newCreatedRepoUpdatedAtModel(conn, c, opts...),
		sf:                               sf,
	}
}

func (m *customCreatedRepoUpdatedAtModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
