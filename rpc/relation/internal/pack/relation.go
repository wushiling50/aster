package pack

import (
	"github.com/wushiling50/aster/gen/relation"
	model_relation "github.com/wushiling50/aster/pkg/model/relation"
)

func BuildModelCreateRepo(createRepo *relation.CreateRepo) *model_relation.CreateRepo {
	return &model_relation.CreateRepo{
		DeveloperId: createRepo.DeveloperId,
		RepoId:      createRepo.RepoId,
	}
}
