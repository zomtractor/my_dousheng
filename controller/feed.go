package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_dousheng/service"
	"strconv"
	"time"
)

//获取推流
func FeedAction(c *gin.Context) {
	lateTime := c.Query("latest_time")
	late, err := strconv.ParseInt(lateTime, 10, 64)

	uid := checkToken(c)
	if err != nil {
		fmt.Println(err)
	}
	videos := service.GetFeed(uid, late, 0)
	if videos == nil {
		c.JSON(200,
			FeedActionResponse{
				BaseResponse: BaseResponse{
					StatusCode: 1,
					StatusMsg:  "video list is nil",
				},
			})
		return
	}
	if len(videos) == 0 {
		c.JSON(200,
			FeedActionResponse{
				BaseResponse: BaseResponse{
					StatusCode: 0,
				},
				VideoList: videos,
				NextTime:  time.Now().Unix(),
			})
		return
	}
	c.JSON(200,
		FeedActionResponse{
			BaseResponse: BaseResponse{
				StatusCode: 0,
			},
			VideoList: videos,
			NextTime:  videos[len(videos)-1].CreateTime,
		})
}
