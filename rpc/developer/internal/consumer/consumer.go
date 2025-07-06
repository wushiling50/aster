package consumer

import (
	"context"

	"github.com/wushiling50/aster/gen/developer"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/developer/internal/config"
	"github.com/wushiling50/aster/rpc/developer/internal/logic"
	"github.com/wushiling50/aster/rpc/developer/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

type DeveloperConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeveloperConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *DeveloperConsumer {
	return &DeveloperConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func Consumers(c config.Config, ctx context.Context, svc *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.KqDeveloperConsumerConf, NewDeveloperConsumer(ctx, svc)),
	}
}

func (c *DeveloperConsumer) Consume(ctx context.Context, key string, value string) (err error) {
	logx.Info("Consume Message: ", value)

	var (
		newDeveloper *developer.Developer
		exist        bool
	)

	if err = jsonx.UnmarshalFromString(value, &newDeveloper); err != nil {
		err = errno.InternalJSONError.WithError(err)
		logx.Error(err)
		return
	}

	if _, exist, err = c.getDeveloper(newDeveloper.Id); err != nil {
		logx.Error(err)
		return
	}

	if exist {
		err = c.updateOldDeveloper(newDeveloper)
		if err != nil {
			logx.Error(err)
			return
		}
	} else {
		err = c.addNewDeveloper(newDeveloper)
		if err != nil {
			logx.Error(err)
			return
		}
	}

	return
}

func (c *DeveloperConsumer) getDeveloper(developerId int64) (*developer.Developer, bool, error) {
	l := logic.NewGetDeveloperByIdLogic(c.ctx, c.svcCtx)

	resp, err := l.GetDeveloperById(&developer.GetDeveloperByIdReq{
		Id: developerId,
	})

	if err != nil {
		err = errno.InternalServiceError.WithError(err)
		return nil, false, err
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return nil, false, err
	}

	if resp.Developer == nil {
		logx.Info("No Found This Developer!")
		return nil, false, nil
	}

	return resp.Developer, true, nil
}

func (c *DeveloperConsumer) updateOldDeveloper(newDeveloper *developer.Developer) error {
	l := logic.NewUpdateDeveloperLogic(c.ctx, c.svcCtx)

	resp, err := l.UpdateDeveloper(&developer.UpdateDeveloperReq{
		Developer: newDeveloper,
	})

	if err != nil {
		err = errno.InternalServiceError.WithError(err)
		return err
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return err
	}

	return nil
}

func (c *DeveloperConsumer) addNewDeveloper(newDeveloper *developer.Developer) error {
	l := logic.NewAddDeveloperLogic(c.ctx, c.svcCtx)

	resp, err := l.AddDeveloper(&developer.AddDeveloperReq{
		Developer: newDeveloper,
	})

	if err != nil {
		err = errno.InternalServiceError.WithError(err)
		return err
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return err
	}

	return nil
}
