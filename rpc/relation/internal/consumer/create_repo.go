package consumer

import (
	"context"

	"github.com/wushiling50/aster/rpc/relation/internal/svc"
)

type CreateRepoConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewCreateRepoConsumer(ctx context.Context, svc *svc.ServiceContext) *CreateRepoConsumer {
	return &CreateRepoConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *CreateRepoConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	return
}
