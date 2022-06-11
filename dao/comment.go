package dao

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

//单例
func NewCommentOnceInstance() *CommentDao {
	cDaoOnce.Do(
		func() {
			cDao = &CommentDao{}
		})
	return cDao
}

// GetCommentsByVideoID 获取评论信息
func (*CommentDao) GetCommentsByVideoID(vid int) []*Comment {
	comments := make([]*Comment, 0)
	err := db.Where("video_id=?", vid).Order("create_time DESC").Find(&comments).Error
	if err != nil && err != sql.ErrNoRows {
		return nil
	}
	return comments
}

// GetCommentByCommentID 获取评论信息
func (*CommentDao) GetCommentByCommentID(cid int) *Comment {
	comments := Comment{}
	err := db.Where("id=?", cid).Find(&comments).Error
	if err != nil && err != sql.ErrNoRows {
		return nil
	}
	return &comments
}

// GetCommentsByFollowerId 获取评论
func (*followerDao) GetCommentsByFollowerId(uid int) []*Follower {
	followers := make([]*Follower, 0)
	err := db.Where("follower_id=?", uid).Find(&followers).Error
	if err != nil && err != sql.ErrNoRows {
		return nil
	}
	return followers
}

// GetCommentsByFollowId 获取评论
func (*followerDao) GetCommentsByFollowId(uid int) []*Follower {
	followers := make([]*Follower, 0)
	err := db.Where("follow_id=?", uid).Find(&followers).Error
	if err != nil && err != sql.ErrNoRows {
		return nil
	}
	return followers
}

// AddComment 添加评论
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

// DeleteComment 删除评论
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
