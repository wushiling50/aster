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

var _ ScoreModel = (*customScoreModel)(nil)

type (
	// ScoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customScoreModel.
	ScoreModel interface {
		scoreModel
		FindOneByDeveloperId(ctx context.Context, developerId int64) (*Score, error)
		CreateDataId() (int64, error)
	}

	customScoreModel struct {
		*defaultScoreModel
		sf *utils.Snowflake
	}
)

// NewScoreModel returns a model for the database table.
func NewScoreModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) ScoreModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customScoreModel{
		defaultScoreModel: newScoreModel(conn, c, opts...),
		sf:                sf,
	}
}

func (m *customScoreModel) FindOneByDeveloperId(ctx context.Context, developerId int64) (*Score, error) {
	cacheScoreDeveloperIdKey := fmt.Sprintf("%s%v", "cache:score:developerId:", developerId)
	var resp Score
	err := m.QueryRowIndexCtx(ctx, &resp, cacheScoreDeveloperIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where developer_id = ? limit 1", scoreRows, m.table)
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

func (m *customScoreModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
