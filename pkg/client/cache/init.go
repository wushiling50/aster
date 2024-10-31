package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/wushiling50/aster/config"
	"github.com/wushiling50/aster/pkg/constants"
)

var (
	RCScore  *redis.Client
	RCNation *redis.Client
)

func Init() {
	RCScore = RCInit(constants.RedisDBScore)
	RCNation = RCInit(constants.RedisDBNation)
}

func RCInit(dbname int) *redis.Client {
	rc := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       dbname,
	})

	_, err := rc.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
	return rc
}
