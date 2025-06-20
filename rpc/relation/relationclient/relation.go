// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: relation.proto

package relationclient

import (
	"context"

	"github.com/wushiling50/aster/gen/relation"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCreateRepoReq                 = relation.AddCreateRepoReq
	AddCreateRepoResp                = relation.AddCreateRepoResp
	AddFollowReq                     = relation.AddFollowReq
	AddFollowResp                    = relation.AddFollowResp
	AddForkReq                       = relation.AddForkReq
	AddForkResp                      = relation.AddForkResp
	AddStarReq                       = relation.AddStarReq
	AddStarResp                      = relation.AddStarResp
	CheckFollowResp                  = relation.CheckFollowResp
	CheckIfFollowReq                 = relation.CheckIfFollowReq
	CheckIfStarReq                   = relation.CheckIfStarReq
	CheckIfStarResp                  = relation.CheckIfStarResp
	CreateRepo                       = relation.CreateRepo
	DelAllCreatedRepoReq             = relation.DelAllCreatedRepoReq
	DelAllCreatedRepoResp            = relation.DelAllCreatedRepoResp
	DelAllFollowerReq                = relation.DelAllFollowerReq
	DelAllFollowerResp               = relation.DelAllFollowerResp
	DelAllFollowingReq               = relation.DelAllFollowingReq
	DelAllFollowingResp              = relation.DelAllFollowingResp
	DelAllForkReq                    = relation.DelAllForkReq
	DelAllForkResp                   = relation.DelAllForkResp
	DelAllStaringDevReq              = relation.DelAllStaringDevReq
	DelAllStaringDevResp             = relation.DelAllStaringDevResp
	DelAllStarredRepoReq             = relation.DelAllStarredRepoReq
	DelAllStarredRepoResp            = relation.DelAllStarredRepoResp
	DelCreateRepoReq                 = relation.DelCreateRepoReq
	DelCreateRepoResp                = relation.DelCreateRepoResp
	DelFollowReq                     = relation.DelFollowReq
	DelFollowResp                    = relation.DelFollowResp
	DelForkReq                       = relation.DelForkReq
	DelForkResp                      = relation.DelForkResp
	DelStarReq                       = relation.DelStarReq
	DelStarResp                      = relation.DelStarResp
	Follow                           = relation.Follow
	Fork                             = relation.Fork
	GetCreatedRepoUpdatedAtReq       = relation.GetCreatedRepoUpdatedAtReq
	GetCreatedRepoUpdatedAtResp      = relation.GetCreatedRepoUpdatedAtResp
	GetCreatorIdReq                  = relation.GetCreatorIdReq
	GetCreatorIdResp                 = relation.GetCreatorIdResp
	GetFollowerUpdatedAtReq          = relation.GetFollowerUpdatedAtReq
	GetFollowerUpdatedAtResp         = relation.GetFollowerUpdatedAtResp
	GetFollowingUpdatedAtReq         = relation.GetFollowingUpdatedAtReq
	GetFollowingUpdatedAtResp        = relation.GetFollowingUpdatedAtResp
	GetForkUpdatedAtReq              = relation.GetForkUpdatedAtReq
	GetForkUpdatedAtResp             = relation.GetForkUpdatedAtResp
	GetOriginReq                     = relation.GetOriginReq
	GetOriginResp                    = relation.GetOriginResp
	GetStarredRepoUpdatedAtReq       = relation.GetStarredRepoUpdatedAtReq
	GetStarredRepoUpdatedAtResp      = relation.GetStarredRepoUpdatedAtResp
	SearchCreatedRepoReq             = relation.SearchCreatedRepoReq
	SearchCreatedRepoResp            = relation.SearchCreatedRepoResp
	SearchFollowerByDeveloperIdReq   = relation.SearchFollowerByDeveloperIdReq
	SearchFollowerByDeveloperIdResp  = relation.SearchFollowerByDeveloperIdResp
	SearchFollowingByDeveloperIdReq  = relation.SearchFollowingByDeveloperIdReq
	SearchFollowingByDeveloperIdResp = relation.SearchFollowingByDeveloperIdResp
	SearchForkReq                    = relation.SearchForkReq
	SearchForkResp                   = relation.SearchForkResp
	SearchStaringDevReq              = relation.SearchStaringDevReq
	SearchStaringDevResp             = relation.SearchStaringDevResp
	SearchStarredRepoReq             = relation.SearchStarredRepoReq
	SearchStarredRepoResp            = relation.SearchStarredRepoResp
	Star                             = relation.Star
	UpdateCreateRepoReq              = relation.UpdateCreateRepoReq
	UpdateCreateRepoResp             = relation.UpdateCreateRepoResp
	UpdateFollowerReq                = relation.UpdateFollowerReq
	UpdateFollowerResp               = relation.UpdateFollowerResp
	UpdateFollowingReq               = relation.UpdateFollowingReq
	UpdateFollowingResp              = relation.UpdateFollowingResp
	UpdateForkReq                    = relation.UpdateForkReq
	UpdateForkResp                   = relation.UpdateForkResp
	UpdateStarredRepoReq             = relation.UpdateStarredRepoReq
	UpdateStarredRepoResp            = relation.UpdateStarredRepoResp

	Relation interface {
		// -----------------------CreateRepo-----------------------
		AddCreateRepo(ctx context.Context, in *AddCreateRepoReq, opts ...grpc.CallOption) (*AddCreateRepoResp, error)
		DelCreateRepo(ctx context.Context, in *DelCreateRepoReq, opts ...grpc.CallOption) (*DelCreateRepoResp, error)
		DelAllCreatedRepo(ctx context.Context, in *DelAllCreatedRepoReq, opts ...grpc.CallOption) (*DelAllCreatedRepoResp, error)
		GetCreatorId(ctx context.Context, in *GetCreatorIdReq, opts ...grpc.CallOption) (*GetCreatorIdResp, error)
		SearchCreatedRepo(ctx context.Context, in *SearchCreatedRepoReq, opts ...grpc.CallOption) (*SearchCreatedRepoResp, error)
		UpdateCreateRepo(ctx context.Context, in *UpdateCreateRepoReq, opts ...grpc.CallOption) (*UpdateCreateRepoResp, error)
		GetCreatedRepoUpdatedAt(ctx context.Context, in *GetCreatedRepoUpdatedAtReq, opts ...grpc.CallOption) (*GetCreatedRepoUpdatedAtResp, error)
		// -----------------------Follow-----------------------
		AddFollow(ctx context.Context, in *AddFollowReq, opts ...grpc.CallOption) (*AddFollowResp, error)
		DelFollow(ctx context.Context, in *DelFollowReq, opts ...grpc.CallOption) (*DelFollowResp, error)
		DelAllFollower(ctx context.Context, in *DelAllFollowerReq, opts ...grpc.CallOption) (*DelAllFollowerResp, error)
		DelAllFollowing(ctx context.Context, in *DelAllFollowingReq, opts ...grpc.CallOption) (*DelAllFollowingResp, error)
		CheckIfFollow(ctx context.Context, in *CheckIfFollowReq, opts ...grpc.CallOption) (*CheckFollowResp, error)
		SearchFollowingByDeveloperId(ctx context.Context, in *SearchFollowingByDeveloperIdReq, opts ...grpc.CallOption) (*SearchFollowingByDeveloperIdResp, error)
		SearchFollowerByDeveloperId(ctx context.Context, in *SearchFollowerByDeveloperIdReq, opts ...grpc.CallOption) (*SearchFollowerByDeveloperIdResp, error)
		UpdateFollowing(ctx context.Context, in *UpdateFollowingReq, opts ...grpc.CallOption) (*UpdateFollowingResp, error)
		UpdateFollower(ctx context.Context, in *UpdateFollowerReq, opts ...grpc.CallOption) (*UpdateFollowerResp, error)
		GetFollowingUpdatedAt(ctx context.Context, in *GetFollowingUpdatedAtReq, opts ...grpc.CallOption) (*GetFollowingUpdatedAtResp, error)
		GetFollowerUpdatedAt(ctx context.Context, in *GetFollowerUpdatedAtReq, opts ...grpc.CallOption) (*GetFollowerUpdatedAtResp, error)
		// -----------------------Fork-----------------------
		AddFork(ctx context.Context, in *AddForkReq, opts ...grpc.CallOption) (*AddForkResp, error)
		DelFork(ctx context.Context, in *DelForkReq, opts ...grpc.CallOption) (*DelForkResp, error)
		DelAllFork(ctx context.Context, in *DelAllForkReq, opts ...grpc.CallOption) (*DelAllForkResp, error)
		GetOrigin(ctx context.Context, in *GetOriginReq, opts ...grpc.CallOption) (*GetOriginResp, error)
		SearchFork(ctx context.Context, in *SearchForkReq, opts ...grpc.CallOption) (*SearchForkResp, error)
		UpdateFork(ctx context.Context, in *UpdateForkReq, opts ...grpc.CallOption) (*UpdateForkResp, error)
		GetForkUpdatedAt(ctx context.Context, in *GetForkUpdatedAtReq, opts ...grpc.CallOption) (*GetForkUpdatedAtResp, error)
		// -----------------------Star-----------------------
		AddStar(ctx context.Context, in *AddStarReq, opts ...grpc.CallOption) (*AddStarResp, error)
		DelStar(ctx context.Context, in *DelStarReq, opts ...grpc.CallOption) (*DelStarResp, error)
		DelAllStarredRepo(ctx context.Context, in *DelAllStarredRepoReq, opts ...grpc.CallOption) (*DelAllStarredRepoResp, error)
		DelAllStaringDev(ctx context.Context, in *DelAllStaringDevReq, opts ...grpc.CallOption) (*DelAllStaringDevResp, error)
		CheckIfStar(ctx context.Context, in *CheckIfStarReq, opts ...grpc.CallOption) (*CheckIfStarResp, error)
		SearchStarredRepo(ctx context.Context, in *SearchStarredRepoReq, opts ...grpc.CallOption) (*SearchStarredRepoResp, error)
		SearchStaringDev(ctx context.Context, in *SearchStaringDevReq, opts ...grpc.CallOption) (*SearchStaringDevResp, error)
		UpdateStarredRepo(ctx context.Context, in *UpdateStarredRepoReq, opts ...grpc.CallOption) (*UpdateStarredRepoResp, error)
		GetStarredRepoUpdatedAt(ctx context.Context, in *GetStarredRepoUpdatedAtReq, opts ...grpc.CallOption) (*GetStarredRepoUpdatedAtResp, error)
	}

	defaultRelation struct {
		cli zrpc.Client
	}
)

func NewRelation(cli zrpc.Client) Relation {
	return &defaultRelation{
		cli: cli,
	}
}

// -----------------------CreateRepo-----------------------
func (m *defaultRelation) AddCreateRepo(ctx context.Context, in *AddCreateRepoReq, opts ...grpc.CallOption) (*AddCreateRepoResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.AddCreateRepo(ctx, in, opts...)
}

func (m *defaultRelation) DelCreateRepo(ctx context.Context, in *DelCreateRepoReq, opts ...grpc.CallOption) (*DelCreateRepoResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.DelCreateRepo(ctx, in, opts...)
}

func (m *defaultRelation) DelAllCreatedRepo(ctx context.Context, in *DelAllCreatedRepoReq, opts ...grpc.CallOption) (*DelAllCreatedRepoResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.DelAllCreatedRepo(ctx, in, opts...)
}

func (m *defaultRelation) GetCreatorId(ctx context.Context, in *GetCreatorIdReq, opts ...grpc.CallOption) (*GetCreatorIdResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.GetCreatorId(ctx, in, opts...)
}

func (m *defaultRelation) SearchCreatedRepo(ctx context.Context, in *SearchCreatedRepoReq, opts ...grpc.CallOption) (*SearchCreatedRepoResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.SearchCreatedRepo(ctx, in, opts...)
}

func (m *defaultRelation) UpdateCreateRepo(ctx context.Context, in *UpdateCreateRepoReq, opts ...grpc.CallOption) (*UpdateCreateRepoResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.UpdateCreateRepo(ctx, in, opts...)
}

func (m *defaultRelation) GetCreatedRepoUpdatedAt(ctx context.Context, in *GetCreatedRepoUpdatedAtReq, opts ...grpc.CallOption) (*GetCreatedRepoUpdatedAtResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.GetCreatedRepoUpdatedAt(ctx, in, opts...)
}

// -----------------------Follow-----------------------
func (m *defaultRelation) AddFollow(ctx context.Context, in *AddFollowReq, opts ...grpc.CallOption) (*AddFollowResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.AddFollow(ctx, in, opts...)
}

func (m *defaultRelation) DelFollow(ctx context.Context, in *DelFollowReq, opts ...grpc.CallOption) (*DelFollowResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.DelFollow(ctx, in, opts...)
}

func (m *defaultRelation) DelAllFollower(ctx context.Context, in *DelAllFollowerReq, opts ...grpc.CallOption) (*DelAllFollowerResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.DelAllFollower(ctx, in, opts...)
}

func (m *defaultRelation) DelAllFollowing(ctx context.Context, in *DelAllFollowingReq, opts ...grpc.CallOption) (*DelAllFollowingResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.DelAllFollowing(ctx, in, opts...)
}

func (m *defaultRelation) CheckIfFollow(ctx context.Context, in *CheckIfFollowReq, opts ...grpc.CallOption) (*CheckFollowResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.CheckIfFollow(ctx, in, opts...)
}

func (m *defaultRelation) SearchFollowingByDeveloperId(ctx context.Context, in *SearchFollowingByDeveloperIdReq, opts ...grpc.CallOption) (*SearchFollowingByDeveloperIdResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.SearchFollowingByDeveloperId(ctx, in, opts...)
}

func (m *defaultRelation) SearchFollowerByDeveloperId(ctx context.Context, in *SearchFollowerByDeveloperIdReq, opts ...grpc.CallOption) (*SearchFollowerByDeveloperIdResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.SearchFollowerByDeveloperId(ctx, in, opts...)
}

func (m *defaultRelation) UpdateFollowing(ctx context.Context, in *UpdateFollowingReq, opts ...grpc.CallOption) (*UpdateFollowingResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.UpdateFollowing(ctx, in, opts...)
}

func (m *defaultRelation) UpdateFollower(ctx context.Context, in *UpdateFollowerReq, opts ...grpc.CallOption) (*UpdateFollowerResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.UpdateFollower(ctx, in, opts...)
}

func (m *defaultRelation) GetFollowingUpdatedAt(ctx context.Context, in *GetFollowingUpdatedAtReq, opts ...grpc.CallOption) (*GetFollowingUpdatedAtResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.GetFollowingUpdatedAt(ctx, in, opts...)
}

func (m *defaultRelation) GetFollowerUpdatedAt(ctx context.Context, in *GetFollowerUpdatedAtReq, opts ...grpc.CallOption) (*GetFollowerUpdatedAtResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.GetFollowerUpdatedAt(ctx, in, opts...)
}

// -----------------------Fork-----------------------
func (m *defaultRelation) AddFork(ctx context.Context, in *AddForkReq, opts ...grpc.CallOption) (*AddForkResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.AddFork(ctx, in, opts...)
}

func (m *defaultRelation) DelFork(ctx context.Context, in *DelForkReq, opts ...grpc.CallOption) (*DelForkResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.DelFork(ctx, in, opts...)
}

func (m *defaultRelation) DelAllFork(ctx context.Context, in *DelAllForkReq, opts ...grpc.CallOption) (*DelAllForkResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.DelAllFork(ctx, in, opts...)
}

func (m *defaultRelation) GetOrigin(ctx context.Context, in *GetOriginReq, opts ...grpc.CallOption) (*GetOriginResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.GetOrigin(ctx, in, opts...)
}

func (m *defaultRelation) SearchFork(ctx context.Context, in *SearchForkReq, opts ...grpc.CallOption) (*SearchForkResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.SearchFork(ctx, in, opts...)
}

func (m *defaultRelation) UpdateFork(ctx context.Context, in *UpdateForkReq, opts ...grpc.CallOption) (*UpdateForkResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.UpdateFork(ctx, in, opts...)
}

func (m *defaultRelation) GetForkUpdatedAt(ctx context.Context, in *GetForkUpdatedAtReq, opts ...grpc.CallOption) (*GetForkUpdatedAtResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.GetForkUpdatedAt(ctx, in, opts...)
}

// -----------------------Star-----------------------
func (m *defaultRelation) AddStar(ctx context.Context, in *AddStarReq, opts ...grpc.CallOption) (*AddStarResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.AddStar(ctx, in, opts...)
}

func (m *defaultRelation) DelStar(ctx context.Context, in *DelStarReq, opts ...grpc.CallOption) (*DelStarResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.DelStar(ctx, in, opts...)
}

func (m *defaultRelation) DelAllStarredRepo(ctx context.Context, in *DelAllStarredRepoReq, opts ...grpc.CallOption) (*DelAllStarredRepoResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.DelAllStarredRepo(ctx, in, opts...)
}

func (m *defaultRelation) DelAllStaringDev(ctx context.Context, in *DelAllStaringDevReq, opts ...grpc.CallOption) (*DelAllStaringDevResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.DelAllStaringDev(ctx, in, opts...)
}

func (m *defaultRelation) CheckIfStar(ctx context.Context, in *CheckIfStarReq, opts ...grpc.CallOption) (*CheckIfStarResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.CheckIfStar(ctx, in, opts...)
}

func (m *defaultRelation) SearchStarredRepo(ctx context.Context, in *SearchStarredRepoReq, opts ...grpc.CallOption) (*SearchStarredRepoResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.SearchStarredRepo(ctx, in, opts...)
}

func (m *defaultRelation) SearchStaringDev(ctx context.Context, in *SearchStaringDevReq, opts ...grpc.CallOption) (*SearchStaringDevResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.SearchStaringDev(ctx, in, opts...)
}

func (m *defaultRelation) UpdateStarredRepo(ctx context.Context, in *UpdateStarredRepoReq, opts ...grpc.CallOption) (*UpdateStarredRepoResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.UpdateStarredRepo(ctx, in, opts...)
}

func (m *defaultRelation) GetStarredRepoUpdatedAt(ctx context.Context, in *GetStarredRepoUpdatedAtReq, opts ...grpc.CallOption) (*GetStarredRepoUpdatedAtResp, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.GetStarredRepoUpdatedAt(ctx, in, opts...)
}
