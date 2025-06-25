package rank

import "github.com/redis/go-redis/v9"

type CacheRank struct {
	client *redis.Client
}

func NewCacheRank(client *redis.Client) *CacheRank {
	return &CacheRank{
		client: client,
	}
}
