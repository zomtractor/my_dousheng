package service

import (
	"my_dousheng/dao"
	"time"
)

// GetFeed 获取推流
func GetFeed(uid int, late int64, cnt int) []*dao.Video {
	videos := dao.NewVideoOnceInstance().GetVideos(late)
	if videos == nil {
		cnt++
		late = time.Now().Unix()
		if cnt > 1 {
			return nil
		}
		return GetFeed(uid, late, cnt+1)
	} else {
		cnt = 0
	}
	setFavoriteByVideos(videos, uid)
	setUser(uid, videos)
	return videos
}

//完善点赞信息
func setFavoriteByVideos(videos []*dao.Video, uid int) {
	favoriteDao := dao.NewFavoriteOnceInstance()
	var f *dao.Favorite
	for i := 0; i < len(videos); i++ {
		f = favoriteDao.GetFavorite(uid, videos[i].Id)
		if f.UserID != 0 {
			videos[i].IsFavorite = true
		}
	}
}

//完善关注信息
func setIsFollower(followId, followerId int, user *dao.User) {
	followDao := dao.NewFollowerOnceInstance()
	rel := followDao.GetFollowerByIDs(followId, followerId)
	if rel.ID > 0 {
		user.IsFollow = true
	}
}

//完善用户信息
func setUser(uid int, videos []*dao.Video) {
	uDao := dao.NewUserOnceInstance()
	var u *dao.User
	for i := 0; i < len(videos); i++ {
		u = uDao.GetUserByUserID(videos[i].AuthorId)
		setIsFollower(videos[i].AuthorId, uid, u)
		videos[i].User = u
	}
}
