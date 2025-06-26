package languages

import (
	"context"
	"strings"

	githublangsgo "github.com/NDoolan360/github-langs-go"
	"github.com/wushiling50/aster/api/internal/svc"
	"github.com/wushiling50/aster/api/internal/types"
	"github.com/wushiling50/aster/pkg/errno"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLanguagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLanguagesLogic {
	return &GetLanguagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLanguagesLogic) GetLanguages(req *types.GetLanguagesReq) (resp *types.GetLanguagesResp, err error) {
	resp = new(types.GetLanguagesResp)

	var allLang map[string]githublangsgo.Language

	if allLang, err = githublangsgo.GetAllLanguages(); err != nil {
		logx.Errorf("service.GetLanguages: Get Github Languages Data failed: %v", err.Error())
		err = errno.InternalLanguagesError.WithError(err)
		return
	}

	for langName, lang := range allLang {
		resp.LanguageList = append(resp.LanguageList, types.Language{
			Id:    strings.ReplaceAll(strings.ToLower(langName), " ", "-"),
			Name:  langName,
			Color: lang.Color,
		})
	}

	logx.Info("Successfully Get Languages")
	return
}
