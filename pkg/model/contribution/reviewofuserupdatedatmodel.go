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

var _ ReviewOfUserUpdatedAtModel = (*customReviewOfUserUpdatedAtModel)(nil)

type (
	// ReviewOfUserUpdatedAtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customReviewOfUserUpdatedAtModel.
	ReviewOfUserUpdatedAtModel interface {
		reviewOfUserUpdatedAtModel
		FindOneByDeveloperId(ctx context.Context, developerId int64) (*ReviewOfUserUpdatedAt, error)
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

func (m *customReviewOfUserUpdatedAtModel) FindOneByDeveloperId(ctx context.Context, developerId int64) (*ReviewOfUserUpdatedAt, error) {
	cacheReviewOfUserUpdatedAtDeveloperIdKey := fmt.Sprintf("%s%v", "cache:reviewOfUserUpdatedAt:developerId:", developerId)
	var resp ReviewOfUserUpdatedAt
	err := m.QueryRowIndexCtx(ctx, &resp, cacheReviewOfUserUpdatedAtDeveloperIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where developer_id = ? limit 1", reviewOfUserUpdatedAtRows, m.table)
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

func (m *customReviewOfUserUpdatedAtModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
