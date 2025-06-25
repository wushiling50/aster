package rank

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var CacheCli *CacheRank

type CacheRank struct {
	client *redis.Redis
}

func NewCacheRank(client *redis.Redis) *CacheRank {
	return &CacheRank{
		client: client,
	}
}
