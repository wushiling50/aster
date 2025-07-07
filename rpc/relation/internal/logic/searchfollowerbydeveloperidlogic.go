package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/pack"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFollowerByDeveloperIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchFollowerByDeveloperIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFollowerByDeveloperIdLogic {
	return &SearchFollowerByDeveloperIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchFollowerByDeveloperIdLogic) SearchFollowerByDeveloperId(in *relation.SearchFollowerByDeveloperIdReq) (*relation.SearchFollowerByDeveloperIdResp, error) {
	resp := new(relation.SearchFollowerByDeveloperIdResp)

	followerIds, err := l.searchFollower(in.DeveloperId, in.Page, in.Limit)
	if err != nil {
		logx.Errorf("service.SearchFollowerByDeveloperId: Search Follower Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if len(followerIds) == 0 {
		logx.Info("service.SearchFollowerByDeveloperId: No Found Follower")
	}

	resp.FollowerIds = followerIds
	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *SearchFollowerByDeveloperIdLogic) searchFollower(developerId, page, limit int64) ([]int64, error) {
	followers, err := l.svcCtx.FollowModel.SearchFollowerByDeveloperId(l.ctx, developerId, page, limit)
	if err != nil {
		return nil, err
	}

	followerIds := make([]int64, len(followers))
	for i, follower := range followers {
		followerIds[i] = follower.FollowerId
	}

	return followerIds, nil
}
