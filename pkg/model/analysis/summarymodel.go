package analysis

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SummaryModel = (*customSummaryModel)(nil)

type (
	// SummaryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSummaryModel.
	SummaryModel interface {
		summaryModel
		CreateDataId() (int64, error)
	}

	customSummaryModel struct {
		*defaultSummaryModel
		sf *utils.Snowflake
	}
)

// NewSummaryModel returns a model for the database table.
func NewSummaryModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) SummaryModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customSummaryModel{
		defaultSummaryModel: newSummaryModel(conn, c, opts...),
		sf:                  sf,
	}
}

func (m *customSummaryModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
