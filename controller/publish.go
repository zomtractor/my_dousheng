package controller

import (
	"github.com/gin-gonic/gin"
	"my_dousheng/service"
	"path"
	"strconv"
)

//视频列表
func PublishList(c *gin.Context) {
	if checkToken(c) <= 0 {
		c.JSON(200, PublishListResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "token is invalid",
			}})
		return
	}
	uid, _ := strconv.Atoi(c.Query("user_id"))
	videos := service.GetPublishList(uid)
	if videos == nil {
		c.JSON(200, PublishListResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "unknown err",
			},
			VideoList: nil,
		})
	} else {
		c.JSON(200, PublishListResponse{
			BaseResponse: BaseResponse{
				StatusCode: 0,
			},
			VideoList: videos,
		})
	}
}

//发布视频
func PublishLogin(c *gin.Context) {
	title := c.PostForm("title")
	uid := service.CheckToken(c.PostForm("token"))
	if uid <= 0 {
		c.JSON(200, PublishActionResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "token is invalid",
			}})
		return
	}
	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(200, PublishActionResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "video is invalid",
			}})
		return
	}
	ext := path.Ext(file.Filename)
	video := service.ParseVedio(title, uid, ext)
	if video == nil {
		c.JSON(200, PublishActionResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "unknown err",
			}})
		return
	}
	videoUrl := service.Url_pf + strconv.FormatInt(int64(video.Id), 10) + path.Ext(file.Filename)
	err = c.SaveUploadedFile(file, videoUrl)
	if err != nil {
		return
	}

	c.JSON(200, PublishActionResponse{
		BaseResponse: BaseResponse{
			StatusCode: 0,
			StatusMsg:  "upload success",
		},
	})
}
