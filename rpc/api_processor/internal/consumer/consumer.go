package consumer

import (
	"context"
	"errors"
	"strconv"

	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/tasks"
	"github.com/wushiling50/aster/rpc/api_processor/internal/logic"
	"github.com/wushiling50/aster/rpc/api_processor/internal/svc"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type APITaskConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAPITaskConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *APITaskConsumer {
	return &APITaskConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (c *APITaskConsumer) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()

	mux.HandleFunc(constants.APITaskName, c.Consume)

	return mux
}

func (c *APITaskConsumer) Consume(ctx context.Context, task *asynq.Task) error {
	logx.Info("Consume Message: ", task.Type(), task.Payload())

	var (
		err  error
		data []byte
		msg  = tasks.APIPayload{}
	)

	err = jsonx.Unmarshal(task.Payload(), &msg)
	if err != nil {
		return err
	}

	switch msg.Type {
	case constants.APIGetDeveloper:
		logx.Info("Consume Message: APIGetDeveloper")

		l := logic.NewGetDeveloperLogic(c.ctx, c.svcCtx)
		data, err = l.GetDeveloper(msg.Id)
	case constants.APIGetLanguage:
		logx.Info("Consume Message: APIGetLanguage")

		l := logic.NewGetLanguageLogic(c.ctx, c.svcCtx)
		data, err = l.GetLanguage(msg.Id)
	case constants.APIGetScore:
		logx.Info("Consume Message: APIGetScore")

		l := logic.NewGetScoreLogic(c.ctx, c.svcCtx)
		data, err = l.GetScore(msg.Id)
	case constants.APIGetNation:
		logx.Info("Consume Message: APIGetNation")

		l := logic.NewGetNationLogic(c.ctx, c.svcCtx)
		data, err = l.GetNation(msg.Id)
	case constants.APIGetSummary:
		logx.Info("Consume Message: APIGetSummary")

		l := logic.NewGetSummaryLogic(c.ctx, c.svcCtx)
		data, err = l.GetSummary(msg.Id)
	default:
		err = errors.New("Unexpected Message Type: " + strconv.FormatInt(int64(msg.Type), 10))
	}

	if err != nil {
		return err
	}

	_, err = task.ResultWriter().Write(data)

	return err
}
