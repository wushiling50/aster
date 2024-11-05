package service

import "context"

type BaseService struct {
	ctx context.Context
}

func NewBaseService(ctx context.Context) *BaseService {
	return &BaseService{
		ctx: ctx,
	}
}
