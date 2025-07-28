package locks

import (
	"context"
	"fmt"
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

func (b *BLock) GetNewLocksKey(lockType string, id int64) string {
	return constants.LockKeyPrefix + constants.LockSeparator +
		lockType + constants.LockSeparator + strconv.Itoa(int(id))
}

func (b *BLock) DelOldLocksKey(ctx context.Context, key string) error {
	_, err := b.client.DelCtx(ctx, key)
	if err != nil {
		err = errno.InternalLockError.WithError(err)
		return err
	}

	return nil
}

func (b *BLock) TryLock(ctx context.Context, key string) (bool, error) {
	result, err := b.client.SetnxExCtx(ctx, key, "", int(20*constants.ONE_SECOND))
	if err != nil {
		err = errno.InternalLockError.WithError(err)
		return false, err
	}
	return result, nil
}

func (b *BLock) TryUnLock(ctx context.Context, key string) error {
	// 使用Lua脚本确保原子性操作
	script := `
	if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("del", KEYS[1])
	else
		return 0
	end
	`

	res, err := b.client.EvalCtx(ctx, script, []string{key})
	fmt.Println(res)
	if err != nil {
		err = errno.InternalLockError.WithError(err)
		return err
	}

	if res != 1 {
		err = errno.InternalLockError.WithMessage("Release Lock Fail")
		return err
	}

	logx.Infof("Release Lock: %s", key)

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
