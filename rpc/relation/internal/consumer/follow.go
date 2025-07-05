package consumer

import (
	"context"

	"github.com/wushiling50/aster/rpc/relation/internal/svc"
)

type FollowConsumer struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewFollowConsumer(ctx context.Context, svc *svc.ServiceContext) *FollowConsumer {
	return &FollowConsumer{
		ctx: ctx,
		svc: svc,
	}
}

func (c *FollowConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	return
}
