package service

import "my_dousheng/dao"

// Favorite 点赞
func Favorite(uid int, vid int) error {
	return dao.NewFavoriteOnceInstance().AddFavoriteToSql(
		&dao.Favorite{
			UserID:  uid,
			VideoID: vid,
		},
	)
}

// DisFavorite 取消点赞
func DisFavorite(uid int, vid int) error {
	return dao.NewFavoriteOnceInstance().DeleteFavoriteFromSql(
		&dao.Favorite{
			UserID:  uid,
			VideoID: vid,
		},
	)
}

// GetFavoriteList 获取点赞列表
func GetFavoriteList(uid int) []*dao.Video {
	favorites := dao.NewFavoriteOnceInstance().GetFavoritesByUid(uid)
	videos := getVideosFromFavorites(favorites)
	return videos
}

//获取点赞视频列表
func getVideosFromFavorites(favorites []*dao.Favorite) []*dao.Video {
	if favorites == nil {
		return nil
	}
	videos := make([]*dao.Video, 0)
	videoDao := dao.NewVideoOnceInstance()
	for i := 0; i < len(favorites); i++ {
		video := videoDao.GetVideoByVid(favorites[i].VideoID)
		videos = append(videos, video)
	}
	return videos
}
