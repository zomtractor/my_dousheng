package service

import "my_dousheng/dao"

// Subscribe 关注
func Subscribe(myUid int, hisUid int) error {
	return dao.NewFollowerOnceInstance().AddFollower(
		&dao.Follower{
			FollowID:   hisUid,
			FollowerID: myUid,
		},
	)
}

// DisSubscribe 取消关注
func DisSubscribe(myUid int, hisUid int) error {
	return dao.NewFollowerOnceInstance().DeleteFollower(
		&dao.Follower{
			FollowID:   hisUid,
			FollowerID: myUid,
		},
	)
}

// GetFollowList 获取关注列表
func GetFollowList(uid int) []*dao.User {
	relations := dao.NewFollowerOnceInstance().GetCommentsByFollowerId(uid)
	users := getUserByFollowID(relations, uid)
	return users
}

// GetFanList 获取粉丝列表
func GetFanList(uid int) []*dao.User {
	relations := dao.NewFollowerOnceInstance().GetCommentsByFollowId(uid)
	users := getUserByFollowerID(relations, uid)
	return users
}

//获取关注用户信息
func getUserByFollowID(rels []*dao.Follower, uid int) []*dao.User {
	uDao := dao.NewUserOnceInstance()
	var u *dao.User
	users := make([]*dao.User, 0)
	for i := 0; i < len(rels); i++ {
		u = uDao.GetUserByUserID(rels[i].FollowID)
		setIsFollower(rels[i].FollowID, uid, u)
		users = append(users, u)
	}
	return users
}

//获取关注用户信息
func getUserByFollowerID(rels []*dao.Follower, uid int) []*dao.User {
	uDao := dao.NewUserOnceInstance()
	var u *dao.User
	users := make([]*dao.User, 0)
	for i := 0; i < len(rels); i++ {
		u = uDao.GetUserByUserID(rels[i].FollowerID)
		setIsFollower(rels[i].FollowerID, uid, u)
		users = append(users, u)
	}
	return users
}
