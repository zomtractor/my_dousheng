package dao

import (
	"database/sql"
	"gorm.io/gorm"
)

func NewVideoOnceInstance() *videoDao {
	videoOnce.Do(
		func() {
			vDao = &videoDao{}
		})
	return vDao
}

func (*videoDao) AddVideoToSql(video *Video) *Video {
	videolock.Lock()
	tx := db.Create(video)
	if tx.Error == nil {
		videolock.Unlock()
		return video
	}

	videolock.Unlock()
	return nil
}

func (*videoDao) GetVideosByUid(uid int) []*Video {
	videos := make([]*Video, 0)
	err := db.Where("author_id = ?", uid).Find(&videos).Error
	if err != nil && err != sql.ErrNoRows {
		return nil
	}
	return videos
}

func (*videoDao) GetVideoByVid(vid int) *Video {
	video := &Video{}
	err := db.Where("id = ?", vid).Find(&video).Error
	if err != nil && err != sql.ErrNoRows {
		return nil
	}
	return video
}

func (*videoDao) UpdateVideo(video *Video) error {
	videolock.Lock()
	err := db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := db.Save(video).Error; err != nil {
			return err
		}
		user := NewUserOnceInstance().GetUserByUserID(video.AuthorId)
		user.VideoCount++
		if err := db.Save(user).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	videolock.Unlock()
	return err
}

func (*videoDao) GetVideos(lastTime int64) []*Video {
	videos := make([]*Video, 0)
	err := db.Where("create_time<?", lastTime).Limit(30).Order("create_time DESC").Find(&videos).Error
	if err != nil && err != sql.ErrNoRows {
		return nil
	}
	return videos
}
