package rank

import (
	"context"

	"github.com/wushiling50/aster/pkg/errno"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func (r *RankModel) GetScores(ctx context.Context, key string, start int64, stop int64) ([]redis.FloatPair, error) {
	rank, err := r.cache.ZrevrangeWithScoresByFloatCtx(ctx, key, start, stop)
	if err != nil {
		logx.Errorf("model.GetScores: Get Cache Failed: %v", err.Error())
		return nil, errno.InternalRedisError.WithMessage(err.Error())
	}

	return rank, nil
}

func (r *RankModel) GetScoresTotal(ctx context.Context, key string) (int64, error) {
	total, err := r.cache.ZcardCtx(ctx, key)
	if err != nil {
		logx.Errorf("model.GetScoresNum: Get Cache Failed: %v", err.Error())
		return 0, errno.InternalRedisError.WithMessage(err.Error())
	}

	return int64(total), nil
}
