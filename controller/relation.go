package controller

import (
	"github.com/gin-gonic/gin"
	"my_dousheng/service"
	"strconv"
)

// RelationAction 关注/取消关注
func RelationAction(c *gin.Context) {
	toUserId, _ := strconv.Atoi(c.Query("to_user_id"))
	actionType := c.Query("action_type")
	uid := checkToken(c)
	if uid <= 0 {
		c.JSON(200, PublishActionResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "token is invalid",
			}})
		return
	}
	var err error
	if actionType == "1" {
		err = service.Subscribe(uid, toUserId)
	} else {
		err = service.DisSubscribe(uid, toUserId)
	}
	if err != nil {
		c.JSON(200, PublishActionResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "failed",
			}})
	} else {
		c.JSON(200, PublishActionResponse{
			BaseResponse: BaseResponse{
				StatusCode: 0,
			}})
	}
}

// RelationFollowList 获取关注列表
func RelationFollowList(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("user_id"))
	ok := checkToken(c)
	if ok <= 0 {
		c.JSON(200, PublishActionResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "token is invalid",
			}})
		return
	}
	userList := service.GetFollowList(uid)
	c.JSON(200, RelationFollowListResponse{
		BaseResponse: BaseResponse{
			StatusCode: 0,
		},
		UserList: userList,
	})
}

// RelationFollowerList 获取粉丝列表
func RelationFollowerList(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("user_id"))
	ok := checkToken(c)
	if ok <= 0 {
		c.JSON(200, PublishActionResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "token is invalid",
			}})
		return
	}
	userList := service.GetFanList(uid)
	c.JSON(200, RelationFollowListResponse{
		BaseResponse: BaseResponse{
			StatusCode: 0,
		},
		UserList: userList,
	})
}
