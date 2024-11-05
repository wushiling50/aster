package service

import (
	"context"

	"github.com/wushiling50/aster/cmd/api/biz/model/api"
)

func (s *BaseService) TalentRank(ctx context.Context, req *api.TalentRankRequest) (userInfo *api.User, err error) {
	userInfo = new(api.User)
	return
}
