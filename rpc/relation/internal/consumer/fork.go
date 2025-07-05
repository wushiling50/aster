package consumer

import (
	"context"

	"github.com/wushiling50/aster/rpc/relation/internal/svc"
)

type ForkConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewForkConsumer(ctx context.Context, svc *svc.ServiceContext) *ForkConsumer {
	return &ForkConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *ForkConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	return
}
