package svc

import (
	"time"

	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/constants"
	analysis "github.com/wushiling50/aster/rpc/analysis/analysisclient"
	"github.com/wushiling50/aster/rpc/api_processor/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	AsynqServer *asynq.Server

	AnalysisRpcClient analysis.Analysis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		AsynqServer: asynq.NewServer(
			asynq.RedisClientOpt{
				Addr:     c.AsynqRedisConf.Addr,
				Password: c.AsynqRedisConf.Pass,
				DB:       c.AsynqRedisConf.DB,
			}, asynq.Config{
				Concurrency: constants.APIConcurrency,
				RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
					return constants.APIRetryDelay
				},
				Queues: map[string]int{constants.APITaskQueue: 1},
			}),

		AnalysisRpcClient: analysis.NewAnalysis(zrpc.MustNewClient(c.Services.Analysis)),
	}
}
