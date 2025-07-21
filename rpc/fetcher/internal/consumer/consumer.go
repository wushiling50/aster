package consumer

import (
	"context"
	"errors"
	"strconv"

	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/tasks"
	"github.com/wushiling50/aster/rpc/fetcher/internal/logic"
	"github.com/wushiling50/aster/rpc/fetcher/internal/svc"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type FetcherTaskConsumer struct {
	svcCtx *svc.ServiceContext
}

func NewFetcherTaskConsumer(svcCtx *svc.ServiceContext) *FetcherTaskConsumer {
	return &FetcherTaskConsumer{
		svcCtx: svcCtx,
	}
}

func (c *FetcherTaskConsumer) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()

	mux.HandleFunc(constants.FetcherTaskName, c.Consume)

	return mux
}

func (c *FetcherTaskConsumer) Consume(ctx context.Context, task *asynq.Task) (err error) {
	logx.Info("Consume Message: ", task.Type(), task.Payload())

	msg := tasks.FetchPayload{}
	if err = jsonx.Unmarshal(task.Payload(), &msg); err != nil {
		return
	}

	switch msg.Type {
	case constants.FetchDeveloper:
		logx.Info("Consume Message: FetchDeveloper")

		l := logic.NewFetchDeveloperLogic(ctx, c.svcCtx)
		err = l.FetchDeveloper(msg.Id)
	case constants.FetchRepo:
		logx.Info("Consume Message: FetchRepo")

		l := logic.NewFetchRepoLogic(ctx, c.svcCtx)
		err = l.FetchRepo(msg.Id)
	case constants.FetchCreatedRepo:
		logx.Info("Consume Message: FetchCreatedRepo")

		l := logic.NewFetchCreatedRepoLogic(ctx, c.svcCtx)
		err = l.FetchCreatedRepo(msg.Id)
	case constants.FetchStarredRepo:
		logx.Info("Consume Message: FetchStarredRepo")

		l := logic.NewFetchStarredRepoLogic(ctx, c.svcCtx)
		err = l.FetchStarredRepo(msg.Id)
	case constants.FetchFollower:
		logx.Info("Consume Message: FetchFollower")

		l := logic.NewFetchFollowerLogic(ctx, c.svcCtx)
		err = l.FetchFollower(msg.Id)
	case constants.FetchFollowing:
		logx.Info("Consume Message: FetchFollowing")

		l := logic.NewFetchFollowingLogic(ctx, c.svcCtx)
		err = l.FetchFollowing(msg.Id)
	case constants.FetchIssuePROfUser:
		logx.Info("Consume Message: FetchIssuePROfUser")

		l := logic.NewFetchIssuePROfUserLogic(ctx, c.svcCtx)
		err = l.FetchIssuePROfUser(msg.Id, msg.UpdateAfter, msg.SearchLimit)
	case constants.FetchCommentOfUser:
		logx.Info("Consume Message: FetchCommentOfUser")

		l := logic.NewFetchCommentOfUserLogic(ctx, c.svcCtx)
		err = l.FetchCommentOfUser(msg.Id, msg.UpdateAfter, msg.SearchLimit)
	case constants.FetchReviewOfUser:
		logx.Info("Consume Message: FetchReviewOfUser")

		l := logic.NewFetchReviewOfUserLogic(ctx, c.svcCtx)
		err = l.FetchReviewOfUser(msg.Id, msg.UpdateAfter, msg.SearchLimit)
	default:
		err = errors.New("Unexpected Message Type: " + strconv.FormatInt(int64(msg.Type), 10))
	}

	return
}
