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

var _ LanguagesModel = (*customLanguagesModel)(nil)

type (
	// LanguagesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLanguagesModel.
	LanguagesModel interface {
		languagesModel
		FindOneByDeveloperId(ctx context.Context, developerId int64) (*Languages, error)
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

func (m *customLanguagesModel) FindOneByDeveloperId(ctx context.Context, developerId int64) (*Languages, error) {
	cacheLanguagesDeveloperIdKey := fmt.Sprintf("%s%v", "cache:languages:developerId:", developerId)
	var resp Languages
	err := m.QueryRowIndexCtx(ctx, &resp, cacheLanguagesDeveloperIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where developer_id = $1 limit 1", languagesRows, m.table)
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

func (m *customLanguagesModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
