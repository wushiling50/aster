package rank

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/wushiling50/aster/api/internal/pack"
	"github.com/wushiling50/aster/api/internal/svc"
	"github.com/wushiling50/aster/api/internal/types"

	"github.com/wushiling50/aster/pkg/constants"

	"github.com/wushiling50/aster/pkg/errno"
	"github.com/wushiling50/aster/pkg/utils"
	analysis "github.com/wushiling50/aster/rpc/analysis/analysisclient"
	developer "github.com/wushiling50/aster/rpc/developer/developerclient"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type GetScoreRankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetScoreRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoreRankLogic {
	return &GetScoreRankLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetScoreRankLogic) GetScoreRank(req *types.GetScoreRankReq) (resp *types.GetScoreRankResp, err error) {
	resp = new(types.GetScoreRankResp)
	var scoreList []redis.FloatPair

	// 获取分数排行 (DeveloperID - Score)
	resp.Rank = make([]*types.DeveloperWithScore, 0, len(scoreList))
	start := req.Limit * req.Offset
	stop := start + req.Limit - 1
	if scoreList, err = l.svcCtx.RankModel.GetScores(l.ctx, constants.ScoreKey, start, stop); err != nil {
		logx.Errorf("service.GetScoreRank: Get Score List failed: %v", err.Error())
		return
	}

	// 整合排名相关信息
	for _, pair := range scoreList {
		var (
			developerId   int64
			nation        string
			languages     = make(map[string]float64)
			typeDeveloper *types.Developer
			typeScore     *types.Score
		)

		if developerId, err = strconv.ParseInt(pair.Key, 10, 64); err != nil {
			logx.Error(err)
			continue
		}

		// filter
		if req.Nation != "" {
			if nation, _, err = l.getNationById(developerId); err != nil {
				logx.Error(err)
				continue
			}
			if nation != req.Nation {
				continue
			}
		}

		if req.Language != "" {
			if languages, err = l.getLanguagesById(developerId); err != nil {
				logx.Error(err)
				continue
			}
			if _, ok := languages[req.Language]; !ok {
				continue
			}
		}

		// 获取开发者信息
		if typeDeveloper, err = l.getDeveloperById(developerId); err != nil {
			logx.Error(err)
			continue
		}

		// 获取分数信息
		if typeScore, err = l.getScoreById(developerId, pair.Score); err != nil {
			logx.Error(err)
			continue
		}

		resp.Rank = append(resp.Rank, &types.DeveloperWithScore{
			Developer: *typeDeveloper,
			Score:     *typeScore,
		})
	}

	resp.Total, err = l.svcCtx.RankModel.GetScoresTotal(l.ctx, constants.ScoreKey)
	if err != nil {
		logx.Errorf("service.GetScoreRank: Get Scores Total failed: %v", err.Error())
		return
	}

	logx.Info("Successfully Get Rank")
	return
}

func (l *GetScoreRankLogic) getNationById(id int64) (nation string, confidence float64, err error) {
	var resp *analysis.GetNationResp

	resp, err = l.svcCtx.AnalysisRpcClient.GetNation(l.ctx, &analysis.GetAnalysisReq{
		DeveloperId: id,
	})
	if err != nil {
		logx.Errorf("GetNationRPC: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	nation = resp.Nation.Nation
	confidence = resp.Nation.Confidence

	return
}

func (l *GetScoreRankLogic) getLanguagesById(developerId int64) (languages map[string]float64, err error) {
	var resp *analysis.GetLanguagesResp

	resp, err = l.svcCtx.AnalysisRpcClient.GetLanguages(l.ctx, &analysis.GetAnalysisReq{
		DeveloperId: developerId,
	})
	if err != nil {
		logx.Errorf("GetLanguagesRPC: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	if err = json.Unmarshal([]byte(resp.Languages.Languages), &languages); err != nil {
		logx.Error(err)
		err = errno.InternalJSONError.WithError(err)
		return
	}

	return
}

func (l *GetScoreRankLogic) getDeveloperById(developerId int64) (typeDeveloper *types.Developer, err error) {
	var resp *developer.GetDeveloperByIdResp

	resp, err = l.svcCtx.DeveloperRpcClient.GetDeveloperById(l.ctx, &developer.GetDeveloperByIdReq{
		Id: developerId,
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

	return
}

func (l *GetScoreRankLogic) getScoreById(developerId int64, score float64) (typeScore *types.Score, err error) {
	var resp *analysis.GetScoreResp

	resp, err = l.svcCtx.AnalysisRpcClient.GetScore(l.ctx, &analysis.GetAnalysisReq{
		DeveloperId: developerId,
	})
	if err != nil {
		logx.Errorf("GetScoreRPC: RPC called failed: %v", err.Error())
		err = errno.InternalServiceError.WithError(err)
		return
	}

	if !utils.IsSuccess(resp.Base) {
		err = errno.BizError.WithMessage(resp.Base.Message)
		return
	}

	typeScore = pack.BuildTypeScore(resp.Score, developerId, score)

	return
}
