package controller

import "my_dousheng/dao"

// BaseResponse 基本返回类型
type BaseResponse struct {
	StatusCode int    `json:"status_code,omitempty"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// GetUserResponse 用户返回类型
type GetUserResponse struct {
	BaseResponse
	User dao.User `json:"user"`
}

// RegisterResponse 注册返回类型
type RegisterResponse struct {
	BaseResponse
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

// CommentActionResponse 评论返回类型
type CommentActionResponse struct {
	BaseResponse
	Comment *dao.Comment `json:"comment"`
}

// CommentListResponse 评论列表返回类型
type CommentListResponse struct {
	BaseResponse
	CommentList []*dao.Comment `json:"comment_list"`
}

// FavoriteActionResponse 点赞返回类型
type FavoriteActionResponse struct {
	BaseResponse
}

// FavoriteListResponse 点赞列表返回类型
type FavoriteListResponse struct {
	BaseResponse
	VideoList []*dao.Video `json:"video_list"`
}

// FeedActionResponse 推流返回类型
type FeedActionResponse struct {
	BaseResponse
	NextTime  int64        `json:"next_time"`
	VideoList []*dao.Video `json:"video_list"`
}

// PublishActionResponse 发布返回类型
type PublishActionResponse struct {
	BaseResponse
}

// PublishListResponse 发布列表返回类型
type PublishListResponse struct {
	BaseResponse
	VideoList []*dao.Video `json:"video_list"`
}

// RelationActionResponse 关注返回类型
type RelationActionResponse struct {
	BaseResponse
}

// RelationFollowListResponse 关注列表返回类型
type RelationFollowListResponse struct {
	BaseResponse
	UserList []*dao.User `json:"user_list"`
}
