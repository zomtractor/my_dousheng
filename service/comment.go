package service

import (
	"my_dousheng/dao"
	"time"
)

// PublishComment 发布评论
func PublishComment(uid int, vid int, content string) *dao.Comment {
	return dao.NewCommentOnceInstance().AddComment(
		&dao.Comment{
			VideoID: vid,
			UserID:  uid,
			Content: content,
		},
	)
}

// DeleteComment 删除评论
func DeleteComment(uid int, cid int) int {
	commentDao := dao.NewCommentOnceInstance()
	cmt := commentDao.GetCommentByCommentID(cid)
	if cmt == nil {
		return -1
	} else if cmt.UserID != uid {
		return -2
	}
	com := commentDao.DeleteComment(cmt)
	if com == nil {
		return -3
	}
	return 0
}

// GetComentList 获得评论列表
func GetComentList(vid int) []*dao.Comment {
	comments := dao.NewCommentOnceInstance().GetCommentsByVideoID(vid)
	if comments == nil {
		return nil
	}
	fillComments(comments)
	return comments
}

//填充评论信息
func fillComments(com []*dao.Comment) {
	uDao := dao.NewUserOnceInstance()
	for i := 0; i < len(com); i++ {
		t := time.Unix(com[i].CreateTime, 0)
		com[i].User = uDao.GetUserByUserID(com[i].UserID)
		com[i].CreateDate = t.Format("01-02")
	}
}
