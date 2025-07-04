package developer

import (
	"context"
	"fmt"

	"github.com/wushiling50/aster/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DeveloperModel = (*customDeveloperModel)(nil)

type (
	// DeveloperModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeveloperModel.
	DeveloperModel interface {
		developerModel
		FindOneById(ctx context.Context, id int64) (*Developer, error)
		FindOneByLogin(ctx context.Context, login string) (*Developer, error)
		CreateDataId() (int64, error)
	}

	customDeveloperModel struct {
		*defaultDeveloperModel
		sf *utils.Snowflake
	}
)

// NewDeveloperModel returns a model for the database table.
func NewDeveloperModel(conn sqlx.SqlConn, c cache.CacheConf, DatancenterId, WorkerId int64, opts ...cache.Option) DeveloperModel {
	sf, err := utils.NewSnowflake(DatancenterId, WorkerId)
	if err != nil {
		logx.Errorf("Init Snowflake Object Error: %v", err.Error())
	}

	return &customDeveloperModel{
		defaultDeveloperModel: newDeveloperModel(conn, c, opts...),
		sf:                    sf,
	}
}

func (m *customDeveloperModel) FindOneById(ctx context.Context, id int64) (*Developer, error) {
	developerDeveloperIdKey := fmt.Sprintf("%s%v", "cache:developer:id:", id)
	var resp Developer
	err := m.QueryRowIndexCtx(ctx, &resp, developerDeveloperIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", developerRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, id); err != nil {
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

func (m *defaultDeveloperModel) FindOneByLogin(ctx context.Context, login string) (*Developer, error) {
	developerDeveloperLoginKey := fmt.Sprintf("%s%v", "cache:developer:login:", login)
	var resp Developer
	err := m.QueryRowIndexCtx(ctx, &resp, developerDeveloperLoginKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where login = $1 limit 1", developerRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, login); err != nil {
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

func (m *customDeveloperModel) CreateDataId() (int64, error) {
	return m.sf.NextVal()
}
