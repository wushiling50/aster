package analysis

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NationModel = (*customNationModel)(nil)

type (
	// NationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNationModel.
	NationModel interface {
		nationModel
		CreateDataId() (int64, error)
	}

	customNationModel struct {
		*defaultNationModel
		sf *utils.Snowflake
	}
)

// NewNationModel returns a model for the database table.
func NewNationModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) NationModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customNationModel{
		defaultNationModel: newNationModel(conn, c, opts...),
		sf:                 sf,
	}
}

func (m *customNationModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
