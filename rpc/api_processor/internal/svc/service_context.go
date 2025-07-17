package svc

import (
	"time"

	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/constants"
	analysis "github.com/wushiling50/aster/rpc/analysis/analysisclient"
	"github.com/wushiling50/aster/rpc/api_processor/internal/config"
	developer "github.com/wushiling50/aster/rpc/developer/developerclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	AsynqServer *asynq.Server

	DeveloperRpcClient developer.DeveloperZrpcClient
	AnalysisRpcClient  analysis.Analysis
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

		DeveloperRpcClient: developer.NewDeveloperZrpcClient(zrpc.MustNewClient(c.Services.Developer)),
		AnalysisRpcClient:  analysis.NewAnalysis(zrpc.MustNewClient(c.Services.Analysis)),
	}
}
