package pack

import (
	"github.com/wushiling50/aster/gen/developer"
	model_developer "github.com/wushiling50/aster/pkg/model/developer"
)

func BuildModelDeveloper(developer *developer.Developer) *model_developer.Developer {
	return &model_developer.Developer{
		Id:              developer.Id,
		Name:            developer.Name,
		Login:           developer.Login,
		AvatarUrl:       developer.AvatarUrl,
		Company:         developer.Company,
		Location:        developer.Location,
		Bio:             developer.Bio,
		Blog:            developer.Blog,
		Email:           developer.Email,
		TwitterUsername: developer.TwitterUsername,
		Repos:           developer.Repos,
		Following:       developer.Following,
		Followers:       developer.Followers,
		Stars:           developer.Stars,
		Gists:           developer.Gists,
	}
}

func BuildGenDeveloper(modelDeveloper *model_developer.Developer) *developer.Developer {
	return &developer.Developer{
		Id:              modelDeveloper.Id,
		Name:            modelDeveloper.Name,
		Login:           modelDeveloper.Login,
		AvatarUrl:       modelDeveloper.AvatarUrl,
		Company:         modelDeveloper.Company,
		Location:        modelDeveloper.Location,
		Bio:             modelDeveloper.Bio,
		Blog:            modelDeveloper.Blog,
		Email:           modelDeveloper.Email,
		TwitterUsername: modelDeveloper.TwitterUsername,
		Repos:           modelDeveloper.Repos,
		Following:       modelDeveloper.Following,
		Followers:       modelDeveloper.Followers,
		Stars:           modelDeveloper.Stars,
		Gists:           modelDeveloper.Gists,
	}
}
