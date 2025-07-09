package logic

import (
	"context"
	"errors"

	"github.com/wushiling50/aster/gen/analysis"
	"github.com/wushiling50/aster/pkg/errno"
	model_analysis "github.com/wushiling50/aster/pkg/model/analysis"
	"github.com/wushiling50/aster/rpc/analysis/internal/pack"
	"github.com/wushiling50/aster/rpc/analysis/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoreLogic {
	return &GetScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetScoreLogic) GetScore(in *analysis.GetAnalysisReq) (*analysis.GetScoreResp, error) {
	resp := new(analysis.GetScoreResp)

	score, err := l.getScore(in.DeveloperId)
	if err != nil {
		if errors.Is(err, model_analysis.ErrNotFound) {
			err = errno.BizNotFoundError.WithError(err)
		}

		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildSuccessResp()
	resp.Score = score
	return resp, nil
}

func (l *GetScoreLogic) getScore(developerId int64) (*analysis.Score, error) {
	var score *model_analysis.Score

	score, err := l.svcCtx.ScoreModel.FindOneByDeveloperId(l.ctx, developerId)
	if err != nil {
		return nil, err
	}

	return pack.BuildGenScore(score), err
}
