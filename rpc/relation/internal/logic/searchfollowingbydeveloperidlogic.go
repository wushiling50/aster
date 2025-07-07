package logic

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"
	"github.com/wushiling50/aster/rpc/relation/internal/pack"
	"github.com/wushiling50/aster/rpc/relation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFollowingByDeveloperIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchFollowingByDeveloperIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFollowingByDeveloperIdLogic {
	return &SearchFollowingByDeveloperIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchFollowingByDeveloperIdLogic) SearchFollowingByDeveloperId(in *relation.SearchFollowingByDeveloperIdReq) (*relation.SearchFollowingByDeveloperIdResp, error) {
	resp := new(relation.SearchFollowingByDeveloperIdResp)

	followingIds, err := l.searchFollowing(in.DeveloperId, in.Page, in.Limit)
	if err != nil {
		logx.Errorf("service.SearchFollowingByDeveloperId: Search Following Failed: %w", err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	if len(followingIds) == 0 {
		logx.Info("service.SearchFollowingByDeveloperId: No Found Following")
	}

	resp.FollowingIds = followingIds
	resp.Base = pack.BuildSuccessResp()

	return resp, nil
}

func (l *SearchFollowingByDeveloperIdLogic) searchFollowing(developerId, page, limit int64) ([]int64, error) {
	followings, err := l.svcCtx.FollowModel.SearchFollowingByDeveloperId(l.ctx, developerId, page, limit)
	if err != nil {
		return nil, err
	}

	followingIds := make([]int64, len(followings))
	for i, following := range followings {
		followingIds[i] = following.FollowerId
	}

	return followingIds, nil
}
