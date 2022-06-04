package controller

import (
	"github.com/gin-gonic/gin"
	"my_dousheng/dao"
	"my_dousheng/service"
	"strconv"
)

//用户注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	id, token := service.Register(username, password)
	if id == -1 {
		c.JSON(200, RegisterResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "User already exist",
			},
		})
	} else if id == -2 {
		c.JSON(200, RegisterResponse{
			BaseResponse: BaseResponse{
				StatusCode: 2,
				StatusMsg:  "unknown error",
			},
		})
	} else if id > 0 {
		c.JSON(200, RegisterResponse{
			BaseResponse: BaseResponse{
				StatusCode: 0,
				StatusMsg:  "register success",
			},
			Token:  token,
			UserId: int64(id),
		})
	}
}

//用户信息
func Action(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("user_id"))
	if checkToken(c) <= 0 {
		c.JSON(200, GetUserResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "token is invalid",
			},
			User: dao.User{},
		})
	} else {
		user := service.GetUserByID(uid)
		if user != nil {
			c.JSON(200, GetUserResponse{
				BaseResponse: BaseResponse{
					StatusCode: 0,
				},
				User: *user,
			})
		} else {
			c.JSON(200, GetUserResponse{
				BaseResponse: BaseResponse{
					StatusCode: 2,
					StatusMsg:  "no user info",
				},
				User: dao.User{},
			})
		}
	}
}

//用户登陆
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	id, token := service.LoginByPassword(username, password)
	if id > 0 {
		c.JSON(200, RegisterResponse{
			BaseResponse: BaseResponse{
				StatusCode: 0,
				StatusMsg:  "login success",
			},
			UserId: int64(id),
			Token:  token,
		})
	} else {
		c.JSON(200, RegisterResponse{
			BaseResponse: BaseResponse{
				StatusCode: 1,
				StatusMsg:  "username or password error",
			},
		})
	}
}

func checkToken(c *gin.Context) int {
	token := c.Query("token")

	t := service.CheckToken(token)
	return t
}
