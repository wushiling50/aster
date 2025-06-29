package pack

import (
	"github.com/wushiling50/aster/pkg/model/relation"
)

func BuildFollow(followerId int64, followingId int64) *relation.Follow {
	return &relation.Follow{
		FollowerId:  followerId,
		FollowingId: followingId,
	}
}

func BuildCompletedFollow(dataId int, userId int64) *relation.Follow {
	return &relation.Follow{
		DataId:      int64(dataId),
		FollowerId:  userId,
		FollowingId: 0,
	}
}
