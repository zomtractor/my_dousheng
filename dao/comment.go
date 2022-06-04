package dao

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

func NewCommentOnceInstance() *CommentDao {
	cDaoOnce.Do(
		func() {
			cDao = &CommentDao{}
		})
	return cDao
}

func (*CommentDao) GetCommentsByVideoID(vid int) []*Comment {
	comments := make([]*Comment, 0)
	err := db.Where("video_id=?", vid).Order("create_time DESC").Find(&comments).Error
	if err != nil && err != sql.ErrNoRows {
		return nil
	}
	return comments
}
func (*CommentDao) GetCommentByCommentID(cid int) *Comment {
	comments := Comment{}
	err := db.Where("id=?", cid).Find(&comments).Error
	if err != nil && err != sql.ErrNoRows {
		return nil
	}
	return &comments
}

func (*CommentDao) AddComment(com *Comment) *Comment {
	com.CreateTime = time.Now().Unix()
	commentlock.Lock()
	err := db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		video := &Video{}
		if err := tx.Create(com).Error; err != nil {
			return err
		}
		if err := tx.Where("id=?", com.VideoID).Find(video).Error; err != nil {
			return err
		}
		video.CommentCount = video.CommentCount + 1
		if err := tx.Save(video).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	commentlock.Unlock()
	if err != nil {
		return nil
	}
	return com
}
func (*CommentDao) DeleteComment(com *Comment) *Comment {
	commentlock.Lock()
	err := db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Delete(com).Error; err != nil {
			return err
		}
		video := &Video{}
		if err := tx.Where("id=?", com.VideoID).Find(video).Error; err != nil {
			return err
		}
		video.CommentCount = video.CommentCount - 1
		if err := tx.Save(video).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	commentlock.Unlock()
	if err != nil {
		return nil
	}
	return com
}
