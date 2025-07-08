package contribution

import (
	"context"
	"fmt"

	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ IssuePrOfUserUpdatedAtModel = (*customIssuePrOfUserUpdatedAtModel)(nil)

type (
	// IssuePrOfUserUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIssuePrOfUserUpdatedAtModel.
	IssuePrOfUserUpdatedAtModel interface {
		issuePrOfUserUpdatedAtModel
		FindOneByDeveloperId(ctx context.Context, developerId int64) (*IssuePrOfUserUpdatedAt, error)
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

func (m *customIssuePrOfUserUpdatedAtModel) FindOneByDeveloperId(ctx context.Context, developerId int64) (*IssuePrOfUserUpdatedAt, error) {
	cacheIssuePrOfUserUpdatedAtDeveloperIdKey := fmt.Sprintf("%s%v", "cache:issuePrOfUserUpdatedAt:developerId:", developerId)
	var resp IssuePrOfUserUpdatedAt
	err := m.QueryRowIndexCtx(ctx, &resp, cacheIssuePrOfUserUpdatedAtDeveloperIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where developer_id = $1 limit 1", issuePrOfUserUpdatedAtRows, m.table)
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

func (m *customIssuePrOfUserUpdatedAtModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
