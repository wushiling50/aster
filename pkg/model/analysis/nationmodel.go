package analysis

import (
	"context"
	"fmt"

	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NationModel = (*customNationModel)(nil)

type (
	// NationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNationModel.
	NationModel interface {
		nationModel
		FindOneByDeveloperId(ctx context.Context, developerId int64) (*Nation, error)
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

func (m *customNationModel) FindOneByDeveloperId(ctx context.Context, developerId int64) (*Nation, error) {
	cacheNationDeveloperIdKey := fmt.Sprintf("%s%v", "cache:nation:developerId:", developerId)
	var resp Nation
	err := m.QueryRowIndexCtx(ctx, &resp, cacheNationDeveloperIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where developer_id = $1 limit 1", nationRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, developerId); err != nil {
			return nil, err
		}
		return resp.DataId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customNationModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
