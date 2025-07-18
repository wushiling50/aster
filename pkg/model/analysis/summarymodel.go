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

var _ SummaryModel = (*customSummaryModel)(nil)

type (
	// SummaryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSummaryModel.
	SummaryModel interface {
		summaryModel
		FindOneByDeveloperId(ctx context.Context, developerId int64) (*Summary, error)
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

func (m *defaultSummaryModel) FindOneByDeveloperId(ctx context.Context, developerId int64) (*Summary, error) {
	cacheSummaryDeveloperIdKey := fmt.Sprintf("%s%v", "cache:summary:developerId:", developerId)
	var resp Summary
	err := m.QueryRowIndexCtx(ctx, &resp, cacheSummaryDeveloperIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where developer_id = ? limit 1", summaryRows, m.table)
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

func (m *customSummaryModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
