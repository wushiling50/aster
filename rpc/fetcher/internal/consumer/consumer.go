package consumer

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/rpc/fetcher/internal/svc"
)

type FetcherTaskConsumer struct {
	svc *svc.ServiceContext
}

func NewFetcherTaskConsumer(ctx context.Context, svc *svc.ServiceContext) *FetcherTaskConsumer {
	return &FetcherTaskConsumer{
		svc: svc,
	}
}

func (c *FetcherTaskConsumer) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()

	mux.HandleFunc(constants.FetcherTaskName, c.Consume)

	return mux
}

func (c *FetcherTaskConsumer) Consume(ctx context.Context, task *asynq.Task) (err error) {
	return nil
}
