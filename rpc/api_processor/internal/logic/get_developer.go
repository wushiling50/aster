package logic

import (
	"context"
	"encoding/json"

	"github.com/wushiling50/aster/gen/developer"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/wushiling50/aster/pkg/utils"
	"github.com/wushiling50/aster/rpc/api_processor/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeveloperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeveloperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeveloperLogic {
	return &GetDeveloperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeveloperLogic) GetDeveloper(id int64) ([]byte, error) {
	if err := l.rpcUpdateDeveloper(id); err != nil {
		logx.Error(err)
		return nil, err
	}

	resp, err := l.rpcGetDeveloperById(id)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	data, err := json.Marshal(resp.Developer)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (l *GetDeveloperLogic) rpcUpdateDeveloper(id int64) (err error) {
	var resp *developer.UpdateDeveloperResp

	resp, err = l.svcCtx.DeveloperRpcClient.UpdateDeveloper(l.ctx, &developer.UpdateDeveloperReq{
		Id: id,
	})
	if err != nil {
		logx.Errorf("UpdateDeveloper: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	return
}

func (l *GetDeveloperLogic) rpcGetDeveloperById(id int64) (resp *developer.GetDeveloperByIdResp, err error) {
	resp, err = l.svcCtx.DeveloperRpcClient.GetDeveloperById(l.ctx, &developer.GetDeveloperByIdReq{
		Id: id,
	})
	if err != nil {
		logx.Errorf("GetDeveloperByIdRPC: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	return
}
