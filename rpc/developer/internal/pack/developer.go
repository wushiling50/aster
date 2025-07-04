package pack

import (
	"time"

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
		CreatedAt:       time.Unix(developer.CreatedAt, 0),
		UpdatedAt:       time.Unix(developer.UpdatedAt, 0),
	}
}

func BuildGenDeveloper(mdoelDveloper *model_developer.Developer) *developer.Developer {
	return &developer.Developer{
		Id:              mdoelDveloper.Id,
		Name:            mdoelDveloper.Name,
		Login:           mdoelDveloper.Login,
		AvatarUrl:       mdoelDveloper.AvatarUrl,
		Company:         mdoelDveloper.Company,
		Location:        mdoelDveloper.Location,
		Bio:             mdoelDveloper.Bio,
		Blog:            mdoelDveloper.Blog,
		Email:           mdoelDveloper.Email,
		TwitterUsername: mdoelDveloper.TwitterUsername,
		Repos:           mdoelDveloper.Repos,
		Following:       mdoelDveloper.Following,
		Followers:       mdoelDveloper.Followers,
		Stars:           mdoelDveloper.Stars,
		Gists:           mdoelDveloper.Gists,
		CreatedAt:       mdoelDveloper.CreatedAt.Unix(),
		UpdatedAt:       mdoelDveloper.UpdatedAt.Unix(),
	}
}
