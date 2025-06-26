package developer

import (
	"context"

	"github.com/wushiling50/aster/api/internal/pack"
	"github.com/wushiling50/aster/api/internal/svc"
	"github.com/wushiling50/aster/api/internal/types"
	"github.com/wushiling50/aster/pkg/errno"
	"github.com/wushiling50/aster/pkg/github"
	"github.com/wushiling50/aster/pkg/utils"
	developer "github.com/wushiling50/aster/rpc/developer/developerclient"

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

func (l *GetDeveloperLogic) GetDeveloper(req *types.GetDeveloperReq) (resp *types.GetDeveloperResp, err error) {
	resp = new(types.GetDeveloperResp)

	var (
		id      int64
		rpcResp *types.Developer
	)

	if id, err = github.GetIdByLogin(l.ctx, req.Login); err != nil {
		logx.Errorf("service.GetDeveloper: Failed To Get Id By Login %v", err.Error())
		return
	}

	if err = l.rpcUpdateDeveloperById(id); err != nil {
		logx.Error(err)
		return
	}

	if rpcResp, err = l.rpcGetDeveloperById(id); err != nil {
		logx.Error(err)
		return
	}

	resp.Developer = *rpcResp

	logx.Info("Successfully Get Developer")
	return
}

func (l *GetDeveloperLogic) rpcUpdateDeveloperById(id int64) (err error) {
	var resp *developer.UpdateDeveloperResp

	resp, err = l.svcCtx.DeveloperRpcClient.UpdateDeveloper(l.ctx, &developer.UpdateDeveloperReq{
		Id: id,
	})
	if err != nil {
		logx.Errorf("UpdateDeveloperByIdRPC: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	logx.Info("Successfully Update Developer By Id")
	return
}

func (l *GetDeveloperLogic) rpcGetDeveloperById(id int64) (typeDeveloper *types.Developer, err error) {
	var resp *developer.GetDeveloperByIdResp

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

	typeDeveloper = pack.BuildTypeDeveloper(resp.Developer)

	logx.Info("Successfully Get Developer By Id")
	return
}
