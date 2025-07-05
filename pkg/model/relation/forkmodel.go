package relation

import (
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ForkModel = (*customForkModel)(nil)

type (
	// ForkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customForkModel.
	ForkModel interface {
		forkModel
		CreateDataId() (int64, error)
	}

	customForkModel struct {
		*defaultForkModel
		sf *utils.Snowflake
	}
)

// NewForkModel returns a model for the database table.
func NewForkModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) ForkModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customForkModel{
		defaultForkModel: newForkModel(conn, c, opts...),
		sf:               sf,
	}
}

func (m *customForkModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
