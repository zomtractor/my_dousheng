package controller

import "my_dousheng/dao"

type BaseResponse struct {
	StatusCode int    `json:"status_code,omitempty"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type GetUserResponse struct {
	BaseResponse
	User dao.User `json:"user"`
}

type RegisterResponse struct {
	BaseResponse
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}
type CommentActionResponse struct {
	BaseResponse
	Comment *dao.Comment `json:"comment"`
}

type CommentListResponse struct {
	BaseResponse
	CommentList []*dao.Comment `json:"comment_list"`
}

type FavoriteActionResponse struct {
	BaseResponse
}

type FavoriteListResponse struct {
	BaseResponse
	VideoList []*dao.Video `json:"video_list"`
}
type FeedActionResponse struct {
	BaseResponse
	NextTime  int64        `json:"next_time"`
	VideoList []*dao.Video `json:"video_list"`
}

type PublishActionResponse struct {
	BaseResponse
}

type PublishListResponse struct {
	BaseResponse
	VideoList []*dao.Video `json:"video_list"`
}
type RelationActionResponse struct {
	BaseResponse
}

type RelationFollowListResponse struct {
	BaseResponse
	UserList []*dao.User `json:"user_list"`
}
