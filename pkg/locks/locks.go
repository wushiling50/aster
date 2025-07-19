package locks

import (
	"context"
	"strconv"
	"time"

	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type BLock struct {
	client *redis.Redis
	expire int64
}

func NewBLock(client *redis.Redis, expire int64) *BLock {
	return &BLock{
		client: client,
		expire: expire,
	}
}

func (b *BLock) GetNewLocksKey(updateType string, id int64) string {
	return constants.LockKeyPrefix + constants.LockSeparator +
		updateType + constants.LockSeparator + strconv.Itoa(int(id))
}

func (b *BLock) DelOldLocksKey(ctx context.Context, key string) error {
	_, err := b.client.DelCtx(ctx, key)
	if err != nil {
		err = errno.InternalLockError.WithError(err)
		return err
	}

	return nil
}

func (b *BLock) Block(ctx context.Context, key string) error {
	node, err := redis.CreateBlockingNode(b.client)
	if err != nil {
		err = errno.InternalLockError.WithError(err)
		return err
	}
	defer node.Close()

	logx.Infof("Block: %s", key)

	_, err = b.client.BlpopWithTimeoutCtx(ctx, node, time.Duration(b.expire)*time.Millisecond, key)
	if err != nil {
		err = errno.InternalLockError.WithError(err)
		return err
	}

	return nil
}

func (b *BLock) Unblock(ctx context.Context, key string) error {
	if _, err := b.client.LpushCtx(ctx, key, ""); err != nil {
		err = errno.InternalLockError.WithError(err)
		return err
	}

	logx.Infof("Unblock: %s", key)
	return nil
}
