package controller

import (
	"github.com/gin-gonic/gin"
	"my_dousheng/service"
	"strconv"
)

// FavoriteAction 点赞/取消点赞
func FavoriteAction(c *gin.Context) {
	vid, _ := strconv.Atoi(c.Query("video_id"))
	actionType := c.Query("action_type")
	err := error(nil)
	uid := checkToken(c)
	if uid <= 0 {
		c.JSON(200, FavoriteActionResponse{BaseResponse{
			StatusCode: 1,
			StatusMsg:  "token is invalid",
		}})
		return
	}
	if actionType == "1" {
		err = service.Favorite(uid, vid)
	} else {
		err = service.DisFavorite(uid, vid)
	}
	if err == nil {
		c.JSON(200, FavoriteActionResponse{BaseResponse{
			StatusCode: 0,
		}})
	} else {
		c.JSON(200, FavoriteActionResponse{BaseResponse{
			StatusCode: 1,
			StatusMsg:  "uknown error",
		}})
	}

}

// FavoriteList 用户点赞列表
func FavoriteList(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("user_id"))
	ok := checkToken(c)
	if ok <= 0 {
		c.JSON(200, PublishListResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "token is invalid",
			}})
		return
	}
	videos := service.GetFavoriteList(uid)
	if videos != nil {
		c.JSON(200, PublishListResponse{
			BaseResponse: BaseResponse{
				StatusCode: 0,
			},
			VideoList: videos,
		})
	} else {
		c.JSON(200, PublishListResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "unknown error",
			},
			VideoList: videos,
		})
	}

}
