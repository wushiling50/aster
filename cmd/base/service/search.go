package service

import (
	"context"

	"github.com/wushiling50/aster/cmd/api/biz/model/api"
)

func (s *BaseService) Search(ctx context.Context, req *api.SearchRequest) (userList []*api.User, err error) {
	userList = make([]*api.User, 0)
	return
}
