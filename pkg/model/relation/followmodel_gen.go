// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package relation

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	followFieldNames          = builder.RawFieldNames(&Follow{})
	followRows                = strings.Join(followFieldNames, ",")
	followRowsExpectAutoSet   = strings.Join(stringx.Remove(followFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	followRowsWithPlaceHolder = strings.Join(stringx.Remove(followFieldNames, "`data_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheFollowDataIdPrefix = "cache:follow:dataId:"
)

type (
	followModel interface {
		Insert(ctx context.Context, data *Follow) (sql.Result, error)
		FindOne(ctx context.Context, dataId int64) (*Follow, error)
		Update(ctx context.Context, data *Follow) error
		Delete(ctx context.Context, dataId int64) error
	}

	defaultFollowModel struct {
		sqlc.CachedConn
		table string
	}

	Follow struct {
		DataId        int64        `db:"data_id"`      // Generated Primary Key, Must Not Be Changed
		FollowerId    int64        `db:"follower_id"`  // Follower ID
		FollowingId   int64        `db:"following_id"` // Following ID
		DataCreatedAt time.Time    `db:"data_created_at"`
		DataUpdatedAt time.Time    `db:"data_updated_at"` // update data time
		DataDeletedAt sql.NullTime `db:"data_deleted_at"`
	}
)

func newFollowModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultFollowModel {
	return &defaultFollowModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`follow`",
	}
}

func (m *defaultFollowModel) Delete(ctx context.Context, dataId int64) error {
	followDataIdKey := fmt.Sprintf("%s%v", cacheFollowDataIdPrefix, dataId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `data_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, dataId)
	}, followDataIdKey)
	return err
}

func (m *defaultFollowModel) FindOne(ctx context.Context, dataId int64) (*Follow, error) {
	followDataIdKey := fmt.Sprintf("%s%v", cacheFollowDataIdPrefix, dataId)
	var resp Follow
	err := m.QueryRowCtx(ctx, &resp, followDataIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `data_id` = ? limit 1", followRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, dataId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFollowModel) Insert(ctx context.Context, data *Follow) (sql.Result, error) {
	followDataIdKey := fmt.Sprintf("%s%v", cacheFollowDataIdPrefix, data.DataId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, followRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DataId, data.FollowerId, data.FollowingId, data.DataCreatedAt, data.DataUpdatedAt, data.DataDeletedAt)
	}, followDataIdKey)
	return ret, err
}

func (m *defaultFollowModel) Update(ctx context.Context, data *Follow) error {
	followDataIdKey := fmt.Sprintf("%s%v", cacheFollowDataIdPrefix, data.DataId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `data_id` = ?", m.table, followRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.FollowerId, data.FollowingId, data.DataCreatedAt, data.DataUpdatedAt, data.DataDeletedAt, data.DataId)
	}, followDataIdKey)
	return err
}

func (m *defaultFollowModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheFollowDataIdPrefix, primary)
}

func (m *defaultFollowModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `data_id` = ? limit 1", followRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFollowModel) tableName() string {
	return m.table
}
