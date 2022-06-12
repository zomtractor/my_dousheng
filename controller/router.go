package controller

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	v1 := r.Group("/douyin")
	v1.GET("/user/", Action)
	v1.POST("/user/register/", Register)
	v1.POST("/user/login/", Login)

	v1.POST("/publish/action/", PublishLogin)
	v1.GET("/publish/list/", PublishList)

	v1.GET("/feed/", FeedAction)

	//v1.GET("/video/", controller.FeedVideo)

	v1.POST("favorite/action/", FavoriteAction)
	v1.GET("favorite/list/", FavoriteList)

	v1.POST("comment/action/", CommentAction)
	v1.GET("comment/list/", CommentList)

	v1.POST("relation/action/", RelationAction)
	v1.GET("relation/follow/list", RelationFollowList)
	v1.GET("relation/follower/list", RelationFollowerList)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
