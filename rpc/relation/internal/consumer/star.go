package consumer

import (
	"context"

	"github.com/wushiling50/aster/rpc/relation/internal/svc"
)

type StarConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewStarConsumer(ctx context.Context, svc *svc.ServiceContext) *StarConsumer {
	return &StarConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *StarConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	return
}
