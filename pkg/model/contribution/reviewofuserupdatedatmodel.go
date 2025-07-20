package contribution

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ReviewOfUserUpdatedAtModel = (*customReviewOfUserUpdatedAtModel)(nil)

type (
	// ReviewOfUserUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customReviewOfUserUpdatedAtModel.
	ReviewOfUserUpdatedAtModel interface {
		reviewOfUserUpdatedAtModel
		CreateDataId() (int64, error)
	}

	customReviewOfUserUpdatedAtModel struct {
		*defaultReviewOfUserUpdatedAtModel
		sf *utils.Snowflake
	}
)

// NewReviewOfUserUpdatedAtModel returns a model for the database table.
func NewReviewOfUserUpdatedAtModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) ReviewOfUserUpdatedAtModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customReviewOfUserUpdatedAtModel{
		defaultReviewOfUserUpdatedAtModel: newReviewOfUserUpdatedAtModel(conn, c, opts...),
		sf:                                sf,
	}
}

func (m *customReviewOfUserUpdatedAtModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
