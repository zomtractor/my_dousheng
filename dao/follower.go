package dao

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
)

func NewFollowerOnceInstance() *followerDao {
	userOnce.Do(
		func() {
			fDao = &followerDao{}
		})
	return fDao
}

// AddFollower 添加关注
func (*followerDao) AddFollower(f *Follower) error {
	followerdlock.Lock()
	err := db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		follow := &User{}
		follower := &User{}
		if err := tx.Create(f).Error; err != nil {
			return err
		}
		if err := tx.Where("id=?", f.FollowID).Find(follow).Error; err != nil {
			return err
		}
		follow.FollowerCount = follow.FollowerCount + 1
		if err := tx.Save(follow).Error; err != nil {
			return err
		}
		if err := tx.Where("id=?", f.FollowerID).Find(follower).Error; err != nil {
			return err
		}
		follower.FollowCount = follower.FollowCount + 1
		if err := tx.Save(follower).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	followerdlock.Unlock()
	return err
}

// DeleteFollower 取消关注
func (*followerDao) DeleteFollower(f *Follower) error {
	followerdlock.Lock()
	err := db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		follow := &User{}
		follower := &User{}
		res := tx.Delete(&Follower{}, "follow_id=? and follower_id=?", f.FollowID, f.FollowerID)
		if res.Error != nil || res.RowsAffected == 0 {
			return errors.New("delete failure")
		}
		if err := tx.Where("id=?", f.FollowID).Find(follow).Error; err != nil {
			return err
		}
		follow.FollowerCount = follow.FollowerCount - 1
		if err := tx.Save(follow).Error; err != nil {
			return err
		}
		if err := tx.Where("id=?", f.FollowerID).Find(follower).Error; err != nil {
			return err
		}
		follower.FollowCount = follower.FollowCount - 1
		if err := tx.Save(follower).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	followerdlock.Unlock()
	return err
}

// GetFollowerByIDs 获取关注信息
func (*followerDao) GetFollowerByIDs(followID, followerId int) *Follower {
	follow := &Follower{}
	err := db.Where("follow_id=? and follower_id=?", followID, followerId).Find(follow).Error
	if err != nil && err != sql.ErrNoRows {
		return nil
	}
	return follow
}
