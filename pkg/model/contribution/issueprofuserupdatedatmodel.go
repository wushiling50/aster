package contribution

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ IssuePrOfUserUpdatedAtModel = (*customIssuePrOfUserUpdatedAtModel)(nil)

type (
	// IssuePrOfUserUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIssuePrOfUserUpdatedAtModel.
	IssuePrOfUserUpdatedAtModel interface {
		issuePrOfUserUpdatedAtModel
		CreateDataId() (int64, error)
	}

	customIssuePrOfUserUpdatedAtModel struct {
		*defaultIssuePrOfUserUpdatedAtModel
		sf *utils.Snowflake
	}
)

// NewIssuePrOfUserUpdatedAtModel returns a model for the database table.
func NewIssuePrOfUserUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) IssuePrOfUserUpdatedAtModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customIssuePrOfUserUpdatedAtModel{
		defaultIssuePrOfUserUpdatedAtModel: newIssuePrOfUserUpdatedAtModel(conn, c, opts...),
		sf:                                 sf,
	}
}

func (m *customIssuePrOfUserUpdatedAtModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
