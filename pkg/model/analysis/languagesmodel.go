package analysis

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LanguagesModel = (*customLanguagesModel)(nil)

type (
	// LanguagesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLanguagesModel.
	LanguagesModel interface {
		languagesModel
		CreateDataId() (int64, error)
	}

	customLanguagesModel struct {
		*defaultLanguagesModel
		sf *utils.Snowflake
	}
)

// NewLanguagesModel returns a model for the database table.
func NewLanguagesModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) LanguagesModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customLanguagesModel{
		defaultLanguagesModel: newLanguagesModel(conn, c, opts...),
		sf:                    sf,
	}
}

func (m *customLanguagesModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
