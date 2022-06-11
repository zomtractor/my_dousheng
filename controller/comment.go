package controller

import (
	"github.com/gin-gonic/gin"
	"my_dousheng/service"
	"strconv"
)

// CommentList 获取评论列表
func CommentList(c *gin.Context) {
	vid, _ := strconv.Atoi(c.Query("video_id"))
	comments := service.GetComentList(vid)
	if comments != nil {
		c.JSON(200, CommentListResponse{
			BaseResponse: BaseResponse{
				StatusCode: 0,
			},
			CommentList: comments,
		})
	} else {
		c.JSON(200, CommentListResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "unknown error",
			},
			CommentList: comments,
		})
	}
}

// CommentAction 发表评论
func CommentAction(c *gin.Context) {
	vid, _ := strconv.Atoi(c.Query("video_id"))
	actionType := c.Query("action_type")
	uid := checkToken(c)
	if uid <= 0 {
		c.JSON(200, CommentActionResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "token is invalid",
			}})
		return
	}
	if actionType == "1" {
		text := c.Query("comment_text")
		comment := service.PublishComment(uid, vid, text)
		if comment != nil {
			c.JSON(200, CommentActionResponse{
				BaseResponse: BaseResponse{
					StatusCode: 0,
				},
				Comment: comment,
			})
		} else {
			c.JSON(200, CommentActionResponse{
				BaseResponse: BaseResponse{
					StatusCode: 1,
					StatusMsg:  "unknown error",
				},
				Comment: nil,
			})
		}
	} else {
		cid, _ := strconv.Atoi(c.Query("comment_id"))
		statue := service.DeleteComment(uid, cid)
		if statue == -2 {
			c.JSON(200, CommentActionResponse{
				BaseResponse: BaseResponse{
					StatusCode: 2,
					StatusMsg:  "no permission",
				},
				Comment: nil,
			})
		} else if statue < 0 {
			c.JSON(200, CommentActionResponse{
				BaseResponse: BaseResponse{
					StatusCode: 1,
					StatusMsg:  "unknown error",
				},
				Comment: nil,
			})
		} else if statue == 0 {
			c.JSON(200, CommentActionResponse{
				BaseResponse: BaseResponse{
					StatusCode: 0,
				},
				Comment: nil,
			})
		}

	}

}
